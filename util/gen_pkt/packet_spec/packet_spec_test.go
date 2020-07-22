package packetspec

import (
	"testing"

	"fmt"
	packet_spec "util/gen_pkt/packet_spec"

	"github.com/alecthomas/participle"
	"github.com/davecgh/go-spew/spew"
)

func parsePacketSpec(t *testing.T, packetSpec string) *packet_spec.PacketSpec {
	parser, err := participle.Build(&packet_spec.PacketSpec{})
	if err != nil {
		fmt.Printf("%v\n", err)
		t.Fail()
		// assert.NoError(t, err)
		return nil
	}

	spec := &packet_spec.PacketSpec{}
	err = parser.ParseString(packetSpec, spec)
	if err != nil {
		fmt.Printf("parse error: %v\n", err)
		t.Fail()
		// assert.NoError(t, err)
		return nil
	}

	return spec
}

// func TestSingleSimplePacket(t *testing.T) {
// 	spec := parsePacketSpec(t, `
// packet Test {
// 	int foo;
// 	int bar;
// }
// `)

// 	assert.Equal(t, 1, len(spec.Packet))
// 	assert.Equal(t, "Test", spec.Packet[0].Name)
// 	assert.Equal(t, 2, len(spec.Packet[0].Fields))
// 	assert.Equal(t, "int", spec.Packet[0].Fields[0].Type)
// 	assert.Equal(t, "foo", spec.Packet[0].Fields[0].Name)
// 	assert.Equal(t, "int", spec.Packet[0].Fields[1].Type)
// 	assert.Equal(t, "bar", spec.Packet[0].Fields[1].Name)
// }

func TestSingleComplexPacket(t *testing.T) {
	spec := parsePacketSpec(t, `
enum Error {
	OK = 0x00
	FAILED_1 = 0x01
	FAILED_2 = 0x02
	BANNED = 0x03
	UNKNOWN_ACCOUNT = 0x04
	UNKNOWN_ACCOUNT_3 = 0x05
	ALREADYONLINE = 0x06
	NOTIME = 0x07
	DBBUSY = 0x08
	BADVERSION = 0x09
	DOWNLOAD_FILE = 0x0A
	FAILED_3 = 0x0B
	SUSPENDED = 0x0C
	FAILED_4 = 0x0D
	CONNECTED = 0x0E
	PARENTAL_CONTROL = 0x0F
	LOCKED_ENFORCED = 0x10
}

packet ClientLoginChallenge {
	string[4] game_name
	int8 version[3]
	int16 build
	string[4] platform
	string[4] os
	string[4] locale
	int32 timezone_offset
	int32b ip_address
	string account_name
}

packet ServerLoginChallenge {
    int8 unk
	Error error

	struct challenge {
		bigint[32] B
		int8 g_len
		int8 g
		int8 N_len
		bigint[32] N
		bigint[32] salt
		bigint[16] crc_salt
		int8 unk = 0
	} if (error is Error.OK)
}
`)

	// assert.Equal(t, 1, len(spec.Packets))
	// assert.Equal(t, "UpdateObject", spec.Packets[0].Name)

	spew.Dump(spec)
	t.Fail()

}
