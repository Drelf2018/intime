package intime

import (
	"time"

	"github.com/Drelf2018/intime/errorf"
	"github.com/itchyny/timefmt-go"
)

const (
	ModeFormat string = "format"
)

var ErrModeFormat = errorf.New("intime: ModeFormat error: %v")
var CN = time.FixedZone("CST", 8*3600)

func Now() Time {
	return Time(time.Now().UnixNano() / 1_000_000)
}

func New(t any, mode ...string) Time {
	if len(mode) == 0 {
		return ParseTime(t)
	}
	switch mode[0] {
	case ModeFormat:
		var s string
		switch x := t.(type) {
		case []byte:
			s = string(x)
		case string:
			s = x
		default:
			panic(ErrModeFormat.Format(`t must be a string or []byte when using "format" mode`))
		}

		format := DefaultFormat
		if len(mode) > 1 {
			format = mode[1]
		}

		x, err := timefmt.Parse(s, format)
		if err != nil {
			panic(ErrModeFormat.Format(err))
		}
		return ConvTime(x)
	default:
		return Time(0)
	}
}
