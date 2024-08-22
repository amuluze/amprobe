// Package nginx
// Date: 2024/8/6 13:55
// Author: Amu
// Description:
package nginx

type Instruction interface {
	Id() int64
	Name() string
	Rename(string)
	Command() *NodeCommand
	
	GetProperties() []string
	
	AddProperty(...string) error
	DelProperty(...string) error
	ClearProperty()
	
	GetKeywordProperties() map[string]string
	
	GetKeywordProperty(string) string
	SetKeywordProperty(string, string) error
	DelKeywordProperty(string) error
	
	BlockInstructions() []Instruction
	BlockInstructionMapping() (map[string][]int, map[int64]int)
	
	AddInstruction(...Instruction) error
	DelInstruction(...Instruction) error
	MoveInstruction(instruction Instruction, pos int, absolute bool) error
	
	FindSubByName(string) ([]Instruction, error)
	FindSubById(id int64) (Instruction, error)
}
