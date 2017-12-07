package smsutil

import "encoding/json"
import "time"
import "strings"

/**
 *
 * @功能概要：检测IP信息
 *
 */

type CheckAddress struct {
	Cm *ConfigManager
}

/**
 * @description 检查地址信息是否可用
 * @param userId
 *        账号
 * @param password
 *        密码
 * @param timestamp
 *        时间戳
 * @param ipAddress
 *        地址信息
 * @return 0:代表可用; 非0:代表不可用
 */
func (c *CheckAddress) checkAddressAvailable(userId string, password string, timestamp string, ipAddress string) int {
	result := ERROR_310099

	// 用查询余额接口get_balance接口去检测地址是否有效
	// 查询余额接口的URLEncode协议格式为：userid=[用户名]&pwd=[密码]&timestamp=[时间戳]
	post_body := []byte("userid=" + strings.ToUpper(userId) + "&pwd=" + password + "&timestamp" + timestamp)

	// 请求地址: http://[域名]/[路径]/get_blance
	balanceHost := "http://" + ipAddress + c.Cm.RequestPath + "get_blance"
	res_code, response, err := HttpPostOnce(balanceHost, post_body, "x-www-form-urlencoded")
	if nil != err {
		return result
	}

	if 200 == res_code {
		//获取响应的实体
		var r interface{}
		err := json.Unmarshal(response, &r)

		if nil != err {
			return ERROR_310099
		}
		// 查询余额成功
		return 0
	}
	return ERROR_310099
}

/**
 * 检查主IP
 * @description
 * @return
 */
func (c *CheckAddress) checkMasterAddress(userid string, password string, timestamp string) string {
	if nil == c.Cm {
		c.Cm = GetConfigManager()
	}

	c.Cm.LastCheckTime = time.Now().Unix()
	result := c.checkAddressAvailable(userid, password, timestamp, c.Cm.MasterIpAndPort)

	if 0 == result {
		c.Cm.AvailableIpAndPort = c.Cm.MasterIpAndPort
		//将主IP设置为可用
		c.Cm.MasterIPState = 0
		// todo 写日志
		// System.out.println("主ipAddress["+ConfigManager.masterIpAndPort+"]恢复正常。");
		return c.Cm.AvailableIpAndPort
	} else {
		if "" != c.Cm.MasterDomainAndPort {
			//新获取的IP和端口
			newIpAndPort, err := c.Cm.AddressToIPAndPort(c.Cm.MasterDomainAndPort)
			if true != err {
				return ""
			}
			//如果失败，通过新获取的IP检查余额
			result = c.checkAddressAvailable(userid, password, timestamp, newIpAndPort)
			// result为0，代表成功 result为非0，则代表失败
			if 0 == result {
				if c.Cm.MasterIpAndPort == newIpAndPort {
					// todo 写日志
					// 	System.out.println("主ipAddress["+ConfigManager.masterIpAndPort+"]恢复正常。");
				} else {
					// todo 写日志
					// 	System.out.println("主ipAddress由["+ConfigManager.masterIpAndPort+"]切换为["+newIpAndPort+"]。");
				}
				c.Cm.AvailableIpAndPort = newIpAndPort
				c.Cm.MasterIPState = 0
				c.Cm.MasterIpAndPort = newIpAndPort
				// todo 写日志
				//	System.out.println("通过域名获取的主IP正常,主IP和端口："+newIpAndPort);
				return c.Cm.AvailableIpAndPort
			}
		}
	}
	return ""
}

func (c *CheckAddress) GetAddressByUserID(userid string, password string, timestamp string) string {
	if nil == c.Cm {
		c.Cm = GetConfigManager()
	}

	// 判断主IP是否正常
	if 0 == c.Cm.MasterIPState {
		// 主IP与可用IP不相等，则将可用IP设置为主IP
		if c.Cm.AvailableIpAndPort != c.Cm.MasterIpAndPort {
			// 将可用IP设置为主IP
			c.Cm.AvailableIpAndPort = c.Cm.MasterIpAndPort
		}
		// 正常
		return c.Cm.AvailableIpAndPort
	} else {
		//主IP异常，如果主IP异常时间超过5分钟，则检测异常主IP
		if time.Now().Unix()-c.Cm.LastCheckTime > c.Cm.CheckTimeInterval {
			address := c.checkMasterAddress(userid, password, timestamp)
			if "" != address {
				return address
			}
		}

		// 循环备用IP和可用IP比较，如果相等，则说明该备用IP可用
		if len(c.Cm.IpAndPortBak) > 0 {
			for _, v := range c.Cm.IpAndPortBak {
				if c.Cm.AvailableIpAndPort == v {
					return v
				}
			}
		}

		// 可用的IP地址不存在，则循环检测备用IP是否可用
		availableAddress := c.checkAddress(userid, password, timestamp)
		if "" != availableAddress {
			c.Cm.AvailableIpAndPort = availableAddress
			return availableAddress
		}
	}

	return ""
}

/**
 * 检测出可用的IP地址
 *
 * @description
 * @param userid
 *        账号
 * @param password
 *        密码
 * @param timestamp
 *        时间戳
 * @return null：无可用的IP地址;否则有可用的IP地址
 */
func (c *CheckAddress) checkAddress(userid string, password string, timestamp string) string {
	address := ""

	result := ERROR_310099
	ipAndPortBakList := c.Cm.IpAndPortBak
	isDomainGet := false
	ipAddressBak := ""
	for _, v := range ipAndPortBakList {
		ipAddressBak = v
		// 调用查询余额的方法检测连接是否可用
		result = c.checkAddressAvailable(userid, password, timestamp, ipAddressBak)
		// 查询余额成功
		if result == 0 {
			address = ipAddressBak
			break
		} else {
			//如果用IP和端口查询失败，则检查IP和端口是否有对应的域名
			domainBak, err := c.Cm.IpAndDomainBakMap[ipAddressBak]
			if false == err {
				//新获取的IP
				newIpAndPort, err := c.Cm.AddressToIPAndPort(c.Cm.MasterDomainAndPort)
				if false == err {
					return ""
				}
				// 调用查询余额的方法检测该备用IP对应的域名是否正常
				result = c.checkAddressAvailable(userid, password, timestamp, newIpAndPort)
				if 0 == result {
					isDomainGet = true
					//移除旧的备用IP和域名的对应关系
					delete(c.Cm.IpAndDomainBakMap, ipAddressBak)
					//添加新的IP和域名的对应关系
					c.Cm.IpAndDomainBakMap[newIpAndPort] = domainBak
					//将新的IP和端口赋值给address
					address = newIpAndPort
					break
				}
			}
		}
	}
	//IP地址不为空并且通过域名获取
	if "" != address && isDomainGet {
		//移除备用IP
		RemoveAddressElement(c.Cm.IpAndPortBak, ipAddressBak)
		//新增备用IP
		c.Cm.IpAndPortBak = append(c.Cm.IpAndPortBak, address)
	}
	return address
}
