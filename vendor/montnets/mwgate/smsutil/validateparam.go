package smsutil

import "strings"
import "regexp"
import "strconv"
import "net"

/**
 * @功能概要： 参数合法性验证类
 * @项目名称： SmsDemo
 * @初创作者： tanglili <jack860127@126.com>
 * @公司名称： ShenZhen Montnets Technology CO.,LTD.
 * @创建时间： 2016-7-6 下午12:24:46
 *        <p>
 *        修改记录1：
 *        </p>
 *
 *        <pre>
 *      修改日期：
 *      修改人：
 *      修改内容：
 * </pre>
 */

/**
 * 验证账号是否合法
 *
 * @param userid
 *        账号
 * @return true:合法;false:非法
 */
func ValidateUserId(userid string) bool {
	uid := strings.TrimSpace(userid)
	if "" != uid && len(uid) <= 16 {
		return true
	}
	Logger.Printf("验证账号不合法,错误码: %d, 账号: %s", ERROR_300001, userid)
	return false
}

/**
 * 验证密码是否合法
 *
 * @param pwd
 *        密码
 * @return true:合法;false:非法
 */
func ValidatePwd(pwd string) bool {
	if "" != pwd && len(pwd) <= 32 {
		return true
	}

	if "" == pwd {
		Logger.Printf("验证密码不合法,错误码: %d, 密码为空。", ERROR_300002)
	} else {
		Logger.Printf("验证密码不合法,错误码: %d,  密码长度为%d.", ERROR_300002, len(pwd))
	}
	return false
}

/**
 * 验证流水号是否合法
 *
 * @param custId
 *        流水号
 * @return true:合法;false:非法
 */
func ValidateCustId(custId string) bool {

	// 流水号为空或者流水号大于0，小于等于64位的字符串合法
	if "" == custId {
		return true
	}

	if len(custId) > 64 {
		Logger.Printf("验证流水号不合法,错误码: %d, ,流水号: %s", ERROR_300009, custId)
		return false
	}

	//字母、数字、下划线、减号
	matched, err := regexp.MatchString("^[0-9a-zA-Z-_]+$", custId)
	if nil != err {
		Logger.Println("验证流水号是否合法失败,流水号: %s", custId)
		return false
	}

	if false == matched {
		Logger.Printf("验证流水号不合法,错误码: %d, ,流水号: %s", ERROR_300009, custId)
		return false
	}
	return true

}

/**
 * 验证业务类型是否合法
 *
 * @param svrType
 *        业务类型
 * @return true:合法;false:非法
 */
func ValidateSvrType(svrType string) bool {
	// 业务类型为空
	if "" == svrType {
		return true
	}

	if len(svrType) > 32 {
		Logger.Printf("验证业务类型不合法,错误码: %d, ,流水号: %s", ERROR_300011, svrType)
		return false
	}

	//业务类型必须是字母和数字
	matched, err := regexp.MatchString("^[0-9a-zA-Z-]+$", svrType)
	if nil != err {
		Logger.Println("验证业务类型不合法,业务类型: ", svrType)
		return false
	}

	if false == matched {
		Logger.Printf("验证业务类型不合法,错误码: %d, ,业务类型: %s", ERROR_300011, svrType)
		return false
	}
	return true
}

/**
 * 验证exdata是否合法
 *
 * @param exData
 *        自定义扩展参数
 * @return true:合法;false:非法
 */
func ValidateExData(exData string) bool {

	// exdata为空
	if "" == exData {
		return true
	}

	if len(exData) > 64 {
		Logger.Printf("验证自定义扩展数据不合法,错误码: %d, ,自定义扩展数据: %s", ERROR_300012, exData)
		return false
	}

	//字母、数字、下划线、减号
	matched, err := regexp.MatchString("^[0-9a-zA-Z-_]+$", exData)
	if nil != err {
		Logger.Println("验证自定义扩展数据不合法,自定义扩展数据: %s", exData)
		return false
	}

	if false == matched {
		Logger.Printf("验证自定义扩展数据不合法,错误码: %d, ,自定义扩展数据: %s", ERROR_300012, exData)
		return false
	}
	return true
}

/**
 * 验证信息内容是否合法
 *
 * @param content
 *        短信内容
 * @return true:合法;false:非法
 */
func ValidateMessage(content string) bool {
	// 短信内容不为空
	content = strings.TrimSpace(content)
	if "" == content {
		Logger.Printf("验证信息内容不合法 ,错误码: %d, 短信内容为空。", ERROR_300007)
		//短信内容为空
		return false
	}

	isChinese := false

	for _, v := range content {
		if v > 127 {
			isChinese = true
			break
		}
	}
	// 中文短信长度必须小于990, 英文短信长度必须小于1980
	if (isChinese && len(content) <= 990) || (!isChinese && len(content) <= 1980) {
		return true
	}
	Logger.Printf("验证信息内容不合法 ,错误码: %d, ,短信内容过长,短信内容: %s", ERROR_300007, content)

	return false
}

/**
 * 验证扩展子号是否合法
 *
 * @param strSubPort
 *        扩展子号
 * @return true:合法;false:非法
 */
func ValidateSubPort(strSubPort string) bool {
	// 扩展子号为空，合法
	if "" == strSubPort {
		return true
	}

	if len(strSubPort) > 6 {
		Logger.Printf("验证扩展子号是否合法失败,错误码: %d, ,扩展子号: %s", ERROR_300008, strSubPort)
		return false
	}

	//字母、数字、下划线、减号
	matched, err := regexp.MatchString("^[0-9]+$", strSubPort)
	if nil != err {
		Logger.Println("验证扩展子号不合法,扩展子号: %s", strSubPort)
		return false
	}

	if false == matched {
		Logger.Printf("验证扩展子号不合法,错误码: %d, ,扩展子号: %s", ERROR_300008, strSubPort)
		return false
	}
	return true
}

/**
 * 验证手机单个号码是否合法
 * 手机号码可能是国际号码，暂时不验证手机号码长度
 *
 * @param mobile
 *        手机号码
 * @return true:合法;false:非法
 */
func ValidateMobile(mobile string) bool {
	// 手机号码不能为空
	if "" == mobile {
		return true
	}

	// 长度小于等于21位
	if len(mobile) > 21 {
		Logger.Printf("验证手机单个号码不合法,错误码: %d, ,手机号码: %s", ERROR_300006, mobile)
		return false
	}

	// 必须是数字
	matched, err := regexp.MatchString("^[0-9]+$", mobile)
	if nil != err {
		Logger.Println("验证手机单个号码不合法,手机号码: %s", mobile)
		return false
	}

	if false == matched {
		Logger.Printf("验证手机单个号码不合法,错误码: %d, ,手机号码: %s", ERROR_300006, mobile)
		return false
	}
	return true
}

/**
 * 验证是否是以英文逗号隔开的1000个手机号码
 * 手机号码可能是国际号码，暂时不验证手机号码长度
 **
 * @param strMobiles
 *        手机号码
 * @return success:合法;fail:检测失败;illegalFormat:格式非法;overNum:超过最大支持号码1000个
 */
func ValidateMobiles(mobiles string) string {
	// 手机号码字符串不能为空
	if "" == mobiles {
		Logger.Println("验证手机号码不合法,错误码:%d, 手机号码 为空。", ERROR_300006)
		return "illegalFormat"
	}

	arrMoblie := strings.Split(mobiles, ",")

	// 手机号码个数必须大于0并且小于等于1000
	if len(arrMoblie) > 0 && len(arrMoblie) <= 1000 {
		for _, m := range arrMoblie {
			matched, err := regexp.MatchString("^[0-9]+$", m)
			if nil != err {
				Logger.Println("验证手机号码不合法,错误码:%d, 错误手机号码: %s, 所有手机号码: %s", ERROR_300006, m, mobiles)
				return "illegalFormat"
			}
			if false == matched {
				Logger.Printf("验证手机号码不合法,错误码: %d, ,错误手机号码: %s,所有手机号码: %s", ERROR_300006, m, mobiles)
				return "illegalFormat"
			}
		}
	} else if len(arrMoblie) > 1000 {
		Logger.Println("验证手机号码不合法,错误码:%d, 手机号码者超过1000, 所有手机号码%s", ERROR_300006, mobiles)
		return "overNum"
	} else {
		Logger.Println("验证手机号码不合法,错误码:%d, 手机号码 为空。", ERROR_300006)
		return "illegalFormat"
	}
	return "success"
}

/**
 * @description 验证IP和端口信息是否合法
 * @param ipAddress
 *        IP和端口信息 IP和端口号以:号连接
 * @return 0：合法;非0:不合法
 * @author tanglili <jack860127@126.com>
 * @datetime 2016-9-22 下午05:02:33
 */
func ValidateIpAndPort(ipAddress string) int {
	// 验证IP和端口 IP和端口不能为空并且必须包含冒号:
	ipAddress = strings.TrimSpace(ipAddress)
	if "" == ipAddress || -1 == strings.LastIndex(ipAddress, ":") {
		Logger.Println("验证IP和端口信息不合法,错误码:%d, ,IP和端口信息: %s", ERROR_300003, ipAddress)
		return ERROR_300003
	}

	// 验证IP
	if nil == net.ParseIP(strings.Split(ipAddress, ":")[0]) {
		Logger.Println("验证IP地址不合法,错误码:%d, ,IP地址: %s", ERROR_300004, strings.Split(ipAddress, ":")[0])
	}
	// 验证端口
	nport, err := strconv.Atoi(strings.Split(ipAddress, ":")[1])
	if nil != err || nport > 65536 {
		Logger.Println("验证端口不合法,错误码:%d, ,端口: %s", ERROR_300005, strings.Split(ipAddress, ":")[1])
		return ERROR_300005
	}
	return 0
}
