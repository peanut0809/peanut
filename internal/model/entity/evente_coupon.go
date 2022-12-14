// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeCoupon is the golang structure for table evente_coupon.
type EventeCoupon struct {
	Id                 uint64      `json:"id"                 description:""`
	OrgId              int64       `json:"orgId"              description:"主办id"`
	Name               string      `json:"name"               description:"优惠券名称"`
	CouponType         int         `json:"couponType"         description:"优惠券类型 1代金券 2:折扣券"`
	Price              float64     `json:"price"              description:"优惠券面额"`
	DiscountPercent    float64     `json:"discountPercent"    description:"折扣比例"`
	UseConditionSwitch int         `json:"useConditionSwitch" description:"是否限制使用条件 0不限制 1限制"`
	UseCondition       float64     `json:"useCondition"       description:"使用条件 订单金额满这个数才可使用"`
	Number             int         `json:"number"             description:"发行量"`
	EventApplyType     int         `json:"eventApplyType"     description:"适用活动方式 1全部活动 2指定活动参与 3指定活动不参与"`
	GoodsApplyType     int         `json:"goodsApplyType"     description:"适用商品方式 1全部商品 2指定商品参与 3指定商品不参与"`
	FetchedNumber      int         `json:"fetchedNumber"      description:"已经领取的数量"`
	FetchTimeSwitch    int         `json:"fetchTimeSwitch"    description:"是否限制领取时间 0不限制 1限制"`
	FetchTimeInfo      string      `json:"fetchTimeInfo"      description:"领取时间限制序列化 兼容前数据"`
	FetchLimit         uint        `json:"fetchLimit"         description:"每人领取限额"`
	FetchEntry         string      `json:"fetchEntry"         description:"领取入口 逗号分隔 1详情页"`
	FetchMemberSwitch  int         `json:"fetchMemberSwitch"  description:"是否限制领取人等级 0不限制 1限制"`
	FetchMemberLevel   string      `json:"fetchMemberLevel"   description:"等级限制 逗号分隔"`
	Status             int         `json:"status"             description:"1-正常，2-删除,3-手动关闭"`
	CreateDate         *gtime.Time `json:"createDate"         description:"新增时间"`
	UpdateDate         *gtime.Time `json:"updateDate"         description:"更新时间"`
}
