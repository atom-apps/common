package tools

import (
	"fmt"
	"time"
)

func HumanDuration(t time.Time) string {
	if t.IsZero() {
		return "-"
	}

	currentTime := time.Now().Unix()

	duration := currentTime - t.Unix()
	seconds := duration % 60
	minutes := (duration / 60) % 60
	hours := (duration / (60 * 60)) % 24
	days := duration / (60 * 60 * 24)

	if days > 0 {
		return fmt.Sprintf("%d天 %d小时 %d分 %d秒", days, hours, minutes, seconds)
	} else if hours > 0 {
		return fmt.Sprintf("%d小时 %d分 %d秒", hours, minutes, seconds)
	} else if minutes > 0 {
		return fmt.Sprintf("%d分 %d秒", minutes, seconds)
	} else {
		return fmt.Sprintf("%d秒", seconds)
	}
}
