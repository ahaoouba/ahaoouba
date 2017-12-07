package smsutil

import "time"
import "net/url"
import "strings"
import "crypto/md5"
import "encoding/hex"
import "encoding/json"
import "net/http"
import "fmt"
import "io/ioutil"
import "log"
import "github.com/axgle/mahonia"
import "net"
import "regexp"
import "github.com/astaxie/beego"

// 日志
var Logger *log.Logger

// 流水号，内部使用
var serial_number uint = 1

/**
 *  格式化输出赶时间，如时间是：2017-5-11 14:22:30 输出为：0511142230
 *  @return 返回时间串
 */
func FormatCurrentTime() string {
	now := time.Now()
	return now.Format("060102150405")[2:]
}

/**
 *  将UTF8串转为GBK并进行URLENCODE
 *  @return 返回URLENCODE的字符串
 */
func FormatContent(content string) string {
	// 去掉两端的空格
	content = strings.TrimSpace(content)
	// 转为GBK
	gbk := mahonia.NewEncoder("gbk").ConvertString(content)
	v := url.Values{}
	v.Set("aa", gbk)
	str := v.Encode()
	arr := strings.Split(str, "=")
	return arr[1]
}

/**
 *  取扩展号：长度不能超过6位，注意通道号+扩展号的总长度不能超过20
 *  位，若超出则exno无效，如不需要扩展号则不用提交此字段或填空, 用户
 *  可以改写里面的内容。
 *  @return 返回扩展号
 */
func GetExno() string {
	return "1234"
}

/**
 *  取用户自定义流水号：该条短信在您业务系统内的ID，比如订单号或者
 *  短信发送记录的流水号。填写后发送状态返回值内将包含用户自定义流
 *  水号。最大可支持64位的ASCII字符串：字母、数字、下划线、减号，如
 *  不需要则不用提交此字段或填空, 用户可以改写里面的内容。
 *  @return 返回用户自定义流水号
 */
func GetCustid() (custid string) {
	now := time.Now()
	serialno := fmt.Sprintf("%06d", serial_number)
	serial_number += 1
	num := now.Format("060102150405")
	custid = "20" + num + serialno
	return custid
}

/**
 *  自定义扩展数据：额外提供的最大64个长度的ASCII字符串：字母、数字、
 *  下划线、减号，作为自定义扩展数据，填写后，状态报告返回时将会包含
 *   这部分数据,如不需要则不用提交此字段或填空, 用户可以改写里面的内容。
 *  @return 返回自定义扩展数据
 */
func GetExdata() (exdata string) {
	now := time.Now()
	num := now.Format("060102150405")
	exdata = "20" + num
	return exdata
}

/**
 *  取业务类型：最大可支持10个长度的英文字母、数字组合的字符串， 用户
 *  可以改写里面的内容。
 *  @return 返回业务类型
 */
func GetSvrtype() string {
	return "SM0001"
}

/**
 *  将密码按格式加密
 *	@param userid 用户名，pwd 密码，strtime 时间串 如时间是2017-5-10 11:20:33, 则串的内容为0510112033
 *  @return 返回加密后的字符串
 */

func CryptPwd(userid string, pwd string, strtime string) string {
	// 密码加密模式：账号：J10003, 密码：111111, 固定字符串：00000000
	// 时间戳：0803192020
	// MD5加密之前的对应字符串：J10003000000001111110803192020
	// MD5加密之后的密码字符串：26dad7f364507df18f3841cc9c4ff94d
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(userid + "00000000" + pwd + strtime))
	encryptPwd := md5Ctx.Sum(nil)
	pwdmd5 := hex.EncodeToString(encryptPwd[:])
	return pwdmd5
}

/**
 *  将数据打包在JSON格式
 *	@param pkg 数据
 *  @return 返回JSON格式的字符数组
 */
func PkgToJson(pkg interface{}) []byte {
	if nil == pkg {
		return nil
	}

	ret, err := json.Marshal(pkg)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}
	return ret
}

/**
 *  短连接发送http post 请求，并返回结果。
 *	@param url 请求的URL, data发送的数据, content_type数据格式
 *  @return http返回码， 内容，出错信息
 */
func HttpPostOnce(url string, data []byte, content_type string) (int, []byte, error) {

	body := strings.NewReader(string(data[:]))
	client := &http.Client{}
	reqest, _ := http.NewRequest("POST", url, body)
	reqest.Header.Set("Accept-Encoding", "gzip, deflate")
	reqest.Header.Set("Content-Type", content_type)
	reqest.Header.Set("Connection", "Close")

	response, err := client.Do(reqest)
	if err != nil {
		return 0, nil, err
	}
	defer response.Body.Close()
	resbody, _ := ioutil.ReadAll(response.Body)
	//bodystr := string(body)
	return response.StatusCode, resbody, err
}

/**
 *  移出指定的字符串数组成员
 *	@param s 字符串数组， element要移出的成员
 *  @return true 返回移出后的数组
 */
func RemoveAddressElement(s []string, element string) []string {
	ret := make([]string, 0)
	for _, v := range s {
		if v != element {
			ret = append(ret, v)
		}
	}
	return ret
}

/**
 *  HTTP POST 短连接发送指定类型的数据并接收返回数据
 *	@param postobj 指定类型
 *  @return true 如果成功返回TRUE,不成功返回FALSE
 */
type PostObj interface {
	GetUserid() string
	GetPwd() string
	GetTimestemp() string
	GetName() string
	ParseRecvData([]byte) bool
}

func SendAndRecvOnce(postobj PostObj) bool {

	fmt.Println("SendAndRecvOnce", postobj.GetName(), "begin")
	// 获取发送IP
	fmt.Println("SendAndRecvOnce", postobj.GetName(), ", get send address begin")
	checkaddr := &CheckAddress{Cm: GetConfigManager()}

	addr := checkaddr.GetAddressByUserID(postobj.GetUserid(), postobj.GetPwd(), postobj.GetTimestemp())
	beego.Debug(postobj.GetUserid())
	beego.Debug(postobj.GetPwd())
	if "" == addr {
		fmt.Println("SendAndRecvOnce", postobj.GetName(), "GetAddressByUserID failed.")
		//return false
	}
	fmt.Println("SendAndRecvOnce", postobj.GetName(), ", get send address is", addr)

	// 获取发送URL
	requestHost := "http://" + addr
	url := requestHost + checkaddr.Cm.RequestPath + postobj.GetName()
	url = "http://api02.monyun.cn:7901/sms/v2/std/single_send"
	// 组包
	data := PkgToJson(postobj)

	fmt.Println("SendAndRecvOnce", postobj.GetName(), ", body is", string(data[:]))

	rep_code, rep_body, err := HttpPostOnce(url, data, "application/json")
	beego.Debug(rep_code)
	beego.Debug(string(rep_body))
	if nil != err || 200 != rep_code {
		fmt.Println("SendAndRecvOnce", postobj.GetName(), "Post data failed, url:", url, "data:", string(data[:]), ", rep_code:", rep_code, ", body:", rep_body, ", err:", err)
		return false
	}

	ret := postobj.ParseRecvData(rep_body)
	return ret
}

/**
 *  将域名格式检查
 *	@param doMain 域名
 *  @return true 如果成功返回TRUE,不成功返回FALSE
 */
func CheckDomain(addr string) bool {
	matched, err := regexp.MatchString("[a-zA-Z0-9][-a-zA-Z0-9]{0,62}((\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+)", addr)
	if nil != err {
		return false
	}

	return matched
}

/**
 *  将域名替换成IP
 *	@param doMain 域名
 *  @return ip, err 如果成功返回IP, 1, 解析不成功返回 addr, 0, 格式错误 "", -1
 */
func AddressToIPAndPort(addr string) (string, int) {
	// 地址形如 www.domain.com:7901或者123.123.13.11:7901
	ipandport := strings.Split(addr, ":")
	if nil != net.ParseIP(ipandport[0]) {
		// 是IP地址，直接返回
		return addr, 0
	}

	if false == CheckDomain(ipandport[0]) {
		Logger.Printf("不符合域名格式，可能是输入错误，请检查, 输入：%s。", addr)
		return "", -1
	}

	// 解析域名，获取IP列表
	iplist, err := net.LookupIP(ipandport[0])
	if nil != err {
		// 域名IP列表失败, 返回失败。
		Logger.Printf("域名解析失败，输入为：%s, %s", addr, err)
		return addr, 0
	}

	Logger.Printf("域名解析成功，输入域名：%s, IP：%s", addr, iplist[0].String())
	// 判断是否有端口信息, 如果有就用第一个IP替换域名，如果没有直接返回IP
	if len(ipandport) > 1 {
		addr = iplist[0].String() + ":" + ipandport[1]
	} else {
		addr = iplist[0].String()
	}

	return addr, 1
}
