package v1

import (
	"fmt"
)

type Service struct {
	Name    string `json:"name"`
	Mode    string `json:"mode"`
	Address string `json:"address"`
	TTL     int    `json:"ttl"`
}

// 验证添加服务信息参数
func (c Service) ValidRegist() error {

	if c.Name == "" {
		return fmt.Errorf("name 不能为空")
	}
	if c.Mode == "" {
		return fmt.Errorf("mode 不能为空")
	}
	if c.Address == "" {
		return fmt.Errorf("address 不能为空")
	}
	return nil
}
