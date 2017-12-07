package smsutil

type MultiSend struct {
	*UserInfo
	Multimt []SMData `json:"multimt"`
}

func NewMultiSend(userid string, pwd string, multmt []SMData) *MultiSend {
	return &MultiSend{NewUserInfo(userid, pwd), multmt}
}
func (m *MultiSend) GetName() string {
	return "multi_send"
}
