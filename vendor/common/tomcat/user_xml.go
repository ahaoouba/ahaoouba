package tomcat

import (
	"encoding/xml"
)

type TomcatUsers struct {
	XMLName xml.Name `xml:"tomcat-users"`
	User    []User   `xml:"user"`
}
type User struct {
	XMLName  xml.Name `xml:"user"`
	Username string   `xml:"username,attr"`
	Password string   `xml:"password,attr"`
	Roles    string   `xml:"roles,attr"`
}

//功能说明:TomcatUsersXML解析
//创建人:孙丽芳
//创建时间:2016-12-16 13:19:29
func AnalysisXML(context []byte) (*TomcatUsers, error) {
	tusers := new(TomcatUsers)
	err := xml.Unmarshal(context, tusers)
	return tusers, err
}
