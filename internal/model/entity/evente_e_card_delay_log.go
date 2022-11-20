// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeECardDelayLog is the golang structure for table evente_e_card_delay_log.
type EventeECardDelayLog struct {
	Id         uint        `json:"id"         description:"自增id"`
	OrgId      uint        `json:"orgId"      description:"主办id"`
	ChildId    uint        `json:"childId"    description:"子账号id"`
	CardId     uint        `json:"cardId"     description:"E通卡id"`
	UserId     uint        `json:"userId"     description:"用户id"`
	CardSaleId uint        `json:"cardSaleId" description:"销售id"`
	OrgName    string      `json:"orgName"    description:"主办名称"`
	BeforeDate *gtime.Time `json:"beforeDate" description:"延期之前时间"`
	ExpireDate *gtime.Time `json:"expireDate" description:"到期时间"`
	CreateDate *gtime.Time `json:"createDate" description:"创建时间"`
}