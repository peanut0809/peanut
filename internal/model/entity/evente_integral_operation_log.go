// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeIntegralOperationLog is the golang structure for table evente_integral_operation_log.
type EventeIntegralOperationLog struct {
	Id             int         `json:"id"             description:"自增id"`
	OrgId          int         `json:"orgId"          description:"主办id"`
	BeforeJsonData string      `json:"beforeJsonData" description:"修改之前操作记录数据"`
	AfterJsonData  string      `json:"afterJsonData"  description:"修改之后操作记录数据"`
	CreateTime     *gtime.Time `json:"createTime"     description:"创建时间"`
}
