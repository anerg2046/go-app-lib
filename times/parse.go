package times

import (
	"go-app/config"
	"time"
)

func ParseTime(strTime string) (t time.Time, err error) {
	var layouts = []string{
		time.DateOnly,
		time.DateTime,
		time.RFC3339,
		time.RFC3339Nano,
	}
	for _, layout := range layouts {
		t, err = time.ParseInLocation(layout, strTime, config.APP.Timezone)
		if err == nil {
			return
		}
	}

	return
}
