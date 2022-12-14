// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeIntegralOrderList is the golang structure for table evente_integral_order_list.
type EventeIntegralOrderList struct {
	Id            int         `json:"id"            description:"自增id"`
	OrgId         int         `json:"orgId"         description:"主办id"`
	UserId        int         `json:"userId"        description:"用户id"`
	IntegralId    int         `json:"integralId"    description:"积分抵现id"`
	IncrementId   string      `json:"incrementId"   description:"订单id"`
	Username      string      `json:"username"      description:""`
	Code          int         `json:"code"          description:"手机区号"`
	Phone         string      `json:"phone"         description:"手机号"`
	Status        int         `json:"status"        description:"订单状态: 10 待支付 20 已支付 30已发货 40 订单完成 50订单超时 60订单取消 70订单关闭 80 订单异常"`
	IntegralUsage int64       `json:"integralUsage" description:"订单积分使用量"`
	TotalAmount   float64     `json:"totalAmount"   description:"抵现总金额"`
	AmountPayable float64     `json:"amountPayable" description:"适用商品应付金额"`
	DiscountRatio int         `json:"discountRatio" description:"当前订单积分抵现 订单折扣比例快照"`
	IntegralRatio int         `json:"integralRatio" description:"当前订单积分抵现 积分金额兑换比例"`
	CreateDate    *gtime.Time `json:"createDate"    description:""`
	UpdateDate    *gtime.Time `json:"updateDate"    description:""`
}
