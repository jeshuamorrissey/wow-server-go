package packetspec

// PacketSpec represents the basic structure of a packet spec file.
// It is made up of one or more enum or packet.
type PacketSpec struct {
	Packets []*Packet `@@*`
}

type Enum struct {
	Name   string       `"enum" @Ident`
	Values []*EnumValue `"{" ( @@ )* "}"`
}

type EnumValue struct {
	Key   string `@Ident`
	Value string `"=" @Ident`
}

// Packet defines the syntax for a packet specification. Packets are
// the word "packet" followed by a name followed by the definition in
// {} brackets.
type Packet struct {
	Name    string         `"packet" @Ident`
	Entries []*PacketEntry `"{" ( @@ )* "}"`
}

// PacketEntry represents a single logical thing within a packet.
// This could be a basic field, and inline struct definition or a
// conditional.
type PacketEntry struct {
	IfConditional *IfConditional `"if" @@`
	Struct        *Struct        `| "struct" @@`
	Field         *Field         `| @@`
}

type IfConditional struct {
	Var1      string         `"(" @Ident( @"." @Ident )?`
	Operation string         `@Ident`
	Var2      string         `@Ident( @"." @Ident )? ")"`
	Entries   []*PacketEntry `"{" ( @@ )* "}"`
}

type Struct struct {
	Name    string         `@Ident`
	Entries []*PacketEntry `"{" ( @@ )* "}"`
}

type Field struct {
	Type     *Type    `@@`
	Name     string   `@Ident`
	ArrayLen int      `( "[" @Int "]" )?`
	Default  *Default `( "=" @@ )?`
}

type Default struct {
	String string  `@String`
	Int    int     `| @Int`
	Float  float64 `| @Float`
}

type Type struct {
	Name string `@Ident`
	Size int    `( "[" @Int "]" )?`
}
