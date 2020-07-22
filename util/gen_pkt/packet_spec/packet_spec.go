package packetspec

// PacketSpec represents the basic structure of a packet spec file.
// It is made up of one or more enum or packet.
type PacketSpec struct {
	Specs []*Specs `@@*`
}

type Specs struct {
	Packet *Packet `"packet" @@`
	Enum   *Enum   `| "enum" @@`
}

// Enum represents a block of named constants.
type Enum struct {
	Type   string       `@Ident`
	Name   string       `@Ident`
	Values []*EnumValue `"{" ( @@ )* "}"`
}

// EnumValue represents a single value line within an enum.
type EnumValue struct {
	Key   string `@Ident "="`
	Value int    `@Int`
}

// Packet defines the syntax for a packet specification. Packets are
// the word "packet" followed by a name followed by the definition in
// {} brackets.
type Packet struct {
	Name    string         `@Ident`
	Entries []*PacketEntry `"{" ( @@ )* "}"`
}

// PacketEntry represents a single logical thing within a packet.
// This could be a basic field, and inline struct definition or a
// conditional.
type PacketEntry struct {
	Struct *Struct `"struct" @@`
	Field  *Field  `| @@`
}

// Struct represents a sub-structure within a packet. It is essential
// the same as a packet, but can have an attached condition.
type Struct struct {
	Name          string         `@Ident`
	Entries       []*PacketEntry `"{" ( @@ )* "}"`
	IfConditional *IfConditional `( "if" "(" @@ ")" )?`
}

// Field represents the most-basic field level of information. This specifies
// the type, length, array length, default and condition for a given field.
type Field struct {
	Type          string         `@Ident`
	Size          int            `( "[" @Int "]" )?`
	Name          string         `@Ident`
	ArrayLen      int            `( "[" @Int "]" )?`
	Default       *Default       `( "=" @@ )?`
	IfConditional *IfConditional `( "if" "(" @@ ")" )?`
}

// IfConditional represents the LHS and RHS of a basic if conditional.
type IfConditional struct {
	Var1      string `@Ident( @"." @Ident )?`
	Operation string `@Ident`
	Var2      string `@Ident( @"." @Ident )?`
}

// Default represents a parsed default value (which is either a string, int
// or float). If the field is not specified, this default will be used.
type Default struct {
	String string  `@String`
	Int    int     `| @Int`
	Float  float64 `| @Float`
}
