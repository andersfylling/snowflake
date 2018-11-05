package snowflake

import (
	"errors"
	"strconv"
	"time"
)

const (
	EpochDiscord uint64 = 1420070400000
	EpochTwitter uint64 = 1288834974657
)

// NewID creates a new Snowflake Snowflake from a uint64.
func NewSnowflake(id uint64) Snowflake {
	return Snowflake(id)
}

// Snowflake twitter snowflake design
type Snowflake uint64

type ID = Snowflake

// JSON can be useful when sending the snowflake Snowflake by a json API
type SnowflakeJSON struct {
	ID    Snowflake `json:"id"`
	IDStr string    `json:"id_str"`
}

// Empty since snowflake exists of several parts, including a timestamp,
//       I assume a valid snowflake Snowflake is never 0.
func (s Snowflake) Empty() bool {
	return uint64(s) == 0
}

// JSONStruct returns a struct that can be embedded in other structs.
//            This is useful if you have a API server, since js can't parse uint64.
//            Therefore there must a snowflake Snowflake string.
func (s Snowflake) JSONStruct() *SnowflakeJSON {
	return &SnowflakeJSON{
		ID:    s,
		IDStr: `"` + s.String() + `"`,
	}
}

// String returns the decimal representation of the snowflake Snowflake.
func (s Snowflake) String() string {
	return strconv.FormatUint(uint64(s), 10)
}

// HexString converts the Snowflake into a hexadecimal string
func (s Snowflake) HexString() string {
	return strconv.FormatUint(uint64(s), 16)
}

// HexPrettyString converts the Snowflake into a hexadecimal string with the hex prefix 0x
func (s Snowflake) HexPrettyString() string {
	return "0x" + strconv.FormatUint(uint64(s), 16)
}

// MarshalBinary create a binary literal representation as a string
func (s Snowflake) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.FormatUint(uint64(s), 2)), nil
}

func (s *Snowflake) UnmarshalBinary(text []byte) (err error) {
	id, err := strconv.ParseUint(string(text), 2, 64)
	if err != nil {
		return
	}
	*s = Snowflake(id)

	return
}

func (s *Snowflake) UnmarshalJSON(data []byte) (err error) {
	*s = 0
	length := len(data)
	if length == 0 {
		return
	}

	// "id":null <- valid null
	// "id":"null" <- not a null
	// no need to check for the entire null word: if the first is a letter, we can't parse it anyways.
	// and since null, is never used in a string, "null", we can safely assume the n is the start of a null.
	first := data[0]
	if length < 6 && first == 'n' {
		return
	}

	// if the snowflake is passed as a string, we account for the double quote wrap
	start := 0
	if first == '"' {
		start++
		length--
	}

	var c byte
	var tmp uint64
	for i := start; i < length; i++ {
		c = data[i] - '0'
		if c < 0 || c > 9 {
			err = errors.New("cannot parse non-integer symbol:" + string(data[i]))
			return
		}
		tmp = tmp*10 + uint64(c)
	}

	*s = Snowflake(tmp)
	return
}

func (s Snowflake) MarshalJSON() (data []byte, err error) {
	// expect to have both "id" and "id_str"
	// but this can't be done, so the SnowflakeJSON type is provided as an alternative.
	if s == 0 {
		data = []byte(`null`)
	} else {
		data = []byte(s.String())
	}
	return
}

func (s Snowflake) MarshalText() (text []byte, err error) {
	text = []byte(s.String())
	err = nil

	return
}

func (s *Snowflake) UnmarshalText(text []byte) (err error) {
	id, err := strconv.ParseUint(string(text), 10, 64)
	if err != nil {
		return
	}
	*s = Snowflake(id)

	return
}

func (s Snowflake) DateByEpoch(epoch uint64) time.Time {
	date := (uint64(s) >> uint64(22)) + epoch
	return time.Unix(int64(date), 0)
}