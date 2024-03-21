package intime

type UnixTimestamp interface {
	Unix() int64
}

type MilliTimestamp interface {
	UnixMilli() int64
}

type NanoTimestamp interface {
	UnixNano() int64
}
