// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeECardActivation is the golang structure of table evente_e_card_activation for DAO operations like Where/Data.
type EventeECardActivation struct {
	g.Meta           `orm:"table:evente_e_card_activation, do:true"`
	Id               interface{} // 自增id
	OrgId            interface{} // 主办id
	CardId           interface{} // E通卡id
	Number           interface{} // 生成数量
	ActivationNumber interface{} // 激活数量
	CreateDate       *gtime.Time // 生成时间(创建时间)
	UpdateDate       *gtime.Time // 更新时间
}
