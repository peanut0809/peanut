// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeVoucherCodeOrderRel is the golang structure of table evente_voucher_code_order_rel for DAO operations like Where/Data.
type EventeVoucherCodeOrderRel struct {
	g.Meta              `orm:"table:evente_voucher_code_order_rel, do:true"`
	Id                  interface{} //
	OrgId               interface{} //
	UserId              interface{} // 用户id
	VoucherActivationId interface{} // 优惠码明细ID
	VoucherCodeId       interface{} // 主表ID
	IncrementId         interface{} // 订单increment_id
	PayStatus           interface{} // 支付状态 0 未支付， 1已支付 2 退款
	OrderMoney          interface{} // 订单金额
	PayMoney            interface{} // 支付金额
	VoucherCodeMoney    interface{} // 优惠金额
	CreateDate          *gtime.Time // 新增时间
	UpdateDate          *gtime.Time // 更新时间
	ActivationCode      interface{} // 优惠码
	UserType            interface{} // 1新用户 2 老用户
	UserName            interface{} //
	Mobile              interface{} // 手机号
	CheckStatus         interface{} // 核销状态 0 未核销 1 核销成功  2 核销失败
	CheckMesaage        interface{} // 核销成功 失败原因
}
