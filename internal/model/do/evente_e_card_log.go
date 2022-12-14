// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeECardLog is the golang structure of table evente_e_card_log for DAO operations like Where/Data.
type EventeECardLog struct {
	g.Meta         `orm:"table:evente_e_card_log, do:true"`
	Id             interface{} // 主键ID
	OrgId          interface{} // 主办ID
	BeforeJsonData interface{} // 修改之前操作记录数据
	AfterJsonData  interface{} // 修改之后操作记录数据
	CreateDate     *gtime.Time // 创建时间
}
