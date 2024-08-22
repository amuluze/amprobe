// Package nginx
// Date: 2024/8/6 16:34
// Author: Amu
// Description:
package nginx

// Property

type PropertyBuilder struct {
	name    string
	options []string
	kvs     []struct {
		key   string
		value string
	}
}

func NewPropertyBuilder(name string) *PropertyBuilder {
	return &PropertyBuilder{
		name: name,
	}
}

// Set update propertyBuilder kvs
func (o *PropertyBuilder) Set(key, value string) *PropertyBuilder {
	o.kvs = append(o.kvs, struct {
		key   string
		value string
	}{key: key, value: value})
	return o
}

// Add update propertyBuilder options
func (o *PropertyBuilder) Add(values ...string) *PropertyBuilder {
	o.options = append(o.options, values...)
	return o
}

func (o *PropertyBuilder) Build(c *Config) (Instruction, error) {
	ins, err := c.NewInstruction(o.name)
	if err != nil {
		return nil, err
	}
	if err = ins.AddProperty(o.options...); err != nil {
		return nil, err
	}
	for _, kv := range o.kvs {
		err := ins.SetKeywordProperty(kv.key, kv.value)
		if err != nil {
			return nil, err
		}
	}
	return ins, nil
}

// Block

type BlockBuilder struct {
	name     string
	builders []ConfigBuilder
}

func NewBlockBuilder(name string) *BlockBuilder {
	return &BlockBuilder{
		name: name,
	}
}

func (b *BlockBuilder) addBuilder(builders ...ConfigBuilder) {
	b.builders = append(b.builders, builders...)
}

func (b *BlockBuilder) AddProperty(name string, ops ...string) *PropertyBuilder {
	p := NewPropertyBuilder(name).Add(ops...)
	b.addBuilder(p)
	return p
}

func (b *BlockBuilder) Build(c *Config) (Instruction, error) {
	ins, err := c.NewInstruction(b.name)
	if err != nil {
		return nil, err
	}

	for _, sub := range b.builders {
		i, err := sub.Build(c)
		if err != nil {
			return nil, err
		}

		if err := ins.AddInstruction(i); err != nil {
			return nil, err
		}
	}

	return ins, err
}

// Upstream

type UpstreamBlockBuilder struct {
	*BlockBuilder
	upstreamName string
}

func NewUpstreamBlock(upstreamName string) *UpstreamBlockBuilder {
	return &UpstreamBlockBuilder{
		BlockBuilder: NewBlockBuilder("upstream"),
		upstreamName: upstreamName,
	}
}

func (b *UpstreamBlockBuilder) UpstreamName() string {
	return b.upstreamName
}

func (b *UpstreamBlockBuilder) AddServer(address string) *PropertyBuilder {
	return b.AddProperty("server").Add(address)
}

func (b *UpstreamBlockBuilder) Build(c *Config) (Instruction, error) {
	ins, err := b.BlockBuilder.Build(c)
	if err != nil {
		return nil, err
	}
	return ins, ins.AddProperty(b.upstreamName)
}

// Server

type ServerBlockBuilder struct {
	*BlockBuilder
}

func NewServerBlock() *ServerBlockBuilder {
	return &ServerBlockBuilder{BlockBuilder: NewBlockBuilder("server")}
}

func (b *ServerBlockBuilder) AddListen(address string, ops ...string) *PropertyBuilder {
	return b.AddProperty("listen").Add(address).Add(ops...)
}

func (b *ServerBlockBuilder) AddServerNames(names ...string) *ServerBlockBuilder {
	b.AddProperty("server_name", names...)
	return b
}

func (b *ServerBlockBuilder) AddLocationB(lb *LocationBlockBuilder) *ServerBlockBuilder {
	b.addBuilder(lb)
	return b
}

func (b *ServerBlockBuilder) AddPropertyB(name string, ops ...string) *ServerBlockBuilder {
	b.AddProperty(name, ops...)
	return b
}

// Location

type LocationBlockBuilder struct {
	*BlockBuilder
	path string
}

func NewLocationBlock(path string) *LocationBlockBuilder {
	return &LocationBlockBuilder{BlockBuilder: NewBlockBuilder("location"), path: path}
}

func (b *LocationBlockBuilder) AddPropertyB(name string, ops ...string) *LocationBlockBuilder {
	b.AddProperty(name, ops...)
	return b
}

func (b *LocationBlockBuilder) Build(c *Config) (Instruction, error) {
	ins, err := b.BlockBuilder.Build(c)
	if err != nil {
		return nil, err
	}

	return ins, ins.AddProperty(b.path)
}

type MixBlockBuilder struct {
	LocationBuilders []*LocationBlockBuilder
}
