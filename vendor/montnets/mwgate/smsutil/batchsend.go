package smsutil

import "strings"

type BatchSend struct {
	*UserInfo
	*SMData
}

func NewBatchSend(userid string, pwd string, mobile string, content string) *BatchSend {
	content = strings.TrimSpace(content)
	return &BatchSend{NewUserInfo(userid, pwd), NewSMData(mobile, content)}
}

func (s *BatchSend) GetName() string {
	return "batch_send"
}
