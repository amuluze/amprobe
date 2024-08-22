// Package nginx
// Date: 2024/8/6 13:57
// Author: Amu
// Description:
package nginx

import (
	"fmt"
	"regexp"
)

var (
	KeywordPattern = regexp.MustCompile("(.+)=(.+)")
)

type NodeArg struct {
	Type   int
	Text   string
	Format *NodeFormatter
}

type NodeBlock struct {
	BlkBgn   *NodeArg
	Commands []*NodeCommand
	BlkEnd   *NodeArg
}

type NodeCommand struct {
	Indent int
	Args   []*NodeArg
	Block  *NodeBlock
}

func (c *NodeCommand) PushArg(a *NodeArg) {
	c.Args = append(c.Args, a)
}

func (c *NodeCommand) PushBlockBegin(a *NodeArg) error {
	if c.Block != nil {
		return fmt.Errorf("unexpected block begin")
	}
	c.Block = &NodeBlock{
		BlkBgn: a,
	}
	return nil
}

func (c *NodeCommand) PushBlockCommand(a *NodeCommand) error {
	if c.Block == nil {
		return fmt.Errorf("unexpected block command")
	}
	c.Block.Commands = append(c.Block.Commands, a)
	return nil
}

func (c *NodeCommand) PushBlockEnd(a *NodeArg) error {
	if c.Block == nil {
		return fmt.Errorf("unexpected block command")
	}
	c.Block.BlkEnd = a
	return nil
}

type NodeFormatter struct {
	Type int    // comment or new line
	Text string // text for comment
}

type NodeInstance struct {
	id          int64
	cmd         *NodeCommand
	block       []Instruction
	nameMapping map[string][]int
	idMapping   map[int64]int
}

func newNodeInstance(id int64, cmd *NodeCommand, block []Instruction) *NodeInstance {
	n, i := MakeInstructionMappingFromSlice(block)
	return &NodeInstance{
		id:          id,
		cmd:         cmd,
		block:       block,
		nameMapping: n,
		idMapping:   i,
	}
}

func MakeInstructionMappingFromSlice(instructions []Instruction) (map[string][]int, map[int64]int) {
	nameMapping := make(map[string][]int)
	idMapping := make(map[int64]int)
	for index := range instructions {
		name := instructions[index].Name()
		nameMapping[name] = append(nameMapping[name], index)
		idMapping[instructions[index].Id()] = index
	}
	return nameMapping, idMapping
}

func (n *NodeInstance) Id() int64 {
	return n.id
}

func (n *NodeInstance) Name() string {
	if len(n.cmd.Args) == 0 || n.cmd.Args[0].Type != ArgTypeIdent {
		return InstructionNameUnknown
	}
	return n.cmd.Args[0].Text
}

func (n *NodeInstance) Rename(name string) {
	if n.Name() != InstructionNameUnknown {
		n.cmd.Args[0].Text = name
	}
}

func (n *NodeInstance) Command() *NodeCommand {
	return n.cmd
}

func (n *NodeInstance) ReMapping() {
	n.nameMapping, n.idMapping = MakeInstructionMappingFromSlice(n.block)
}

/* plain Property */

func (n *NodeInstance) GetProperties() []string {
	var ps []string
	for _, arg := range n.cmd.Args {
		switch arg.Type {
		case ArgTypeIdent, ArgTypeString:
			if !KeywordPattern.MatchString(arg.Text) {
				ps = append(ps, arg.Text)
			}
		default:
			// pass
		}
	}
	return ps
}

func (n *NodeInstance) AddProperty(ps ...string) error {
	for _, p := range ps {
		if len(p) == 0 {
			return fmt.Errorf("invalid")
		}
		// add
		if err := n.addCommandArg(p); err != nil {
			return err
		}
	}
	return nil
}

func (n *NodeInstance) DelProperty(ps ...string) error {
	for _, p := range ps {
		if len(p) == 0 {
			return fmt.Errorf("invalid")
		}

		// del
		if err := n.delCommandArg(p); err != nil {
			return err
		}
	}
	return nil
}

func (n *NodeInstance) ClearProperty() {
	args := make([]*NodeArg, 1)
	args[0] = n.cmd.Args[0]
	n.cmd.Args = args
}

/* keyword Property */

func (n *NodeInstance) GetKeywordProperties() map[string]string {
	ps := make(map[string]string)
	for _, arg := range n.cmd.Args {
		switch arg.Type {
		case ArgTypeIdent, ArgTypeString:
			if match := KeywordPattern.FindStringSubmatch(arg.Text); len(match) == 3 {
				ps[match[1]] = match[2]
			}
		default:
			// pass
		}
	}
	if len(ps) == 0 {
		return nil
	}
	return ps
}

func (n *NodeInstance) GetKeywordProperty(k string) string {
	ps := n.GetKeywordProperties()
	if v, ok := ps[k]; ok {
		return v
	}
	return ""
}

func (n *NodeInstance) SetKeywordProperty(k, v string) error {
	// check
	if len(k) == 0 || len(v) == 0 {
		return fmt.Errorf("invalid")
	}
	ps := n.GetKeywordProperties()
	if _, ok := ps[k]; ok {
		if err := n.DelKeywordProperty(k); err != nil {
			return err
		}
	}
	// add
	return n.addCommandArg(fmt.Sprintf("%s=%s", k, v))
}

func (n *NodeInstance) DelKeywordProperty(k string) error {
	// check
	if len(k) == 0 {
		return fmt.Errorf("invalid")
	}
	ps := n.GetKeywordProperties()
	if _, ok := ps[k]; !ok {
		return fmt.Errorf("already exists")
	}
	return n.delCommandArg(fmt.Sprintf("%s=%s", k, ps[k]))
}

/* block */

func (n *NodeInstance) BlockInstructions() []Instruction {
	return n.block
}

func (n *NodeInstance) BlockInstructionMapping() (map[string][]int, map[int64]int) {
	return n.nameMapping, n.idMapping
}

func (n *NodeInstance) AddInstruction(instructions ...Instruction) error {
	for _, instruction := range instructions {
		if i, existed := n.idMapping[instruction.Id()]; existed {
			if instruction != n.block[i] {
				return fmt.Errorf("invalid")
			}
			// existed item, ignore
			continue
		}

		// instruction
		index := len(n.block)
		name := instruction.Name()
		n.nameMapping[name] = append(n.nameMapping[name], index)
		n.idMapping[instruction.Id()] = index
		n.block = append(n.block, instruction)

		if n.cmd != nil { // exception: root instruction
			// create block if needed
			if n.cmd.Block == nil {
				n.cmd.Block = &NodeBlock{
					BlkBgn:   nil,
					Commands: nil,
					BlkEnd:   nil,
				}
			}
			// add
			n.cmd.Block.Commands = append(n.cmd.Block.Commands, instruction.Command())
		}
	}
	return nil
}

func (n *NodeInstance) DelInstruction(instructions ...Instruction) error {
	var err = fmt.Errorf("not found")
	for _, instruction := range instructions {
		name := instruction.Name()
		id := instruction.Id()
		if index, exists := n.idMapping[id]; exists {
			err = nil // if all instruction not found, return ErrNotFound, otherwise nil
			// check if instruction matches mappings
			for _, v := range n.nameMapping[name] {
				if v == index && n.block[v] == instruction {
					goto VALID
				}
			}
			return fmt.Errorf("invalid")
		VALID:
			// instruction
			n.block = append(n.block[:index], n.block[index+1:]...)
			// frequently rebuild will reduce efficiency
			n.ReMapping()
			if n.cmd != nil {
				indexOfCmd := index
				if n.cmd.Block.Commands[index] != instruction.Command() {
					return fmt.Errorf("invalid")
				}
				n.cmd.Block.Commands = append(n.cmd.Block.Commands[:indexOfCmd], n.cmd.Block.Commands[indexOfCmd+1:]...)
			}
		}
	}
	return err
}

func (n *NodeInstance) MoveInstruction(instruction Instruction, pos int, absolute bool) error {
	if len(n.block) == 1 { // if only one, do nothing
		return nil
	}
	id := instruction.Id()
	if index, exists := n.idMapping[id]; exists {
		// check if instruction matches mappings
		if n.block[index] != instruction {
			return fmt.Errorf("invalid")
		}
		if n.cmd != nil && n.cmd.Block.Commands[index] != instruction.Command() {
			// FIXME: if here returned, it must be something fatal wrong
			return fmt.Errorf("invalid")
		}
		// since will rebuild mapping later, so name mapping can be ignored

		if !absolute {
			pos = index + pos
		}
		// TODO: may cancel this limit
		if pos > len(n.block) || pos < 0 {
			return fmt.Errorf("invalid")
		}
		var fd, bk int
		// instruction
		switch {
		/*
			*	     fd       bk
			*	[][][0][][][][x][][][]
			back:
			*	[:fd],[fd+1:bk+1],0,[bk+1:]
			forward:
			*	[:fd],x,[fd:bk],[bk+1:]
		*/
		case index < pos: // move back
			fd, bk = index, pos
			newSlice := make([]Instruction, 0)
			newSlice = append(newSlice, n.block[:fd]...)
			newSlice = append(newSlice, n.block[fd+1:bk+1]...)
			newSlice = append(newSlice, instruction)
			n.block = append(newSlice, n.block[bk+1:]...)
		case index > pos: // move forward
			fd, bk = pos, index
			newSlice := make([]Instruction, 0)
			newSlice = append(newSlice, n.block[:fd]...)
			newSlice = append(newSlice, instruction)
			newSlice = append(newSlice, n.block[fd:bk]...)
			n.block = append(newSlice, n.block[bk+1:]...)
		default:
			// dont move
		}
		// adjust mapping
		n.ReMapping()
		// command
		if n.cmd != nil { // exception: root instruction
			switch {
			/*
				*	     fd       bk
				*	[][][0][][][][x][][][]
				back:
				*	[:fd],[fd+1:bk+1],0,[bk+1:]
				forward:
				*	[:fd],x,[fd:bk],[bk+1:]
			*/
			case index < pos: // move back
				fd, bk = index, pos
				newSlice := make([]*NodeCommand, 0) // MUST CREATE NEW
				newSlice = append(newSlice, n.cmd.Block.Commands[:fd]...)
				newSlice = append(newSlice, n.cmd.Block.Commands[fd+1:bk+1]...)
				newSlice = append(newSlice, n.cmd.Block.Commands[index])
				n.cmd.Block.Commands = append(newSlice, n.cmd.Block.Commands[bk+1:]...)
			case index > pos: // move forward
				fd, bk = pos, index
				newSlice := make([]*NodeCommand, 0) // MUST CREATE NEW
				newSlice = append(newSlice, n.cmd.Block.Commands[:fd]...)
				newSlice = append(newSlice, n.cmd.Block.Commands[index])
				newSlice = append(newSlice, n.cmd.Block.Commands[fd:bk]...)
				n.cmd.Block.Commands = append(newSlice, n.cmd.Block.Commands[bk+1:]...)
			default:
				// dont move
			}
		}
		return nil
	}
	return fmt.Errorf("not found")
}

func (n *NodeInstance) FindSubByName(name string) ([]Instruction, error) {
	if v, exists := n.nameMapping[name]; exists {
		ret := make([]Instruction, len(v))
		for item, index := range v {
			ret[item] = n.block[index]
		}
		return ret, nil
	}
	return nil, fmt.Errorf("not found")
}

func (n *NodeInstance) FindSubById(id int64) (Instruction, error) {
	if index, exists := n.idMapping[id]; exists {
		return n.block[index], nil
	} else {
		return nil, fmt.Errorf("not found")
	}
}

// private

func (n *NodeInstance) addCommandArg(a string) error {
	index := -1
loop:
	// find index to add
	for an := len(n.cmd.Args) - 1; an >= 0; an-- {
		switch n.cmd.Args[an].Type {
		case ArgTypeIdent, ArgTypeString:
			index = an
			break loop
		default:
			// pass
		}
	}
	// add
	if index < 0 {
		// not found
		n.cmd.Args = append(
			[]*NodeArg{
				{
					Type:   ArgTypeIdent,
					Text:   a,
					Format: nil,
				},
			},
			n.cmd.Args...,
		)
	} else {
		// found
		n.cmd.Args = append(
			n.cmd.Args[:index],
			append(
				[]*NodeArg{
					{
						Type:   n.cmd.Args[index].Type,
						Text:   n.cmd.Args[index].Text,
						Format: nil,
					},
					{
						Type:   ArgTypeIdent,
						Text:   a,
						Format: n.cmd.Args[index].Format,
					},
				},
				n.cmd.Args[index+1:]...,
			)...,
		)
	}
	return nil
}

func (n *NodeInstance) delCommandArg(a string) error {
	index := -1
loop:
	for n, arg := range n.cmd.Args {
		switch arg.Type {
		case ArgTypeIdent, ArgTypeString:
			if arg.Text == a {
				index = n
				break loop
			}
		default:
			// pass
		}
	}

	if index == -1 {
		return fmt.Errorf("not found")
	}
	// delete
	if n.cmd.Args[index].Format != nil {
		if index > 0 && n.cmd.Args[index-1].Format == nil {
			n.cmd.Args[index-1].Format = n.cmd.Args[index].Format
		}
	}
	n.cmd.Args = append(n.cmd.Args[:index], n.cmd.Args[index+1:]...)
	return nil
}
