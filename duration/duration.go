package duration

import "time"

type Duration uint64

func (d Duration) ToTime() time.Duration {
	return time.Duration(d) * time.Second
}

const (
	Second Duration = 1
	Minute Duration = 60
	Hour   Duration = 60 * Minute
	Day    Duration = 24 * Hour
	Week   Duration = 7 * Day
	Month  Duration = 30 * Day
	Year   Duration = 365 * Day
)
