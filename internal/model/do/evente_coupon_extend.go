// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeCouponExtend is the golang structure of table evente_coupon_extend for DAO operations like Where/Data.
type EventeCouponExtend struct {
	g.Meta        `orm:"table:evente_coupon_extend, do:true"`
	Id            interface{} //
	OrgId         interface{} //
	CouponId      interface{} // 卡券id
	ApplyType     interface{} // 适用方式 同主表 2指定参与 3指定不参与
	ProductType   interface{} // 产品类型 （event活动 goods商品 hotel酒店 scenic景点）
	ProductId     interface{} // 产品id
	AllScreenings interface{} // 是否全部场次 0 否 1 是
	ScreeningsId  interface{} // 场次id集合,以,号分割，只在product_type=event&&all_screenings=0 时有效
	Status        interface{} // 1-正常，2-删除
	CreateDate    *gtime.Time // 新增时间
	UpdateDate    *gtime.Time // 更新时间
}
