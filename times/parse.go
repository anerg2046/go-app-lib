package times

import (
	"go-app/config"
	"time"
)

func ParseTime(strTime string) (t time.Time, err error) {
	var layouts = []string{
		"2006-01-02",
		"2006-01-02T15:04:05",
		"2006-01-02T15:04:05.000Z",
		"2006-01-02T15:04:05Z07:00",
	}
	for _, layout := range layouts {
		t, err = time.ParseInLocation(layout, strTime, config.APP.Timezone)
		if err == nil {
			return
		}
	}

	return
}
