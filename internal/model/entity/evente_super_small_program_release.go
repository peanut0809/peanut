// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeSuperSmallProgramRelease is the golang structure for table evente_super_small_program_release.
type EventeSuperSmallProgramRelease struct {
	Id             int         `json:"id"             description:""`
	AppId          string      `json:"appId"          description:"app_id"`
	TemplateId     int         `json:"templateId"     description:"代码库中的代码模版ID"`
	ExtJson        string      `json:"extJson"        description:"提交审核代码的扩展内容json"`
	UserVersion    string      `json:"userVersion"    description:"当前提交版本"`
	UserDesc       string      `json:"userDesc"       description:"当前提交简介"`
	SubmitItemList string      `json:"submitItemList" description:"提交审核项的一个列表"`
	AuditId        int         `json:"auditId"        description:"审核编号"`
	AuditStatus    int         `json:"auditStatus"    description:"0为审核成功，1为审核失败，2为审核中"`
	Reason         string      `json:"reason"         description:"当status=1，审核被拒绝时，返回的拒绝原因"`
	ReleaseStatus  int         `json:"releaseStatus"  description:"发布状态 0为未发布，1为审核成功后发布成功"`
	AuditDate      *gtime.Time `json:"auditDate"      description:"最后一次审核通知或查询时间"`
	ReleaseDate    *gtime.Time `json:"releaseDate"    description:"发布时间"`
	CreateDate     *gtime.Time `json:"createDate"     description:"新增时间"`
	UpdateDate     *gtime.Time `json:"updateDate"     description:"更新时间"`
}
