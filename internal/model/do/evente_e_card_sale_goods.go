// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeECardSaleGoods is the golang structure of table evente_e_card_sale_goods for DAO operations like Where/Data.
type EventeECardSaleGoods struct {
	g.Meta      `orm:"table:evente_e_card_sale_goods, do:true"`
	Id          interface{} // 自增id
	OrgId       interface{} // 主办id
	UserId      interface{} // 用户id
	IncrementId interface{} // 订单id
	SaleId      interface{} // sale_id
	CardId      interface{} // E通卡id
	GoodsId     interface{} // 商品id
	GiveNumber  interface{} // 赠送数量
	UseNumber   interface{} // 已使用数量
	CreateDate  *gtime.Time // 创建时间
	UpdateDate  *gtime.Time // 更新时间
}