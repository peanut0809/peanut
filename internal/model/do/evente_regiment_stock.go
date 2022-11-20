// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeRegimentStock is the golang structure of table evente_regiment_stock for DAO operations like Where/Data.
type EventeRegimentStock struct {
	g.Meta          `orm:"table:evente_regiment_stock, do:true"`
	Id              interface{} //
	OrgId           interface{} // 主办id
	RegimentId      interface{} // 团id
	ProductId       interface{} // 产品id
	ProductTitle    interface{} // 产品名称
	ProductSubId    interface{} // 二级id  例如场次id
	ProductSubTitle interface{} // 二级名称  例如场次名称
	ProductRelId    interface{} // 三级id  例如 票品 规格id
	ProductRelTitle interface{} // 三级名称 例如 票品 规格名称
	ProductType     interface{} // 产品类型（event 票务 goods 商品）
	Money           interface{} // 实售价格
	ActivityNum     interface{} // 拼团库存
	IsJoin          interface{} // 是否参加 拼团 1 - 参加 0 - 未参加
	CreateDate      *gtime.Time // 创建时间
	UpdateDate      *gtime.Time // 修改时间
}