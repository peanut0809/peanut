// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EventeCouponDao is the data access object for table evente_coupon.
type EventeCouponDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns EventeCouponColumns // columns contains all the column names of Table for convenient usage.
}

// EventeCouponColumns defines and stores column names for table evente_coupon.
type EventeCouponColumns struct {
	Id                 string //
	OrgId              string // 主办id
	Name               string // 优惠券名称
	CouponType         string // 优惠券类型 1代金券 2:折扣券
	Price              string // 优惠券面额
	DiscountPercent    string // 折扣比例
	UseConditionSwitch string // 是否限制使用条件 0不限制 1限制
	UseCondition       string // 使用条件 订单金额满这个数才可使用
	Number             string // 发行量
	EventApplyType     string // 适用活动方式 1全部活动 2指定活动参与 3指定活动不参与
	GoodsApplyType     string // 适用商品方式 1全部商品 2指定商品参与 3指定商品不参与
	FetchedNumber      string // 已经领取的数量
	FetchTimeSwitch    string // 是否限制领取时间 0不限制 1限制
	FetchTimeInfo      string // 领取时间限制序列化 兼容前数据
	FetchLimit         string // 每人领取限额
	FetchEntry         string // 领取入口 逗号分隔 1详情页
	FetchMemberSwitch  string // 是否限制领取人等级 0不限制 1限制
	FetchMemberLevel   string // 等级限制 逗号分隔
	Status             string // 1-正常，2-删除,3-手动关闭
	CreateDate         string // 新增时间
	UpdateDate         string // 更新时间
}

// eventeCouponColumns holds the columns for table evente_coupon.
var eventeCouponColumns = EventeCouponColumns{
	Id:                 "id",
	OrgId:              "org_id",
	Name:               "name",
	CouponType:         "coupon_type",
	Price:              "price",
	DiscountPercent:    "discount_percent",
	UseConditionSwitch: "use_condition_switch",
	UseCondition:       "use_condition",
	Number:             "number",
	EventApplyType:     "event_apply_type",
	GoodsApplyType:     "goods_apply_type",
	FetchedNumber:      "fetched_number",
	FetchTimeSwitch:    "fetch_time_switch",
	FetchTimeInfo:      "fetch_time_info",
	FetchLimit:         "fetch_limit",
	FetchEntry:         "fetch_entry",
	FetchMemberSwitch:  "fetch_member_switch",
	FetchMemberLevel:   "fetch_member_level",
	Status:             "status",
	CreateDate:         "create_date",
	UpdateDate:         "update_date",
}

// NewEventeCouponDao creates and returns a new DAO object for table data access.
func NewEventeCouponDao() *EventeCouponDao {
	return &EventeCouponDao{
		group:   "default",
		table:   "evente_coupon",
		columns: eventeCouponColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EventeCouponDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EventeCouponDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EventeCouponDao) Columns() EventeCouponColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EventeCouponDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EventeCouponDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EventeCouponDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
