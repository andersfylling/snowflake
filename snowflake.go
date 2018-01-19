package snowflake

import (
	"strconv"
)

// NewID creates a new Snowflake ID from a uint64.
func NewID(id uint64) ID {
	return ID(id)
}

// Generate a snowflake ID.
func Generate() ID {
	return ID(0) // TODO
}

// ID Snowflake ID created by twitter
type ID uint64

// JSON can be useful when sending the snowflake ID by a json API
type JSON struct {
	ID    ID     `json:"id"`
	IDStr string `json:"id_str"`
}

// Empty since snowflake exists of several parts, including a timestamp,
//       I assume a valid snowflake ID is never 0.
func (s ID) Empty() bool {
	return uint64(s) == 0
}

// JSONStruct returns a struct that can be embedded in other structs.
//            This is useful if you have a API server, since js can't parse uint64.
//            Therefore there must a snowflake ID string.
func (s ID) JSONStruct() *JSON {
	return &JSON{
		ID:    s,
		IDStr: s.String(),
	}
}

// String returns the decimal representation of the snowflake ID.
func (s ID) String() string {
	return strconv.FormatUint(uint64(s), 10)
}

// HexString converts the ID into a hexadecimal string
func (s ID) HexString() string {
	return strconv.FormatUint(uint64(s), 16)
}

// HexPrettyString converts the ID into a hexadecimal string with the hex prefix 0x
func (s ID) HexPrettyString() string {
	return "0x" + strconv.FormatUint(uint64(s), 16)
}

// MarshalBinary create a binary literal representation as a string
func (s ID) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.FormatUint(uint64(s), 2)), nil
}
func (s *ID) UnmarshalBinary(text []byte) (err error) {
	id, err := strconv.ParseUint(string(text), 2, 64)
	if err != nil {
		return
	}
	*s = ID(id)

	return
}

func (s ID) MarshalText() (text []byte, err error) {
	text = []byte(s.String())
	err = nil

	return
}

func (s *ID) UnmarshalText(text []byte) (err error) {
	id, err := strconv.ParseUint(string(text), 10, 64)
	if err != nil {
		return
	}
	*s = ID(id)

	return
}
