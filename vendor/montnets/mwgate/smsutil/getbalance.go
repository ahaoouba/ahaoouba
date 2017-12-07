package smsutil

type GetBalance struct {
	*UserInfo
}

func NewGetBalance(userid string, pwd string) *GetBalance {
	return &GetBalance{NewUserInfo(userid, pwd)}
}

func (s *GetBalance) GetName() string {
	return "get_balance"
}
