package snowflake

import (
	"errors"
	"fmt"
	"strconv"
)

// ParseSnowflakeString interprets a string with a decimal number.
//         Note that in contrast to ParseUint, this function assumes the given string is
//         always valid and thus will panic rather than return an error.
//         This should only be used on checks that can be done at compile time,
//         unless you want to trust other modules to returns valid data.
func ParseSnowflakeString(v string) Snowflake {
	id, err := ParseSnowflakeUint(v, 10)
	if err != nil {
		panic(fmt.Errorf("unable to parse %s into snowflake", v)) // TODO
	}
	return id
}

// ParseUint converts a string and given base to a Snowflake
func ParseSnowflakeUint(v string, base int) (Snowflake, error) {
	if v == "" {
		return Snowflake(0), nil
	}

	id, err := strconv.ParseUint(v, base, 64)
	return Snowflake(id), err
}

func GetSnowflake(v interface{}) (s Snowflake, err error) {
	switch x := v.(type) {
	case int:
		s = NewSnowflake(uint64(x))
	case int8:
		s = NewSnowflake(uint64(x))
	case int16:
		s = NewSnowflake(uint64(x))
	case int32:
		s = NewSnowflake(uint64(x))
	case int64:
		s = NewSnowflake(uint64(x))

	case uint:
		s = NewSnowflake(uint64(x))
	case uint8:
		s = NewSnowflake(uint64(x))
	case uint16:
		s = NewSnowflake(uint64(x))
	case uint32:
		s = NewSnowflake(uint64(x))
	case uint64:
		s = NewSnowflake(x)

	case string:
		i, err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			s = NewSnowflake(0)
		} else {
			s = NewSnowflake(i)
		}

	case Snowflake:
		s = x

	default:
		s = NewSnowflake(0)
		err = errors.New("not supported type for snowflake")
	}

	return
}
