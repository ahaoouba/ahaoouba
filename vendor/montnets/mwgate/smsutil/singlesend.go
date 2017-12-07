package smsutil

import "strings"

type SingleSend struct {
	*UserInfo
	*SMData
}

func NewSingleSend(userid string, pwd string, mobile string, content string) *SingleSend {
	content = strings.TrimSpace(content)
	return &SingleSend{NewUserInfo(userid, pwd), NewSMData(mobile, content)}
}
func (s *SingleSend) GetName() string {
	return "single_send"
}
