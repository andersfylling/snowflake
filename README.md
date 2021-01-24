# Snowflake [![GoDoc](https://godoc.org/github.com/andersfylling/snowflake?status.svg)](https://godoc.org/github.com/andersfylling/snowflake) [![codecov](https://codecov.io/gh/andersfylling/snowflake/branch/master/graph/badge.svg?token=w5FS4B9fou)](https://codecov.io/gh/andersfylling/snowflake) [![Go Report Card](https://goreportcard.com/badge/github.com/andersfylling/snowflake)](https://goreportcard.com/report/github.com/andersfylling/snowflake)
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)

> This lib supports signed numbers, but will convert them to uint64.

Does not hold functionality to connect a snowflake service, nor generating snowflakes. But supports parsing snowflakes sent by Discord (both integer and string), and renders zero value as "null" and higher values as a string version to comply with Discord.

For module usage you must utilise the suffix found in the go.mod file (/v2 for v2.x.x releases, /v3 for v3.x.x releases, etc.).

Usage:
>Note: if you are against dot imports, which I can understand, instead of writing snowflake.Snowflake, you can write snowflake.ID. ID is just an alias for Snowflake.

```go
import . "github.com/andersfylling/snowflake/v5"

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

### Benchmarks

```markdown
name                              time/op
UnmarshalJSON/string-8            52.5ns ±14%
UnmarshalJSON/snowflake-8         23.0ns ± 6%
UnmarshalJSON/string-struct-8     1.23µs ± 5%
UnmarshalJSON/snowflake-struct-8  1.15µs ±10%
NullCheck/string-8                0.58ns ± 8%
```