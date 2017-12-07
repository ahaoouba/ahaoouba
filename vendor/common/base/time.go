package base

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 获取当前时间
func GetCurrentData() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//返回当前时间的时间戳格式
func GetCurrentDataUnix() int64 {
	return time.Now().Unix()
}

func ParsePeriod(p string) (time.Duration, error) {
	// 5s   5d
	p = strings.TrimSpace(p)
	reg, err := regexp.Compile("(\\d+)\\s{0,}([s|m|h|d])")
	if err != nil {
		return 0, err
	}
	s := reg.FindStringSubmatch(p)
	if len(s) == 0 {
		return 0, errors.New(fmt.Sprintf("解析period参数失败，period为：[ %s ]", p))
	}
	tm, err := strconv.ParseInt(s[1], 10, 64)
	unit := s[2]
	switch unit {
	case "s":
		return time.Duration(tm) * time.Second, nil
	case "m":
		return time.Duration(tm) * time.Minute, nil
	case "h":
		return time.Duration(tm) * time.Hour, nil
	case "d":
		return time.Duration(tm*24) * time.Hour, nil
	}
	return 0, nil

}

// 将unix时间戳转换成time.Time
func ConvertUnixToTime(tm string) (time.Time, error) {
	itm, err := strconv.ParseInt(tm, 10, 64)
	if err != nil {
		return time.Now(), err
	}
	return time.Unix(itm, 0), nil
}
