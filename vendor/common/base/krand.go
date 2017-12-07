package base

import (
	mrand "math/rand"
	"time"
)

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

func KrandTypeNum(length int) []byte {
	return krand(length, KC_RAND_KIND_ALL)
}

func KrandTypeLower(length int) []byte {
	return krand(length, KC_RAND_KIND_LOWER)
}

func KrandTypeUpper(length int) []byte {
	return krand(length, KC_RAND_KIND_UPPER)
}

func KrandTypeAll(length int) []byte {
	return krand(length, KC_RAND_KIND_ALL)
}

// 随机字符串
// size 是生成的字符串长度
func krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	mrand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = mrand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + mrand.Intn(scope))
	}
	return result
}
