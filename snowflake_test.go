package snowflake

import "testing"

func TestString(t *testing.T) {
	id := NewID(uint64(435834986943))
	if id.String() != "435834986943" {
		t.Errorf("String conversion failed. Got %s, wants %s", id.String(), "435834986943")
	}

	if id.JSONStruct().IDStr != "435834986943" {
		t.Errorf("String conversion failed. Got %s, wants %s", id.String(), "435834986943")
	}

	if id.HexString() != "6579ca21bf" {
		t.Errorf("String conversion for Hex failed. Got %s, wants %s", id.HexString(), "6579ca21bf")
	}

	if id.HexPrettyString() != "0x6579ca21bf" {
		t.Errorf("String conversion for Pretty Hex failed. Got %s, wants %s", id.HexPrettyString(), "0x6579ca21bf")
	}

	if id.BinaryString() != "110010101111001110010100010000110111111" {
		t.Errorf("String conversion for binary failed. Got %s, wants %s", id.BinaryString(), "110010101111001110010100010000110111111")
	}
}

func TestEmpty(t *testing.T) {
	id := NewID(0)
	if !id.Empty() {
		t.Errorf("Expects ID to be viewed as empty when value is 0")
	}
}
