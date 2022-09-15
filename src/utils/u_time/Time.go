package u_time

import (
	"time"
)

const (
	FMT_YYYY_MM_DD          string = `2006-01-02`
	FMT_YYYY_MM_DD_HH_MM_SS string = `2006-01-02 15:04:05`
)

func FmtYmdUnix(timeUnix int64) string {
	return time.Unix(timeUnix, 0).Format(FMT_YYYY_MM_DD)
}
