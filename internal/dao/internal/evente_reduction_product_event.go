// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EventeReductionProductEventDao is the data access object for table evente_reduction_product_event.
type EventeReductionProductEventDao struct {
	table   string                             // table is the underlying table name of the DAO.
	group   string                             // group is the database configuration group name of current DAO.
	columns EventeReductionProductEventColumns // columns contains all the column names of Table for convenient usage.
}

// EventeReductionProductEventColumns defines and stores column names for table evente_reduction_product_event.
type EventeReductionProductEventColumns struct {
	Id           string //
	OrgId        string // 主办ID
	ReductionId  string // 满减ID
	ProductId    string // 产品ID
	ScreeningsId string // 场次ID
	PriceList    string // 票品列表集合（逗号分隔）
	Status       string // 状态（1.正常 ，2.删除）
	CreateDate   string // 创建时间
	UpdateDate   string // 更新时间
}

// eventeReductionProductEventColumns holds the columns for table evente_reduction_product_event.
var eventeReductionProductEventColumns = EventeReductionProductEventColumns{
	Id:           "id",
	OrgId:        "org_id",
	ReductionId:  "reduction_id",
	ProductId:    "product_id",
	ScreeningsId: "screenings_id",
	PriceList:    "price_list",
	Status:       "status",
	CreateDate:   "create_date",
	UpdateDate:   "update_date",
}

// NewEventeReductionProductEventDao creates and returns a new DAO object for table data access.
func NewEventeReductionProductEventDao() *EventeReductionProductEventDao {
	return &EventeReductionProductEventDao{
		group:   "default",
		table:   "evente_reduction_product_event",
		columns: eventeReductionProductEventColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EventeReductionProductEventDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EventeReductionProductEventDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EventeReductionProductEventDao) Columns() EventeReductionProductEventColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EventeReductionProductEventDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EventeReductionProductEventDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EventeReductionProductEventDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
