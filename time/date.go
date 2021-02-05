package time

import (
	"time"
	"github.com/jinzhu/now"
)

// 字符串转时间，支持大部分格式
func Parse(t string) (time.Time, error) {
	return now.Parse(t)
}