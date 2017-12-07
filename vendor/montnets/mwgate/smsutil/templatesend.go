package smsutil

type VoiceData struct {

	//短信接收的手机号，用英文逗号(,)分隔，最大1000个号码。一次提交的号码类型不受限制，但手机会做验证，若有不合法的手机号将会被退回。号码段类型分为：移动、联通、电信手机
	// 注意：请不要使用中文的逗号
	Mobile string `json:"mobile"`

	//最大支持350个字，一个字母或一个汉字都视为一个字
	Content string `json:"content"`

	// 消息类型：
	// 1：语音验证码
	// 3：语音通知：只有当显号为12590时，实际发出的消息类型仍为语音验证码，并且使用梦网自带的语音模板发送语音验证码，其他显号下仍然使用语音模板编号对应的模板发送语音通知。
	Msgtype string `json:"msgtype"`

	// 语音模版编号：
	// 当msgtype为1时，语音模板编号为非必须项，如提交此字段则使用与提交模板编号对应的模板发送语音验证码，如不提交此字段或填空则使用梦网自带的语音模板发送语音验证码
	// 当msgtype为3时，语音模板编号为必填项
	Tmplid string `json:"tmplid"`

	// 扩展号
	// 长度由账号类型定4-6位，通道号总长度不能超过20位。如：10657****主通道号，3321绑定的扩展端口，主+扩展+子端口总长度不能超过20位。
	Exno string `json:"exno"`

	// 该条短信在您业务系统内的用户自定义流水编号，比如订单号或者短信发送记录的流水号。填写后发送状态返回值内将包含这个ID.最大可支持64位的字符串
	Custid string `json:"custid"`
}

func NewVoiceData(mobile string, content string, msgtype string, tmplid string) (voidata *VoiceData) {
	voidata = &VoiceData{
		Mobile:  mobile,
		Content: FormatContent(content),
		Msgtype: msgtype,
		Tmplid:  tmplid,
		Custid:  GetCustid(),
		Exno:    ""}
	return voidata
}

type TemplateSend struct {
	*UserInfo
	*VoiceData
}

func NewTemplateSend(userid string, pwd string, mobile string, content string, msgtype string, tmplid string) *TemplateSend {
	return &TemplateSend{NewUserInfo(userid, pwd), NewVoiceData(mobile, content, msgtype, tmplid)}
}

func (s *TemplateSend) GetName() string {
	return "template_send"
}
