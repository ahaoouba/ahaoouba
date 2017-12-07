package tomcat

import (
	"encoding/xml"
)

type server struct {
	XMLName               xml.Name                `xml:"Server"`
	Shutdown              string                  `xml:"shutdown,attr"`
	Port                  int                     `xml:"port,attr"`
	Listener              []listener              `xml:"Listener"`
	GlobalNamingResources []globalNamingResources `xml:"GlobalNamingResources"`
	Service               []service               `xml:"Service"`
}
type listener struct {
	XMLName   xml.Name `xml:"Listener"`
	SSLEngine string   `xml:"SSlEngine,attr"`
	ClassName string   `xml:"className,attr"`
}

type globalNamingResources struct {
	XMLName  xml.Name `xml:"GlobalNamingResources"`
	Resource resource `xml:"Resource"`
}
type resource struct {
	XMLName     xml.Name `xml:"Resource"`
	Auth        string   `xml:"auth,attr"`
	Description string   `xml:"description,attr"`
	Factory     string   `xml:"factory,attr"`
	Nmae        string   `xml:"name,attr"`
	Pathname    string   `xml:"pathname,attr"`
	Type        string   `xml:"type,attr"`
}

type service struct {
	XMLName   xml.Name    `xml:"Service"`
	Name      string      `xml:"name,attr"`
	Connector []connector `xml:"Connector"`
	Engine    engine      `xml:"Engine"`
}

type connector struct {
	XMLName           xml.Name `xml:"Connector"`
	Port              string   `xml:"port,attr"`
	RedirectPort      string   `xml:"redirectPort,attr"`
	Protocol          string   `xml:"protocol,attr"`
	ConnectionTimeout string   `xml:"connectionTimeout,attr"`
	URIEncoding       string   `xml:"URIEncoding,attr"`
}
type engine struct {
	XMLName     xml.Name `xml:"Engine"`
	Name        string   `xml:"name,attr"`
	DefaultHost string   `xml:"defaultHost,attr"`
	Realm       realm    `xml:"Realm"`
	Host        host     `xml:"Host"`
}
type realm struct {
	XMLName      xml.Name `xml:"Realm"`
	ClassName    string   `xml:"className,attr"`
	ResourceName string   `xml:"resourceName,attr"`
	Realm        realms   `xml:"Realm"`
}
type realms struct {
	XMLName      xml.Name `xml:"Realm"`
	ClassName    string   `xml:"className,attr"`
	ResourceName string   `xml:"resourceName,attr"`
}
type host struct {
	XMLName           xml.Name `xml:"Host"`
	Name              string   `xml:"name,attr"`
	AutoDeploy        string   `xml:"autoDeploy,attr"`
	AppBase           string   `xml:"appBase,attr"`
	UnpackWARs        string   `xml:"unpackWARs,attr"`
	XmlValidation     string   `xml:"xmlValidation,attr"`
	XmlNamespaceAware string   `xml:"xmlNamespaceAware,attr"`
	Valve             valve    `xml:"Valve"`
	Context           context  `xml:"Context"`
}
type valve struct {
	XMLName   xml.Name `xml:"Valve"`
	ClassName string   `xml:"classname,attr"`
	Suffix    string   `xml:"suffix,attr"`
	Prefix    string   `xml:"prefix,attr"`
	Pattern   string   `xml:"pattern,attr"`
	Directory string   `xml:"directory,attr"`
}
type context struct {
	XMLName    xml.Name `xml:"Context"`
	Source     string   `xml:"source,attr"`
	Reloadable string   `xml:"reloadable,attr"`
	Path       string   `xml:"path,attr"`
	BocBase    string   `xml:"docBase,attr"`
}

//功能说明:TomcatServerXML解析
//创建人:夏龙飞
//创建时间:2016-12-16 13:19:29
func ServerAnaluze(content []byte) (*server, error) {
	result := new(server)
	err := xml.Unmarshal(content, result)

	return result, err
}
