package quark

import (
	"github.com/spf13/cast"
	"regexp"
	"time"
)

var pattern = regexp.MustCompile(`https://pan\.quark\.cn/s/([a-f0-9]{12})`)

func parseShareCode(url string) string {
	var matched []string
	if matched = pattern.FindStringSubmatch(url); len(matched) > 1 {
		return matched[1]
	}
	return ""
}
func GetTimestamp() string {
	// 获取当前时间的纳秒级时间戳
	nanoseconds := time.Now().UnixNano()

	// 将纳秒级时间戳转换为毫秒级时间戳
	milliseconds := nanoseconds / int64(time.Millisecond)
	return cast.ToString(milliseconds)

}
