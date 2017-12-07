package base

import (
	"os"
)

// 建立多个目录
func Mkdirs(dirs []string) error {
	for _, d := range dirs {
		err := os.MkdirAll(d, 0775)
		if err != nil {
			return err
		}
	}
	return nil
}

// 检测文件是否存在
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
