# Snowflake [![GoDoc](https://godoc.org/github.com/andersfylling/snowflake?status.svg)](https://godoc.org/github.com/andersfylling/snowflake) [![CircleCI](https://circleci.com/gh/andersfylling/snowflake/tree/master.svg?style=shield)](https://circleci.com/gh/andersfylling/snowflake/tree/master) [![Go Report Card](https://goreportcard.com/badge/github.com/andersfylling/snowflake)](https://goreportcard.com/report/github.com/andersfylling/snowflake) [![Test Coverage](https://api.codeclimate.com/v1/badges/5fe3da7a87c3185b5f33/test_coverage)](https://codeclimate.com/github/andersfylling/snowflake/test_coverage)
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)

Does not hold functionality to connect a snowflake service, nor generating snowflakes. But rather parsing the snowflakes from JSON only. The default behaviour is to parse Discord snowflakes, but build constraints exists to handle twitter snowflakes as well. Note that there also exist `DateByEpoch` to which you can pass any desired epoch to affect the date parsing.

For module usage you must utilise the suffix found in the go.mod file (/v2 for v2.x.x releases, /v3 for v3.x.x releases, etc.).

Usage:
>Note: if you are against dot imports, which I can understand, instead of writing snowflake.Snowflake, you can write snowflake.ID. ID is just an alias for Snowflake.

```go
import . "github.com/andersfylling/snowflake"

type DiscordRole struct {
    ID          Snowflake    `json:"id"`
    Name        string       `json:"name"`
    Managed     bool         `json:"managed"`
    Mentionable bool         `json:"mentionable"`
    Hoist       bool         `json:"hoist"`
    Color       int          `json:"color"`
    Position    int          `json:"position"`
    Permissions uint64       `json:"permissions"`
}
```

If you're creating an API that sends JSON to a multiple different language clients, some might not be able to process uint64, such as javascript. To support both uint64 and string use the JSON struct included:

```go
import . "github.com/andersfylling/snowflake"

type DiscordRole struct {
    *SnowflakeJSON           `json:"snowflake"`
    Name        string       `json:"name"`
    Managed     bool         `json:"managed"`
    Mentionable bool         `json:"mentionable"`
    Hoist       bool         `json:"hoist"`
    Color       int          `json:"color"`
    Position    int          `json:"position"`
    Permissions uint64       `json:"permissions"`
}
```

This adds two fields: `ID` and `IDStr`. Where the first is of a snowflake.ID(uint64), and the second is a string. This creates the JSON format (IDs only. Where the dots represents the remaining DiscordRole fields):

```json
{
    "snowflake": {
          "id": 74895735435643,
          "id_str": "74895735435643",
    }
}
```

Now an alternative is to send only the string version by adding `,string` to the json tag. Since Snowflake utilizes a custom Marshaler, this won't function. if you want to support 32bit languages, you must use the SnowflakeJSON as mentioned above.
You can also extract a SnowflakeJSON from a Snowflake by calling `Snowflake.JSONStruct()`.

### Build constraints
if you want the Snowflake.Date() method to parse snowflakes based on the twitter epoch, you can do `go build -tags=snowflake_twitter`. The default behaviour will use the Discord epoch.


### Benchmarks

```markdown
BenchmarkUnmarshalJSON/string-8                        50000000	        24.70 ns/op
BenchmarkUnmarshalJSON/uint64-a-8                     300000000	         4.17 ns/op
BenchmarkUnmarshalJSON/uint64-b-8                      20000000	        77.30 ns/op

BenchmarkUnmarshalJSON/string-struct-8                  3000000	       500.00 ns/op
BenchmarkUnmarshalJSON/snowflake-struct-8               3000000	       476.00 ns/op

BenchmarkUnmarshal_snowflakeStrategies/dereference-8  100000000	        15.40 ns/op
BenchmarkUnmarshal_snowflakeStrategies/tmp-var-8      100000000	        11.00 ns/op
```
In the first 3 tests, I test out the convertion strategy directly.

 1. string: byte slice is copied by calling `dst = string(src)`
 2. uint64-a: a custom converted (loop) is used
 3. uint64-b: strconv.ParseUint is used

In the next 2 tests, I call json.Unmarshal to see how it will perform in real life.

 4. a struct with `ID string 'json:"snowflake"'`
 5. a struct with `ID Snowflake 'json:"snowflake"'`, which utilises the custom loop found in #2

These last tests simply regards optimazations of the current custom loop, just as removing dereference. (The UnmarshalJSON method is called here which is why it slower than #2).

 6. using dereference to update the snowflake during loop
 7. using a local var during the loop, then updating the snowflake when finishing with no error