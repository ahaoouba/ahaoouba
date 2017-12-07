package smsutil

import "strings"
import "net"
import "regexp"

/**
 * 配置管理类
 * @author Administrator
 *
 */
type ConfigManager struct {
	// 请求路径
	RequestPath string

	//主IP和端口
	MasterIpAndPort string

	//主域名和端口    主域名和端口可能为空
	MasterDomainAndPort string

	// 备用IP端口信息 IP和端口号以:号连接
	IpAndPortBak []string

	//备IP和备域名的对应关系。key:备用IP和端口   value:域名和端口
	IpAndDomainBakMap map[string]string

	// 主IP状态 0正常 1异常
	MasterIPState int

	//可用IP
	AvailableIpAndPort string

	//主IP最近检测时间 1970年距今的毫秒数
	LastCheckTime int64

	//主IP异常检测时间间隔 5分钟
	CheckTimeInterval int64

	//IP是否设置
	IpIsSet bool

	//密码是否加密 默认加密
	IsEncryptPwd bool
}

var singleConfigManager = &ConfigManager{
	RequestPath:         "/sms/v2/std/",
	MasterIpAndPort:     "",
	MasterDomainAndPort: "",
	IpAndPortBak:        []string(nil),
	IpAndDomainBakMap:   make(map[string]string, 3),
	MasterIPState:       0,
	AvailableIpAndPort:  "",
	LastCheckTime:       0,
	CheckTimeInterval:   5 * 60,
	IpIsSet:             false,
	IsEncryptPwd:        true}

/**
 *  取配置对象
 *  @return 返回配置对象
 */
func GetConfigManager() *ConfigManager {
	return singleConfigManager
}

/**
 *  将域名替换成IP
 *	@param doMain 域名
 *  @return ip, err 如果成功返回首IP,不成功返回错误信息
 */
func (cm *ConfigManager) AddressToIPAndPort(addr string) (string, bool) {

	// 地址形如 www.domain.com:7901或者123.123.13.11:7901
	ipandport := strings.Split(addr, ":")
	//  先假设地址为域名，解析域名，获取IP列表
	iplist, err := net.LookupIP(ipandport[0])
	if nil != err {
		// 获取IP列表失败，再判断是否是合法IP,不是合法IP,返回失败。
		if nil == net.ParseIP(ipandport[0]) {
			return "", false
		}
		return addr, true
	}

	// 判断是否有端口信息, 如果有就用第一个IP替换域名，如果没有直接返回IP
	if len(ipandport) > 1 {
		addr = iplist[0].String() + ":" + ipandport[1]
	} else {
		addr = iplist[0].String()
	}

	return addr, true
}

/**
 * 取配置管理信息
 * @return 配置管理信息
 */

/**
 * 设置短信平台地址的方法
 * @param masterAddress 主地址，可以是域名和端口或者IP和端口
 * @param bakAddress1 备用，可以是域名或者IP
 * @param bakAddress2 备用，可以是域名或者IP
 *  * @param bakAddress3 备用，可以是域名或者IP
 * @return 0:成功 ;-1:失败
 */

func (cm *ConfigManager) SetIpInfo(masterAddress string, bakAddress1 string,
	bakAddress2 string, bakAddress3 string) (ret int) {

	// 主IP是""直接返回失败
	if "" == masterAddress {
		return -1

	}

	// 将域名替换成IP
	ipAndPort, ok := cm.AddressToIPAndPort(masterAddress)
	if false == ok {
		return -1
	}
	//设置主域名
	cm.MasterDomainAndPort = masterAddress
	// 设置主IP
	cm.MasterIpAndPort = ipAndPort

	// 将域名替换成IP
	ipAndPort, ok = cm.AddressToIPAndPort(bakAddress1)
	if ok {
		cm.IpAndDomainBakMap[ipAndPort] = bakAddress1
		cm.IpAndPortBak = append(cm.IpAndPortBak, ipAndPort)
	}

	// 将域名替换成IP
	ipAndPort, ok = cm.AddressToIPAndPort(bakAddress2)
	if ok {
		cm.IpAndDomainBakMap[ipAndPort] = bakAddress2
		cm.IpAndPortBak = append(cm.IpAndPortBak, ipAndPort)
	}

	// 将域名替换成IP
	ipAndPort, ok = cm.AddressToIPAndPort(bakAddress3)
	if ok {
		cm.IpAndDomainBakMap[ipAndPort] = bakAddress3
		cm.IpAndPortBak = append(cm.IpAndPortBak, ipAndPort)
	}
	return 0
}

/**
 * 判断是否是域名
 * @param address IP信息
 * @return
 */
func (cm *ConfigManager) isDomain(addr string) bool {
	matched, err := regexp.MatchString("[a-zA-Z0-9][-a-zA-Z0-9]{0,62}((\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+)", addr)
	if nil != err {
		return false
	}

	return matched
}

/**
 * 清除所有设置的IP
 * @description
 * @return
 */
func (cm *ConfigManager) RemoveAllIpInfo() {
	cm.MasterDomainAndPort = ""
	cm.MasterIpAndPort = ""
	cm.IpAndPortBak = []string(nil)
	cm.IpAndDomainBakMap = make(map[string]string, 3)
}
