// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeECardGoods is the golang structure for table evente_e_card_goods.
type EventeECardGoods struct {
	Id           uint64      `json:"id"           description:""`
	OrgId        int64       `json:"orgId"        description:"主办id"`
	CardId       int64       `json:"cardId"       description:"年卡id"`
	Name         string      `json:"name"         description:"商品名称"`
	ImgId        int         `json:"imgId"        description:"商品图片ID"`
	Price        float64     `json:"price"        description:"商品原价"`
	GiveNumber   int         `json:"giveNumber"   description:"赠送数量"`
	FetchedCount int         `json:"fetchedCount" description:"已经领取的次数"`
	Status       int         `json:"status"       description:"1-正常，2-删除,3-手动关闭"`
	CreateDate   *gtime.Time `json:"createDate"   description:"新增时间"`
	UpdateDate   *gtime.Time `json:"updateDate"   description:"更新时间"`
}
