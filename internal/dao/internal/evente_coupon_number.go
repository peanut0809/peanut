// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EventeCouponNumberDao is the data access object for table evente_coupon_number.
type EventeCouponNumberDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns EventeCouponNumberColumns // columns contains all the column names of Table for convenient usage.
}

// EventeCouponNumberColumns defines and stores column names for table evente_coupon_number.
type EventeCouponNumberColumns struct {
	Id             string //
	OrgId          string // 主办id
	UserId         string // 用户id
	CouponId       string // 卡券id
	Status         string // 1正常，2已删除，3已使用
	UseIncrementId string // 使用订单
	UseDate        string // 使用时间
	CouponSource   string // 优惠券来源  normal  普通领取  integral  积分赠送
	CreateDate     string // 新增时间
	UpdateDate     string // 更新时间
}

// eventeCouponNumberColumns holds the columns for table evente_coupon_number.
var eventeCouponNumberColumns = EventeCouponNumberColumns{
	Id:             "id",
	OrgId:          "org_id",
	UserId:         "user_id",
	CouponId:       "coupon_id",
	Status:         "status",
	UseIncrementId: "use_increment_id",
	UseDate:        "use_date",
	CouponSource:   "coupon_source",
	CreateDate:     "create_date",
	UpdateDate:     "update_date",
}

// NewEventeCouponNumberDao creates and returns a new DAO object for table data access.
func NewEventeCouponNumberDao() *EventeCouponNumberDao {
	return &EventeCouponNumberDao{
		group:   "default",
		table:   "evente_coupon_number",
		columns: eventeCouponNumberColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EventeCouponNumberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EventeCouponNumberDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EventeCouponNumberDao) Columns() EventeCouponNumberColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EventeCouponNumberDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EventeCouponNumberDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EventeCouponNumberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
