// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeSmallProgramRelease is the golang structure of table evente_small_program_release for DAO operations like Where/Data.
type EventeSmallProgramRelease struct {
	g.Meta         `orm:"table:evente_small_program_release, do:true"`
	Id             interface{} //
	OrgId          interface{} // 主办ID
	AppId          interface{} // app_id
	TemplateId     interface{} // 代码库中的代码模版ID
	ExtJson        interface{} // 提交审核代码的扩展内容json
	UserVersion    interface{} // 当前提交版本
	UserDesc       interface{} // 当前提交简介
	SubmitItemList interface{} // 提交审核项的一个列表
	AuditId        interface{} // 审核编号
	AuditStatus    interface{} // 0为审核成功，1为审核失败，2为审核中
	Reason         interface{} // 当status=1，审核被拒绝时，返回的拒绝原因
	ReleaseStatus  interface{} // 发布状态 0为未发布，1为审核成功后发布成功
	AuditDate      *gtime.Time // 最后一次审核通知或查询时间
	ReleaseDate    *gtime.Time // 发布时间
	CreateDate     *gtime.Time // 新增时间
	UpdateDate     *gtime.Time // 更新时间
}
