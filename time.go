package player

import "time"

func ParseTime(num int) time.Duration {
	floatNum := float64(num) * float64(5668.9342403628117913832199546485260770975056689342403628117)
	return time.Duration(floatNum)
}
