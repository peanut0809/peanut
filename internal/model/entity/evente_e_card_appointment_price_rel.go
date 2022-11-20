// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeECardAppointmentPriceRel is the golang structure for table evente_e_card_appointment_price_rel.
type EventeECardAppointmentPriceRel struct {
	Id            uint        `json:"id"            description:"自增id"`
	IncrementId   string      `json:"incrementId"   description:"主单号"`
	AppointmentId uint        `json:"appointmentId" description:"预约主表id"`
	OrgId         int         `json:"orgId"         description:""`
	UserId        int         `json:"userId"        description:""`
	PriceId       int         `json:"priceId"       description:"票品"`
	PriceName     string      `json:"priceName"     description:""`
	Price         float64     `json:"price"         description:""`
	SeatsAttr     string      `json:"seatsAttr"     description:"座位信息(json 结构化数据)"`
	Number        int         `json:"number"        description:"票品数量"`
	CreateDate    *gtime.Time `json:"createDate"    description:"(预约时间)创建时间"`
	UpdateDate    *gtime.Time `json:"updateDate"    description:"修改时间"`
}