package packetspec

import (
	"testing"

	packet_spec "util/gen_pkt/packet_spec"

	"github.com/alecthomas/participle"
	"github.com/stretchr/testify/assert"
)

func parsePacketSpec(t *testing.T, packetSpec string) *packet_spec.PacketSpec {
	parser, err := participle.Build(&packet_spec.PacketSpec{})
	if err != nil {
		assert.NoError(t, err)
		return nil
	}

	spec := &packet_spec.PacketSpec{}
	err = parser.ParseString(packetSpec, spec)
	if err != nil {
		assert.NoError(t, err)
		return nil
	}

	return spec
}

func TestSingleSimplePacket(t *testing.T) {
	spec := parsePacketSpec(t, `
packet Test {
	int foo;
	int bar;
}	
`)

	assert.Equal(t, 1, len(spec.Packet))
	assert.Equal(t, "Test", spec.Packet[0].Name)
	assert.Equal(t, 2, len(spec.Packet[0].Fields))
	assert.Equal(t, "int", spec.Packet[0].Fields[0].Type)
	assert.Equal(t, "foo", spec.Packet[0].Fields[0].Name)
	assert.Equal(t, "int", spec.Packet[0].Fields[1].Type)
	assert.Equal(t, "bar", spec.Packet[0].Fields[1].Name)
}
