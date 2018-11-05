// +build !snowflake_twitter

package snowflake

import "time"

func (s Snowflake) Date() time.Time {
	var epoch uint64 = (uint64(s) >> uint64(22)) + EpochDiscord
	return time.Unix(int64(epoch), 0)
}
