package base

import (
	"fmt"
	"testing"
)

func TestFormatSize(t *testing.T) {
	s := " 1024 mB "
	i, err := FormatSize(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
}

func TestFormatTimeToSecond(t *testing.T) {
	s := " 2 hour "
	i, err := FormatTimeToSecond(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
}

func TestHumanSize(t *testing.T) {
	size := float64(92 * 982 * 9182300)
	fmt.Println(HumanSize(size, "B"))
}
