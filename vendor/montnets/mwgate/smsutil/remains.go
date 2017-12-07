package smsutil

/**
 *
 * @功能概要：余额对象
 * @项目名称： GoSmsSdk5.3
 * @初创作者： wangdongyu
 * @公司名称： ShenZhen Montnets Technology CO.,LTD.
 * @创建时间： 2017-4-28 上午10:13
 * <p>修改记录1：</p>
 * <pre>
 *      修改日期：
 *      修改人：
 *      修改内容：
 * </pre>
 */
type Remains struct {
	//是否成功的标识   0:成功;其他则为错误码
	result int

	//计费类型  0:按条计费;1:按金额计费
	chargetype int

	//剩余条数  chargetype为1时,balance为0.
	blance int

	//剩余金额   chargetype为0时,money为0.
	money string
}

func NewRemains() *Remains {
	return &Remains{result: -310099, chargetype: 0, blance: 0, money: "0"}
}
