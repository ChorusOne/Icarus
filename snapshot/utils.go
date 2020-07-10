package snapshot

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func UTCTime(timestampHex string) time.Time {
	timestampInt, _ := strconv.ParseInt(timestampHex, 0, 32)
	return time.Unix(timestampInt, 0).UTC()
}

func SameDateOfTimestamps(ts1, ts2 string) bool {

	date1 := UTCTime(ts1).Format(("02-01-2006"))
	date2 := UTCTime(ts2).Format(("02-01-2006"))
	return (date1 == date2)
}
