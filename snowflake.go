package snowflake

import (
	"errors"
	"strconv"
	"time"
)

const (
	EpochDiscord uint64 = 1420070400000
)

// NewID creates a new Snowflake Snowflake from a uint64.
func NewSnowflake(id uint64) Snowflake {
	return Snowflake(id)
}

// Snowflake twitter snowflake design
type Snowflake uint64

type ID = Snowflake

// IsZero since snowflake exists of several parts, including a timestamp,
//       I assume a valid snowflake Snowflake is never 0.
func (s Snowflake) IsZero() bool {
	return uint64(s) == 0
}

// Valid makes sure the snowflake is after the fixed epoch
func (s Snowflake) Valid() bool {
	return (s >> 22) >= 1 // older than 1 millisecond
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
		// Blank value.
		return
	}
	if length == 4 && string(data) == "null" {
		// This is a zero value.
		return
	}

	// if the snowflake is passed as a string, we account for the double quote wrap
	start := 0
	if data[0] == '"' {
		start++
		length--
	}
	if signed := data[start] == '-'; signed {
		start++
		*s |= 1 << 63
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

	*s |= Snowflake(tmp)
	return
}

func (s Snowflake) MarshalJSON() (data []byte, err error) {
	if s == 0 {
		data = []byte(`null`)
	} else {
		data = []byte(`"` + s.String() + `"`)
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

func (s Snowflake) Date() time.Time {
	epoch := (uint64(s) >> uint64(22)) + EpochDiscord
	return time.Unix(int64(epoch)/1000, 0)
}
