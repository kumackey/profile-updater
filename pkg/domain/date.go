package domain

import "time"

const (
	JSTOffset = 9 * 60 * 60
	asiaTokyo = "Asia/Tokyo"
)

var locationJST = time.FixedZone(asiaTokyo, JSTOffset)

func convertTimeToString(time time.Time) string {
	return time.In(locationJST).Format("Jan 2")
}
