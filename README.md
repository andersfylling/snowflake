# Snowflake [![GoDoc](https://godoc.org/github.com/andersfylling/snowflake?status.svg)](https://godoc.org/github.com/andersfylling/snowflake) [![codecov](https://codecov.io/gh/andersfylling/snowflake/branch/master/graph/badge.svg?token=w5FS4B9fou)](https://codecov.io/gh/andersfylling/snowflake) [![Go Report Card](https://goreportcard.com/badge/github.com/andersfylling/snowflake)](https://goreportcard.com/report/github.com/andersfylling/snowflake)
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)

> This lib supports signed numbers, but will convert them to uint64.

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

If you're creating an API that sends JSON, the snowflake will automatically be converted to/from a string for you.

### Build constraints
if you want the Snowflake.Date() method to parse snowflakes based on the twitter epoch, you can do `go build -tags=snowflake_twitter`. The default behaviour will use the Discord epoch.


### Benchmarks

```markdown
BenchmarkUnmarshalJSON/string-4            	40553596	        27.8 ns/op
BenchmarkUnmarshalJSON/uint64-a-4          	195255220	        5.90 ns/op
BenchmarkUnmarshalJSON/uint64-b-4          	10915821	        92.0 ns/op

BenchmarkUnmarshalJSON/string-struct-4     	 1363028	       784 ns/op
BenchmarkUnmarshalJSON/snowflake-struct-4  	 1645940	       757 ns/op

BenchmarkUnmarshal_snowflakeStrategies/dereference-4         	59154159	        22.4 ns/op
BenchmarkUnmarshal_snowflakeStrategies/tmp-var-4             	72053302	        18.2 ns/op
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
