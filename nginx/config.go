// Package nginx
// Date: 2022/11/9 10:18
// Author: Amu
// Description:
package nginx

import "fmt"

type Config struct {
	nextId       func() int64
	config       string
	root         Instruction
	instructions map[int64]Instruction
}

type ConfigBuilder interface {
	Build(c *Config) (Instruction, error)
}

func NewConfig(config string) *Config {
	conf := &Config{
		nextId: func() func() int64 {
			var id int64
			return func() (ret int64) {
				ret = id
				id++
				return
			}
		}(),
		config:       config,
		root:         newNodeInstance(-1, nil, nil),
		instructions: make(map[int64]Instruction),
	}
	conf.instructions[conf.root.Id()] = conf.root
	return conf
}

func (c *Config) Generate() (string, error) {
	for _, instruction := range c.root.BlockInstructions() {
		err := Populate(instruction)
		if err != nil {
			return "", err
		}
	}
	ret := ""
	for _, instruction := range c.root.BlockInstructions() {
		s, err := PrintCommand(instruction.Command())
		if err != nil {
			return "", err
		}
		ret += s
	}
	return ret, nil
}

/* use (id = -1) to add into top level */

func (c *Config) Root() Instruction {
	return c.root
}

func (c *Config) Add(id int64, instructions ...Instruction) error {
	var find func(id int64, root Instruction) Instruction
	find = func(id int64, root Instruction) Instruction {
		if root == nil {
			return nil
		}
		_, idMapping := root.BlockInstructionMapping()
		if n, exists := idMapping[id]; exists {
			return root.BlockInstructions()[n]
		} else {
			if root.Id() == id {
				return root
			}
			for _, i := range root.BlockInstructions() {
				if p := find(id, i); p != nil {
					return p
				}
			}
		}
		return nil
	}
	parent := find(id, c.root)
	if parent == nil {
		return fmt.Errorf("not found")
	}
	return parent.AddInstruction(instructions...)
}

func (c *Config) Del(id int64) error {
	var findParent func(id int64, root Instruction) Instruction
	findParent = func(id int64, root Instruction) Instruction {
		if root == nil {
			return nil
		}
		_, idMapping := root.BlockInstructionMapping()
		if _, exists := idMapping[id]; exists {
			return root
		}
		for _, i := range root.BlockInstructions() {
			if p := findParent(id, i); p != nil {
				return p
			}
		}
		return nil
	}

	if id == -1 {
		return fmt.Errorf("invalid")
	}
	parent := findParent(id, c.root)
	if parent == nil {
		return fmt.Errorf("invalid")
	}
	ins, err := parent.FindSubById(id)
	if err != nil {
		return err
	}
	return parent.DelInstruction(ins)
}

func (c *Config) Move(id int64, pos int, absolute bool) error {
	var findParent func(id int64, root Instruction) Instruction
	findParent = func(id int64, root Instruction) Instruction {
		if root == nil {
			return nil
		}
		_, idMapping := root.BlockInstructionMapping()
		if _, exists := idMapping[id]; exists {
			return root
		}
		for _, i := range root.BlockInstructions() {
			if p := findParent(id, i); p != nil {
				return p
			}
		}
		return nil
	}

	if id == -1 {
		return fmt.Errorf("invalid")
	}
	parent := findParent(id, c.root)
	if parent == nil {
		return fmt.Errorf("invalid")
	}
	ins, err := parent.FindSubById(id)
	if err != nil {
		return err
	}
	return parent.MoveInstruction(ins, pos, absolute)
}

/* add format (eg: new line, comment) into block of this instruction */

func (c *Config) AddFormat(id int64, args ...string) error {
	return nil
}

func (c *Config) ApplyBuilders(builders ...ConfigBuilder) error {
	for _, builder := range builders {
		if ins, err := builder.Build(c); err != nil {
			return err
		} else if err := c.Add(-1, ins); err != nil {
			return err
		}
	}
	return nil
}

func (c *Config) Instructions() []Instruction {
	return c.root.BlockInstructions()
}

func (c *Config) NewInstruction(args ...string) (Instruction, error) {
	cmd := &NodeCommand{
		Indent: 0,
		Args:   nil,
		Block:  nil,
	}
	instruction := newNodeInstance(c.nextId(), cmd, nil) // no block for now
	if err := instruction.AddProperty(args...); err != nil {
		return nil, err
	}
	c.instructions[instruction.Id()] = instruction
	return instruction, nil
}
