# snowflake

Does not hold functionality to connect a snowflake service, but rather dealing with the snowflake ID itself.

Usage:
```golang
import "github.com/andersfylling/snowflake"

type DiscordRole struct {
	ID          snowflake.ID `json:"id"`
	Name        string       `json:"name"`
	Managed     bool         `json:"managed"`
	Mentionable bool         `json:"mentionable"`
	Hoist       bool         `json:"hoist"`
	Color       int          `json:"color"`
	Position    int          `json:"position"`
	Permissions uint64       `json:"permissions"`
}
```

If you're creating an API that sends JSON to a javascript client, or any other language that can't process uint64. USe the JSON struct included:
```golang
import "github.com/andersfylling/snowflake"

type DiscordRole struct {
	*snowflake.JSON           `json:"snowflake"`
	Name        string       `json:"name"`
	Managed     bool         `json:"managed"`
	Mentionable bool         `json:"mentionable"`
	Hoist       bool         `json:"hoist"`
	Color       int          `json:"color"`
	Position    int          `json:"position"`
	Permissions uint64       `json:"permissions"`
}
```

This adds two fields: `ID` and `IDStr`. Where the first is of a snowflake.ID type, and the second is a string. This creates the JSON format (IDs only. Where the dots represents the remaining DiscordRole fields):
```json
{
	"snowflake": {
  		"id": 74895735435643,
  		"id_str": "74895735435643",
	},
	...
}
```

This does fulfill the twitter snowflake usecase described here: https://developer.twitter.com/en/docs/basics/twitter-ids
