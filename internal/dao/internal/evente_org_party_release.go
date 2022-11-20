// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EventeOrgPartyReleaseDao is the data access object for table evente_org_party_release.
type EventeOrgPartyReleaseDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns EventeOrgPartyReleaseColumns // columns contains all the column names of Table for convenient usage.
}

// EventeOrgPartyReleaseColumns defines and stores column names for table evente_org_party_release.
type EventeOrgPartyReleaseColumns struct {
	Id             string //
	OrgId          string // 主办ID
	AppId          string // app_id
	PartyType      string // 第三方工具类型
	TemplateId     string // 代码库中的代码模版ID
	ExtJson        string // 提交审核代码的扩展内容json
	UserVersion    string // 当前提交版本
	UserDesc       string // 当前提交简介
	SubmitItemList string // 提交审核项的一个列表
	AuditId        string // 审核编号
	AuditStatus    string // 0为审核成功，1为审核失败，2为审核中
	Reason         string // 当status=1，审核被拒绝时，返回的拒绝原因
	ReleaseStatus  string // 发布状态 0为未发布，1为审核成功后发布成功
	AuditDate      string // 最后一次审核通知或查询时间
	ReleaseDate    string // 发布时间
	CreateDate     string // 新增时间
	UpdateDate     string // 更新时间
}

// eventeOrgPartyReleaseColumns holds the columns for table evente_org_party_release.
var eventeOrgPartyReleaseColumns = EventeOrgPartyReleaseColumns{
	Id:             "id",
	OrgId:          "org_id",
	AppId:          "app_id",
	PartyType:      "party_type",
	TemplateId:     "template_id",
	ExtJson:        "ext_json",
	UserVersion:    "user_version",
	UserDesc:       "user_desc",
	SubmitItemList: "submit_item_list",
	AuditId:        "audit_id",
	AuditStatus:    "audit_status",
	Reason:         "reason",
	ReleaseStatus:  "release_status",
	AuditDate:      "audit_date",
	ReleaseDate:    "release_date",
	CreateDate:     "create_date",
	UpdateDate:     "update_date",
}

// NewEventeOrgPartyReleaseDao creates and returns a new DAO object for table data access.
func NewEventeOrgPartyReleaseDao() *EventeOrgPartyReleaseDao {
	return &EventeOrgPartyReleaseDao{
		group:   "default",
		table:   "evente_org_party_release",
		columns: eventeOrgPartyReleaseColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EventeOrgPartyReleaseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EventeOrgPartyReleaseDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EventeOrgPartyReleaseDao) Columns() EventeOrgPartyReleaseColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EventeOrgPartyReleaseDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EventeOrgPartyReleaseDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EventeOrgPartyReleaseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
