// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeRegiment is the golang structure for table evente_regiment.
type EventeRegiment struct {
	Id              uint        `json:"id"              description:"主键id"`
	OrgId           uint        `json:"orgId"           description:"主办id"`
	Title           string      `json:"title"           description:"名称"`
	RegimentType    int         `json:"regimentType"    description:"拼团类型 0 - 普通团 1 - 新人团 2 - 抽奖团"`
	StartDate       *gtime.Time `json:"startDate"       description:"开始时间"`
	EndDate         *gtime.Time `json:"endDate"         description:"结束时间"`
	RegimentPeople  int         `json:"regimentPeople"  description:"成团人数"`
	ExpiryDate      int         `json:"expiryDate"      description:"有效期"`
	ProductId       int         `json:"productId"       description:"产品id"`
	ProductType     string      `json:"productType"     description:"产品类型 event - 活动 goods - 商品"`
	LimitNum        int         `json:"limitNum"        description:"单次购买上限"`
	OpenRegimentNum int         `json:"openRegimentNum" description:"开团次数"`
	JoinRegimentNum int         `json:"joinRegimentNum" description:"参团次数"`
	RegimentStatus  int         `json:"regimentStatus"  description:"拼团状态 1 - 正常 0 - 关闭"`
	IsDel           uint        `json:"isDel"           description:"逻辑删除状态：默认 0 未删除  1 - 已删除"`
	CreateType      int         `json:"createType"      description:"创建类型 0 - h5 1 - 小程序"`
	CreateDate      *gtime.Time `json:"createDate"      description:"创建时间"`
	DeleteDate      *gtime.Time `json:"deleteDate"      description:"删除时间"`
	UpdateDate      *gtime.Time `json:"updateDate"      description:"修改时间"`
	EndHandleStatus int         `json:"endHandleStatus" description:"活动结束处理状态 -1 有待支付 0 未处理 1 已处理"`
}
