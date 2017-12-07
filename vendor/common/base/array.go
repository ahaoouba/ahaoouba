package base

import (
	"reflect"
	"strings"
)

func In_Array(a interface{}, b interface{}) bool {
	str1 := reflect.TypeOf(a).Kind().String()

	if str1 != "slice" {
		panic("第一个参数必须是个slice!")
	}
	str3 := strings.Split(reflect.TypeOf(a).String(), "[]")[1]
	str2 := reflect.TypeOf(b).String()

	if str3 != str2 {
		panic("查找的参数与切片内元素类型不匹配!")
	}
	var ok bool
	for i := 0; i < reflect.ValueOf(a).Len(); i++ {

		if reflect.ValueOf(b).Interface() == reflect.ValueOf(a).Index(i).Interface() {
			ok = true
			break
		}
	}
	return ok
}
