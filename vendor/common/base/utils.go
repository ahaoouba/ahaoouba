package base

import (
	"fmt"
	"reflect"
)

func IsZero(i interface{}) bool {
	if i == nil {
		return true
	}
	fmt.Println(reflect.TypeOf(i).Kind())
	v := reflect.ValueOf(i)
	switch reflect.TypeOf(i).Name() {
	case "string":
		return v.String() == ""
	case "int":
		return v.Int() == 0
	case "float":
		return v.Float() == 0

	}
	return false
}
