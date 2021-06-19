package datetime

import "time"

const (
	MINUTE    = 60
	MINUTE_5  = 300
	MINUTE_15 = 900
	MINUTE_30 = 1800
	HOUR      = 3600
	HOUR_2    = 7200
	HOUR_6    = 21600
	HOUR_12   = 43200
	DAY       = 86400
	WEEK      = 604800
	MONTH     = 2592000
	YEAR      = 31104000
)

func Now() float32 {
	return float32(0.000001 * float64(time.Now().UnixNano()))
}

func Current() int64 {
	return time.Now().Unix()
}
