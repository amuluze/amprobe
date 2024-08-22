// Package nginx
// Date: 2024/8/6 13:59
// Author: Amu
// Description:
package nginx

/* AST */

const (
	IndentString   = "    "
	GeneralSep     = "_"
	UpstreamPrefix = "stream"
)

const (
	ArgTypeIdent  = iota // identifier
	ArgTypeString        // string, e.g., 'aaa', "bbb"
	ArgTypeBlkBgn        // {
	ArgTypeBlkEnd        // }
	ArgTypeCmdFin        // ;
	ArgTypeFormat        // just newline or comment
)

const (
	FormatterTypeNewLine = iota
	FormatterTypeComment
)

const (
	InstructionNameUnknown = "unknown"
)
