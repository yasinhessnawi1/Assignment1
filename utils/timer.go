package utils

import (
	"time"
)

var uptimeStartTime time.Time

/*
StartUptime starts recording uptime
*/
func StartUptime() {
	uptimeStartTime = time.Now()
}

/*
GetUptime returns the current uptime duration
*/
func GetUptime() time.Duration {
	return time.Since(uptimeStartTime)
}
