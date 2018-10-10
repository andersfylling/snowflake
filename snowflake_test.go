package snowflake

import (
	"encoding"
	"encoding/json"
	"strconv"
	"testing"
)

func TestString(t *testing.T) {
	var b []byte
	var err error
	id := NewSnowflake(uint64(435834986943))
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

	b, err = id.MarshalBinary()
	if err != nil {
		t.Error(err)
	}
	if string(b) != "110010101111001110010100010000110111111" {
		t.Errorf("String conversion for binary failed. Got %s, wants %s", string(b), "110010101111001110010100010000110111111")
	}
}

func TestEmpty(t *testing.T) {
	id := NewSnowflake(0)
	if !id.Empty() {
		t.Errorf("Expects ID to be viewed as empty when value is 0")
	}
}

func TestBinaryMarshalling(t *testing.T) {
	if _, ok := interface{}((*Snowflake)(nil)).(encoding.BinaryMarshaler); !ok {
		t.Error("does not implement encoding.BinaryMarshaler")
	}

	id := NewSnowflake(4598345)
	b, err := id.MarshalBinary()
	if err != nil {
		t.Error(err)
	}

	id2 := NewSnowflake(4534)
	err = id2.UnmarshalBinary(b)
	if err != nil {
		t.Error(err)
	}

	if id2 != id {
		t.Errorf("Value differs. Got %d, wants %d", id2, id)
	}
}

func TestTextMarshalling(t *testing.T) {
	if _, ok := interface{}((*Snowflake)(nil)).(encoding.TextMarshaler); !ok {
		t.Error("does not implement encoding.TextMarshaler")
	}

	target := "80351110224678912"

	id := NewSnowflake(4534)
	err := id.UnmarshalText([]byte(target))
	if err != nil {
		t.Error(err)
	}

	b, err := id.MarshalText()
	if err != nil {
		t.Error(err)
	}

	if string(b) != target {
		t.Errorf("Value differs. Got %s, wants %s", string(b), target)
	}
}

func TestJSONMarshalling(t *testing.T) {
	if _, ok := interface{}((*Snowflake)(nil)).(json.Marshaler); !ok {
		t.Error("does not implement json.Marshaler")
	}

	target := `"80351110224678912"`

	id := NewSnowflake(0)
	err := json.Unmarshal([]byte(target), &id)
	if err != nil {
		t.Error(err)
	}

	b, err := json.Marshal(id)
	if err != nil {
		t.Error(err)
	}

	if string(b) != target {
		t.Errorf("Incorrect snowflake value. Got %s, wants %s", string(b), target)
	}



	id = NewSnowflake(0)
	b, err = json.Marshal(&id)
	if err != nil {
		t.Error(err)
	}
	if string(b) != `null` {
		t.Error("expected 0 Snowflake to display as null")
	}
}

func TestDate(t *testing.T) {
	s := NewSnowflake(228846961774559232)
	if s.Date().Unix() != 1474631767458 {
		t.Error("date is incorrect")
	}
}

type testSet struct {
	result uint64 
	data []byte 
}

func TestSnowflake_UnmarshalJSON(t *testing.T) {
	if _, ok := interface{}((*Snowflake)(nil)).(json.Unmarshaler); !ok {
		t.Error("does not implement json.Unmarshaler")
	}
	
	data := []testSet{
		{8994537984753, []byte(`{"id":"8994537984753"}`)},
		{4573485, []byte(`{"id":"4573485"}`)},
		{2349872349, []byte(`{"id":"00002349872349"}`)},
		{435453, []byte(`{"id":"435453"}`)},
		{4987598525434463, []byte(`{"id":"4987598525434463"}`)},
		{59823042, []byte(`{"id":"059823042"}`)},
		{698734534634, []byte(`{"id":"698734534634"}`)},
		{24795873495, []byte(`{"id":"024795873495"}`)},
		{598360703000, []byte(`{"id":"0598360703000"}`)},
		{0, []byte(`{"id":"null"}`)},
		{0, []byte(`{"id":null}`)},
	}

	type Foo struct {
		Bar Snowflake `json:"id,omitempty"`
	}

	foo := &Foo{}
	for i := range data {
		wants := data[i].result
		input := data[i].data
		err := json.Unmarshal(input, foo)
		if err != nil {
			t.Error(err.Error() + " | " + "given input of: " + string(input))
		}

		if uint64(foo.Bar) != wants {
			t.Errorf("incorrect snowflake ID. Got %d, wants %d", foo.Bar, wants)
		}
	}

	// fail on letters in snowflake ID's
	evilJSON := []byte("{\"id\":\"89945379a84753\"}")
	err := json.Unmarshal(evilJSON, foo)
	if err == nil {
		t.Error("expected to fail, but continued to parse string as uint64")
	}
}

func BenchmarkSnowflake_UnmarshalJSON(b *testing.B) {
	dataSets := [][]byte{
		[]byte("\"8994537984753\""),
		[]byte("\"4573485\""),
		[]byte("\"00002349872349\""),
		[]byte("\"435453\""),
		[]byte("\"4987598525434463\""),
		[]byte("\"059823042\""),
		[]byte("\"698734534634\""),
		[]byte("\"024795873495\""),
		[]byte("\"0598360703000\""),
	}
	b.Run("string", func(b *testing.B) {
		var result string
		var i int
		length := len(dataSets)
		for n := 0; n < b.N; n++ {
			result = string(dataSets[i])
			if i == length {
				i = 0
			}
		}
		if result == "" {
		}
	})
	b.Run("uint64-a", func(b *testing.B) {
		var result uint64
		var i int
		lengthi := len(dataSets)
		for n := 0; n < b.N; n++ {
			data := dataSets[i]
			result = 0
			length := len(data) - 1
			for j := 1; j < length; j++ {
				result = result*10 + uint64(data[j]-'0')
			}
			if i == lengthi {
				i = 0
			}
		}
		if result == 0 {
		}
	})
	b.Run("uint64-b", func(b *testing.B) {
		var result uint64
		var i int
		length := len(dataSets)
		for n := 0; n < b.N; n++ {
			data := dataSets[i]
			var tmp uint64
			tmp, _ = strconv.ParseUint(string(data), 10, 64)
			result = uint64(tmp)
			if i == length {
				i = 0
			}
		}
		if result == 0 {
		}
	})
	type fooOld struct {
		Foo string `json:"id,omitempty"`
	}
	type fooNew struct {
		Foo Snowflake `json:"id,omitempty"`
	}
	dataSetsJSON := [][]byte{
		[]byte("{\"id\":\"8994537984753\"}"),
		[]byte("{\"id\":\"4573485\"}"),
		[]byte("{\"id\":\"00002349872349\"}"),
		[]byte("{\"id\":\"435453\"}"),
		[]byte("{\"id\":\"4987598525434463\"}"),
		[]byte("{\"id\":\"059823042\"}"),
		[]byte("{\"id\":\"698734534634\"}"),
		[]byte("{\"id\":\"024795873495\"}"),
		[]byte("{\"id\":\"0598360703000\"}"),
	}
	b.Run("string-struct", func(b *testing.B) {
		foo := &fooOld{}
		var i int
		length := len(dataSetsJSON)
		for n := 0; n < b.N; n++ {
			_ = json.Unmarshal(dataSetsJSON[i], foo)
			if i == length {
				i = 0
			}
		}
	})
	b.Run("snowflake-struct", func(b *testing.B) {
		foo := &fooNew{}
		var i int
		length := len(dataSetsJSON)
		for n := 0; n < b.N; n++ {
			_ = json.Unmarshal(dataSetsJSON[i], foo)
			if i == length {
				i = 0
			}
		}
	})
}
