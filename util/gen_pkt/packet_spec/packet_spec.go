package packetspec

type PacketSpec struct {
	Packet []*Packet `@@*`
}

type Packet struct {
	Name   string   `"packet" @Ident`
	Fields []*Field `"{" { @@ } "}"`
}

type Field struct {
	Type string `@Ident`
	Name string `@Ident ";"`
}
