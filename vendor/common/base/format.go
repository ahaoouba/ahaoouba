package base

import (
	"fmt"
	"strconv"
	"strings"
)

// 格式化百分比字符串
func FormatPercent(s string) (float64, error) {

	return strconv.ParseFloat(strings.TrimSpace(strings.TrimSuffix(s, "%")), 10)

}

// 格式化容量大小字符串,结尾必须是KB,MB,GB,TB
// 返回结果为已KB为单位的int64数值
// 1024 KB  -->  1024
// 1024 MB  -->  1048576
func FormatSize(s string) (int64, error) {

	s = strings.ToUpper(strings.TrimSpace(s))
	size := "0"
	unit := int64(1)
	if strings.Contains(s, "KB") {
		size = strings.TrimSpace(strings.TrimSuffix(s, "KB"))
	} else if strings.Contains(s, "MB") {
		size = strings.TrimSpace(strings.TrimSuffix(s, "MB"))
		unit = 1024
	} else if strings.Contains(s, "GB") {
		size = strings.TrimSpace(strings.TrimSuffix(s, "GB"))
		unit = 1024 * 1024
	} else if strings.Contains(s, "TB") {
		size = strings.TrimSpace(strings.TrimSuffix(s, "TB"))
		unit = 1024 * 1024 * 1024
	} else if strings.Contains(s, "PB") {
		size = strings.TrimSpace(strings.TrimSuffix(s, "PB"))
		unit = 1024 * 1024 * 1024 * 1024
	} else {
		return 0, fmt.Errorf("不支持的数据存储单位，[ %s ]", s)
	}

	i, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		return 0, err
	}

	return i * unit, nil
}

// 将字符串格式化成秒
//443010 sec --> 443010
//2 min --> 120
//1 hour --> 3600
func FormatTimeToSecond(s string) (int64, error) {
	s = strings.ToLower(strings.TrimSpace(s))
	size := "0"
	unit := int64(1)
	if strings.Contains(s, "sec") {
		size = strings.TrimSpace(strings.TrimSuffix(s, "sec"))
	} else if strings.Contains(s, "min") {
		size = strings.TrimSpace(strings.TrimSuffix(s, "min"))
		unit = 60
	} else if strings.Contains(s, "hour") {
		size = strings.TrimSpace(strings.TrimSuffix(s, "hour"))
		unit = 60 * 60
	}
	i, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		return 0, err
	}
	return i * unit, nil
}

// 存储单位转换
//unit 表示size的单位
func HumanSize(size float64, unit string) string {
	u := strings.ToUpper(unit)
	m := size / 1024
	if m >= 1 {
		switch u {
		case "B":
			return HumanSize(m, "KB")
		case "KB":
			return HumanSize(m, "MB")
		case "MB":
			return HumanSize(m, "GB")
		case "GB":
			return HumanSize(m, "TB")
		case "TB":
			return HumanSize(m, "PB")
		case "PB":
			return HumanSize(m, "EB")
		case "EB":
			return HumanSize(m, "ZB")
		case "ZB":
			return HumanSize(m, "YB")
		case "YB":
			return HumanSize(m, "BB")
		}
	}

	return fmt.Sprintf("%.2f %s", size, unit)
}
