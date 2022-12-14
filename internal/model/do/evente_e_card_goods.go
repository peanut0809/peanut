// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeECardGoods is the golang structure of table evente_e_card_goods for DAO operations like Where/Data.
type EventeECardGoods struct {
	g.Meta       `orm:"table:evente_e_card_goods, do:true"`
	Id           interface{} //
	OrgId        interface{} // 主办id
	CardId       interface{} // 年卡id
	Name         interface{} // 商品名称
	ImgId        interface{} // 商品图片ID
	Price        interface{} // 商品原价
	GiveNumber   interface{} // 赠送数量
	FetchedCount interface{} // 已经领取的次数
	Status       interface{} // 1-正常，2-删除,3-手动关闭
	CreateDate   *gtime.Time // 新增时间
	UpdateDate   *gtime.Time // 更新时间
}
