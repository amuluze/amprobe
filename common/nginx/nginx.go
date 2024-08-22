// Package nginx
// Date: 2024/8/6 13:44
// Author: Amu
// Description:
package nginx

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// populate format for AST

func Populate(i Instruction) error {
	if i.Command().Block == nil {
		n := len(i.Command().Args)
		if n == 0 {
			return fmt.Errorf("invalid")
		}
		if i.Command().Args[n-1].Type != ArgTypeCmdFin {
			i.Command().Args = append(i.Command().Args,
				&NodeArg{
					Type: ArgTypeCmdFin,
					Text: ";",
					Format: &NodeFormatter{
						Type: FormatterTypeNewLine,
						Text: "\n",
					},
				},
			)
		}
	} else {
		if i.Command().Block.BlkBgn == nil {
			i.Command().Block.BlkBgn = &NodeArg{
				Type: ArgTypeBlkBgn,
				Text: "{",
				Format: &NodeFormatter{
					Type: FormatterTypeNewLine,
					Text: "\n",
				},
			}
		}
		if i.Command().Block.BlkEnd == nil {
			i.Command().Block.BlkEnd = &NodeArg{
				Type: ArgTypeBlkEnd,
				Text: "}",
				Format: &NodeFormatter{
					Type: FormatterTypeNewLine,
					Text: "\n",
				},
			}
		}

		// fix indent
		for _, ins := range i.BlockInstructions() {
			ins.Command().Indent = i.Command().Indent + 1
			if err := Populate(ins); err != nil { // recursive
				return err
			}
		}
	}
	return nil
}

func PrintCommand(c *NodeCommand) (string, error) {
	s, _, err := printCommand(c, true)
	if err != nil {
		return "", err
	}
	return s, nil
}

func printCommand(c *NodeCommand, first bool) (string, bool, error) {
	var ret string

	argPrefix := func(indent, index int) string {
		if index != 0 {
			indent++
		}
		return strings.Repeat(IndentString, indent)
	}

	// args
	for i, arg := range c.Args {
		switch arg.Type {
		case ArgTypeIdent:
			if first {
				ret += argPrefix(c.Indent, i) + arg.Text
			} else {
				ret += " " + arg.Text
			}
		case ArgTypeString:
			if first {
				ret += argPrefix(c.Indent, i) + arg.Text
			} else {
				ret += " " + arg.Text
			}
		case ArgTypeCmdFin:
			if first {
				ret += argPrefix(c.Indent, i) + arg.Text
			} else {
				ret += arg.Text
			}
		case ArgTypeFormat:
			// pass
		case ArgTypeBlkBgn, ArgTypeBlkEnd:
			return "", false, fmt.Errorf("block braces appears in command arguments")
		}

		if arg.Format != nil {
			if arg.Type == ArgTypeFormat {
				// standalone formatter
				switch arg.Format.Type {
				case FormatterTypeNewLine:
					ret += arg.Format.Text // new line does not need prefix
				case FormatterTypeComment:
					ret += argPrefix(c.Indent, i) + arg.Format.Text
				}
			} else {
				// trailing formatter
				switch arg.Format.Type {
				case FormatterTypeNewLine:
					ret += arg.Format.Text
				case FormatterTypeComment:
					ret += " " + arg.Format.Text
				}
			}
		}

		first = arg.Format != nil
	}

	// block
	if c.Block != nil {
		// block begin
		if c.Block.BlkBgn == nil {
			return "", false, errors.New("block has no beginning")
		}
		if first {
			ret += argPrefix(c.Indent, 0) + c.Block.BlkBgn.Text
		} else {
			ret += " " + c.Block.BlkBgn.Text
		}
		if c.Block.BlkBgn.Format != nil {
			// trailing formatter
			switch c.Block.BlkBgn.Format.Type {
			case FormatterTypeNewLine:
				ret += c.Block.BlkBgn.Format.Text
				first = true
			case FormatterTypeComment:
				ret += " " + c.Block.BlkBgn.Format.Text
				first = true
			}
		}

		// block commands
		for _, cmd := range c.Block.Commands {
			s, f, err := printCommand(cmd, first)
			if err != nil {
				return "", false, err
			}
			ret += s
			first = f
		}

		// block end
		if c.Block.BlkEnd == nil {
			return "", false, fmt.Errorf("block has no end")
		}
		if first {
			ret += argPrefix(c.Indent, 0) + c.Block.BlkEnd.Text
		} else {
			ret += " " + c.Block.BlkEnd.Text
		}
		if c.Block.BlkEnd.Format != nil {
			// trailing formatter
			switch c.Block.BlkEnd.Format.Type {
			case FormatterTypeNewLine:
				ret += c.Block.BlkEnd.Format.Text
				first = true
			case FormatterTypeComment:
				ret += " " + c.Block.BlkEnd.Format.Text
				first = true
			}
		}
	}

	// ok
	return ret, first, nil
}

func GenerateProxyNames(config *ProxyConfig) map[string]string {
	return map[string]string{"proxy_pass": "proxy_pass",
		"proxy_bind":            "proxy_bind",
		"proxy_mark":            "proxy_mark",
		"proxy_set_header":      "proxy_set_header",
		"proxy_hide_header":     "proxy_hide_header",
		"proxy_ssl_server_name": "proxy_ssl_server_name",
		"proxy_ssl_name":        "proxy_ssl_name",
		"proxy_ssl_gm":          "proxy_ssl_gm",
		"proxy_ssl_ciphers":     "proxy_ssl_ciphers",
		"scheme":                config.Schema,
	}
}

func WriteConfig(resourcesDir string, siteGroup SiteGroup) error {
	g := NewGenerator()
	config, err := g.Generate(siteGroup)
	if err != nil {
		return err
	}
	configFile := filepath.Join(resourcesDir, "server.conf")
	fmt.Printf("configFile: %s\n", configFile)
	err = os.WriteFile(configFile, []byte(config.SiteConfig["siteConfig"]), 0644)
	if err != nil {
		return err
	}
	return nil
}
