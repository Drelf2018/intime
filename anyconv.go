package intime

import (
	"reflect"
	"strconv"
	"time"

	"github.com/Drelf2018/intime/errorf"
)

var ErrParseInt = errorf.New("intime: ParseInt of unknown type: %v")

func ConvTime(t time.Time) Time {
	return Time(t.UnixNano() / 1_000_000)
}

func ParseInt(t any) (i int64, err error) {
	if t == nil {
		return time.Now().UnixNano() / 1_000_000, nil
	}
	switch x := t.(type) {
	case []byte:
		return strconv.ParseInt(string(x), 10, 64)
	case string:
		return strconv.ParseInt(x, 10, 64)
	case int:
		i = int64(x)
	case int8:
		i = int64(x)
	case int16:
		i = int64(x)
	case int32:
		i = int64(x)
	case int64:
		i = int64(x)
	case uint:
		i = int64(x)
	case uint8:
		i = int64(x)
	case uint16:
		i = int64(x)
	case uint32:
		i = int64(x)
	case uint64:
		i = int64(x)
	case float32:
		i = int64(x)
	case float64:
		i = int64(x)
	case time.Duration:
		i = int64(x)
	case MilliTimestamp:
		i = x.UnixMilli()
	case NanoTimestamp:
		i = x.UnixNano() / 1_000_000
	case UnixTimestamp:
		i = x.Unix() * 1000
	default:
		err = ErrParseInt.Format(reflect.TypeOf(x).String())
	}
	return
}

func ParseTime(t any) Time {
	i, err := ParseInt(t)
	if err != nil {
		panic(err)
	}
	return Time(i)
}
