package main

import (
	"fmt"
	"os"
	"strings"
	packet_spec "util/gen_pkt/packet_spec"

	"github.com/achiku/varfmt"
	"github.com/alecthomas/participle"
	"github.com/dave/jennifer/jen"
)

var typeStrToJenType = map[string]func() *jen.Statement{
	"int8":   jen.Int8,
	"int16":  jen.Int16,
	"int32":  jen.Int32,
	"int32b": jen.Int32,
	"string": jen.String,
	"bigint": func() *jen.Statement { return jen.Add(jen.Op("*"), jen.Qual("math/big", "Int")) },
}

func parseEnumValue(name, val string) string {
	return varfmt.PublicVarName(name + "_" + strings.ToLower(val))
}

func parseEnumType(name string) string {
	return varfmt.PublicVarName(name)
}

func loadEntries(entries []*packet_spec.PacketEntry) ([]jen.Code, error) {
	structFields := []jen.Code{}
	for _, entry := range entries {
		if entry.Field != nil {
			fieldName := varfmt.PublicVarName(entry.Field.Name)
			fieldTypeFn, ok := typeStrToJenType[entry.Field.Type]
			var fieldType jen.Code
			if ok {
				fieldType = fieldTypeFn()
			} else {
				fieldType = jen.Id(parseEnumType(entry.Field.Type))
			}

			// If there is a condition, the field is optional.
			if entry.Field.IfConditional != nil {
				fieldType = jen.Add(jen.Op("*"), fieldType)
			}

			structFields = append(structFields, jen.Add(jen.Id(fieldName), fieldType))
		} else if entry.Struct != nil {
			subStructFields, err := loadEntries(entry.Struct.Entries)
			if err != nil {
				return nil, err
			}

			// Add sub-struct. If there is a condition, the struct is optional.
			subStruct := jen.Struct(subStructFields...)
			if entry.Struct.IfConditional != nil {
				subStruct = jen.Add(jen.Op("*"), subStruct)
			}

			structFields = append(structFields, jen.Add(jen.Id(entry.Struct.Name), subStruct))
		}
	}

	return structFields, nil
}

func genGoEncodeEntries(entries []*packet_spec.PacketEntry) ([]jen.Code, error) {
	statements := []jen.Code{}

	for _, entry := range entries {
		if entry.Field != nil {
			// switch based on type
		} else if entry.Struct != nil {
			structStatements, err := genGoEncodeEntries(entry.Struct.Entries)
			if err != nil {
				return nil, err
			}

			statements = append(statements, structStatements...)
		}
	}

	return statements, nil
}

func genGoEncodeFn(packet *packet_spec.Packet, goFile *jen.File) error {
	// Generate a list of code statements to write the packet to a binary buffer.
	statements := []jen.Code{
		jen.Id("buffer").Op(":=").New(jen.Qual("bytes", "Buffer")),
	}

	encodeStatements, err := genGoEncodeEntries(packet.Entries)
	if err != nil {
		return err
	}

	// Return statement.
	statements = append(statements, encodeStatements...)
	statements = append(statements, jen.Return(jen.Id("buffer").Dot("Bytes").Call(), jen.Nil()))

	fn := goFile.Func()
	fn.Params(jen.Id("pkt").Op("*").Id(packet.Name)) // the method params list
	fn.Id("Encode")                                  // the method name
	fn.Params()                                      // the method arguments (none)
	fn.Params(jen.Index().Byte(), jen.Error())       // the method return values
	fn.Block(statements...)                          // the method statements

	return nil
}

func genGoFile(parser *participle.Parser, packetFilePath string) error {
	f, err := os.Open(packetFilePath)
	if err != nil {
		return err
	}

	spec := &packet_spec.PacketSpec{}
	err = parser.Parse(f, spec)
	if err != nil {
		return err
	}

	goFile := jen.NewFile("packets")

	for _, s := range spec.Specs {
		if s.Enum != nil {
			enumTypeFn := typeStrToJenType[s.Enum.Type]
			enumTypeName := parseEnumType(s.Enum.Name)
			goFile.Type().Add(jen.Id(enumTypeName), enumTypeFn())

			for _, val := range s.Enum.Values {
				enumName := jen.Id(parseEnumValue(s.Enum.Name, val.Key))
				enumValue := jen.Lit(val.Value)
				goFile.Const().Add(jen.Id(enumTypeName), enumName, jen.Op("="), enumValue)
			}
		} else if s.Packet != nil {
			// Generate the struct.
			structFields, err := loadEntries(s.Packet.Entries)
			if err != nil {
				return err
			}

			// Add the struct type.
			goFile.Type().Add(jen.Id(s.Packet.Name), jen.Struct(structFields...))

			// Add a method to create a new packet. This should also set defaults (if appropriate).
			err = genGoEncodeFn(s.Packet, goFile)
			if err != nil {
				return err
			}

			// Add a function to take the struct and produce a binary blob.

			// Add a function to take a binary blob and produce a struct.
			// TODO
		}
	}

	fmt.Printf("%#v", goFile)

	return nil
}

func main() {
	// Build the parser.
	parser, err := participle.Build(&packet_spec.PacketSpec{})
	if err != nil {
		fmt.Printf("Failed to build parser: %v\n", err)
		return
	}

	err = genGoFile(parser, "/home/jeshua/code/wow-server-go/login_server/packets/login_challenge.packet")
	if err != nil {
		fmt.Printf("Failed to generate go file: %v\n", err)
	}
}
