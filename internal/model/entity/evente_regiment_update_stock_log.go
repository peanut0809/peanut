// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeRegimentUpdateStockLog is the golang structure for table evente_regiment_update_stock_log.
type EventeRegimentUpdateStockLog struct {
	Id               uint        `json:"id"               description:""`
	OrgId            int         `json:"orgId"            description:"主办id"`
	UserId           int         `json:"userId"           description:"主办id"`
	RegimentId       uint        `json:"regimentId"       description:"团id"`
	IncrementId      string      `json:"incrementId"      description:"主订单id"`
	GroupIncrementId string      `json:"groupIncrementId" description:"团订单id"`
	ProductId        int         `json:"productId"        description:"产品id"`
	ProductSubId     int         `json:"productSubId"     description:"二级id  例如场次id"`
	ProductRelId     string      `json:"productRelId"     description:"三级id 例如票品 规格id"`
	StockId          int         `json:"stockId"          description:"库存id"`
	Money            float64     `json:"money"            description:"实售价格"`
	UpdateNum        uint        `json:"updateNum"        description:"扣减数量"`
	IsRollback       uint        `json:"isRollback"       description:"是否回滚"`
	CreateDate       *gtime.Time `json:"createDate"       description:"创建时间"`
	UpdateDate       *gtime.Time `json:"updateDate"       description:"修改时间"`
}