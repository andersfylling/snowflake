package snowflake

import (
	"reflect"
	"testing"
)

func TestParseSnowflakeString(t *testing.T) {
	// test panic
	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()

	id := "435639843545"
	if ParseSnowflakeString(id).String() != id {
		t.Errorf("Incorrect string parsing for ID, base 10. Wants %s, got %s", id, ParseSnowflakeString(id).String())
	}
	ParseSnowflakeString("11")
	ParseSnowflakeString("0")
	ParseSnowflakeString("")
}

func TestParseSnowflakeStringWithPanicTriggerLetters(t *testing.T) {
	// test panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ParseID did not panic")
		}
	}()

	id := "435639sd843545gf453s"
	if ParseSnowflakeString(id).String() != id {
		t.Errorf("Incorrect string parsing for ID, base 10. Wants %s, got %s", id, ParseSnowflakeString(id).String())
	}
}

func TestParseSnowflakeStringWithPanicTriggerOverflow(t *testing.T) {
	// test panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ParseID did not panic")
		}
	}()

	id := "184467440737095516151" // string(uint64(0) - 1) + "1"
	if ParseSnowflakeString(id).String() != id {
		t.Errorf("Incorrect string parsing for ID, base 10. Wants %s, got %s", id, ParseSnowflakeString(id).String())
	}
}

func TestGetSnowflake(t *testing.T) {
	data := []interface{}{
		"348563",
		int(2345),
		int(-2345),
		uint(8345),
		int8(1),
		int16(1),
		int32(1),
		int64(1),
		uint8(1),
		uint16(1),
		uint32(1),
		uint64(1),
		Snowflake(3),
		Snowflake(246),
	}
	for _, v := range data {
		s, err := GetSnowflake(v)
		if err != nil || s.IsZero() {
			t.Error("cannot parse", reflect.TypeOf(v).Name(), "of value", v)
		}
	}

	type unknown struct {}
	s, err := GetSnowflake(unknown{})
	if s != 0 {
		t.Error("snowflake should not have a value")
	}
	if err == nil {
		t.Error("should fail to process a unknown type")
	}
}
