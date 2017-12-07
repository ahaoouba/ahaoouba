package smsutil

type SMData struct {

	//短信接收的手机号，用英文逗号(,)分隔，最大1000个号码。一次提交的号码类型不受限制，但手机会做验证，若有不合法的手机号将会被退回。号码段类型分为：移动、联通、电信手机
	// 注意：请不要使用中文的逗号
	Mobile string `json:"mobile"`

	//最大支持350个字，一个字母或一个汉字都视为一个字
	Content string `json:"content"`

	// 扩展号
	// 长度由账号类型定4-6位，通道号总长度不能超过20位。如：10657****主通道号，3321绑定的扩展端口，主+扩展+子端口总长度不能超过20位。
	Exno string `json:"exno"`

	// 该条短信在您业务系统内的用户自定义流水编号，比如订单号或者短信发送记录的流水号。填写后发送状态返回值内将包含这个ID.最大可支持64位的字符串
	Custid string `json:"custid"`

	// 额外提供的最大64个长度的自定义扩展数据.填写后发送状态返回值内将会包含这部分数据
	Exdata string `json:"exdata"`

	//业务类型
	SvrType string `json:"svrtype"`
}

func NewSMData(mobile string, content string) (smdata *SMData) {
	smdata = &SMData{
		Mobile:  mobile,
		Content: FormatContent(content),
		Custid:  GetCustid(),
		Exdata:  GetExdata(),
		Exno:    GetExno(),
		SvrType: GetSvrtype()}
	return smdata
}
