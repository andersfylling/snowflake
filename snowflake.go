package snowflake

import (
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
	if data[0] == '"' {
		if length == 1 {
			// This can't be anything.
			return
		}
		dataRemainder := make([]byte, 0, length - 2)
		if data[1] == '-' {
			// Negative value.
			*s |= 1 << 63
		} else {
			// Handle the first byte.
			dataRemainder = append(dataRemainder, data[1])
		}
		for i := 2; i < length; i++ {
			switch x := data[i]; x {
			case '"':
				// End of string.
				break
			default:
				// Add to remainder.
				dataRemainder = append(dataRemainder, x)
			}
		}
		x, err := ParseSnowflakeUint(string(dataRemainder), 10)
		if err != nil {
			return err
		}
		*s |= x
	} else {
		// Take the yolo strategy and try and parse un-compliant JSON.
		var dataRemainder []byte
		if data[0] == '-' {
			// Negative value.
			dataRemainder = []byte{}
			*s |= 1 << 63
		} else {
			// Start of value.
			dataRemainder = []byte{data[0]}
		}
		dataRemainder = append(dataRemainder, data[1:]...)
		x, err := ParseSnowflakeUint(string(dataRemainder), 10)
		if err != nil {
			return err
		}
		*s |= x
	}
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

func (s Snowflake) DateByEpoch(epoch uint64) time.Time {
	date := (uint64(s) >> uint64(22)) + epoch
	return time.Unix(int64(date), 0)
}
