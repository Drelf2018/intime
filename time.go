package intime

import (
	"strconv"
	"time"

	"github.com/itchyny/timefmt-go"
)

const DefaultFormat string = "%Y-%m-%d %H:%M:%S"

type Time int64

func (t *Time) Add(x any) error {
	i, err := ParseInt(x)
	*t += Time(i)
	return err
}

func (t *Time) Set(x any, mode ...string) error {
	i, err := ParseInt(x)
	*t = Time(i)
	return err
}

func (t *Time) Now() {
	*t = ConvTime(time.Now())
}

func (t Time) Time() time.Time {
	return time.Unix(int64(t/1000), int64(t%1000)*int64(time.Millisecond))
}

func (t Time) String() string {
	return strconv.FormatInt(int64(t), 10)
}

func (t Time) Unix() int64 {
	return int64(t / 1000)
}

func (t Time) UnixMilli() int64 {
	return int64(t)
}

func (t Time) UnixNano() int64 {
	return int64(t * 1_000_000)
}

func (t Time) IsZero() bool {
	return t == 0
}

func (t Time) Format(format ...string) string {
	if len(format) != 0 {
		return timefmt.Format(t.Time(), format[0])
	}
	return timefmt.Format(t.Time(), DefaultFormat)
}

var _ UnixTimestamp = Time(0)
