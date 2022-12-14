// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeECardAppointment is the golang structure of table evente_e_card_appointment for DAO operations like Where/Data.
type EventeECardAppointment struct {
	g.Meta              `orm:"table:evente_e_card_appointment, do:true"`
	Id                  interface{} // 自增id
	IncrementId         interface{} //
	UserId              interface{} // C端用户id
	OrgId               interface{} // 主办id
	Mobile              interface{} // 手机号
	Name                interface{} // 预约联系人
	CardId              interface{} // E通卡id
	CardSaleId          interface{} //
	PicId               interface{} // 活动图片id
	EventeId            interface{} // 活动id
	EventeName          interface{} // 活动名称
	ScreeningsId        interface{} // 场次id
	ScreeningsName      interface{} // 场次名称
	ScreeningsStartDate *gtime.Time // 场次开始时间
	ScreeningsEndDate   *gtime.Time // 场次结束时间
	Number              interface{} // 预约数量
	Status              interface{} // 状态 1预约 2取消预约 3已核销
	VerificationName    interface{} // 核销人-主办名称
	VerificationDate    *gtime.Time // 核销时间
	CreateDate          *gtime.Time // 创建时间
	UpdateDate          *gtime.Time // 修改时间
	UseNumber           interface{} // 本次消耗年卡次数
}
