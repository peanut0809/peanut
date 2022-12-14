// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EventeOrgWechatAuthDao is the data access object for table evente_org_wechat_auth.
type EventeOrgWechatAuthDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns EventeOrgWechatAuthColumns // columns contains all the column names of Table for convenient usage.
}

// EventeOrgWechatAuthColumns defines and stores column names for table evente_org_wechat_auth.
type EventeOrgWechatAuthColumns struct {
	Id                     string //
	OrgId                  string // org_id
	NickName               string // 授权方昵称
	ServiceTypeInfo        string // 类型，0订阅号，1升级后订阅号，2服务号
	VerifyTypeInfo         string // 认证类型，-1未认证，0微信认证，1新浪微博认证，2腾讯微博认证，3:0+未通过名称认证，4:1+未通过名称认证，5:2+未通过名称认证
	UserName               string // 授权方公众号的原始ID
	PrincipalName          string // 公众号的主体名称
	Alias                  string // 公众号的主体名称
	HeadImg                string // 授权方头像
	QrcodeUrl              string // 二维码
	BusinessInfo           string // 商户功能 序列化
	AuthorizerAppid        string // app_id
	AuthorizerRefreshToken string //
	FuncInfo               string // 公众号授权给开发者的权限集列表 序列化
	AuthorizerInfo         string // 微信返回的序列化
	CreateDate             string // 新增时间
	UpdateDate             string // 更新时间
	DeleteDate             string // 删除时间
	Status                 string // 状态 1:正常 2:取消授权
	TemplateList           string //
}

// eventeOrgWechatAuthColumns holds the columns for table evente_org_wechat_auth.
var eventeOrgWechatAuthColumns = EventeOrgWechatAuthColumns{
	Id:                     "id",
	OrgId:                  "org_id",
	NickName:               "nick_name",
	ServiceTypeInfo:        "service_type_info",
	VerifyTypeInfo:         "verify_type_info",
	UserName:               "user_name",
	PrincipalName:          "principal_name",
	Alias:                  "alias",
	HeadImg:                "head_img",
	QrcodeUrl:              "qrcode_url",
	BusinessInfo:           "business_info",
	AuthorizerAppid:        "authorizer_appid",
	AuthorizerRefreshToken: "authorizer_refresh_token",
	FuncInfo:               "func_info",
	AuthorizerInfo:         "authorizer_info",
	CreateDate:             "create_date",
	UpdateDate:             "update_date",
	DeleteDate:             "delete_date",
	Status:                 "status",
	TemplateList:           "template_list",
}

// NewEventeOrgWechatAuthDao creates and returns a new DAO object for table data access.
func NewEventeOrgWechatAuthDao() *EventeOrgWechatAuthDao {
	return &EventeOrgWechatAuthDao{
		group:   "default",
		table:   "evente_org_wechat_auth",
		columns: eventeOrgWechatAuthColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EventeOrgWechatAuthDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EventeOrgWechatAuthDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EventeOrgWechatAuthDao) Columns() EventeOrgWechatAuthColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EventeOrgWechatAuthDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EventeOrgWechatAuthDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EventeOrgWechatAuthDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
