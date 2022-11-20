// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EventeRegimentStockDao is the data access object for table evente_regiment_stock.
type EventeRegimentStockDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns EventeRegimentStockColumns // columns contains all the column names of Table for convenient usage.
}

// EventeRegimentStockColumns defines and stores column names for table evente_regiment_stock.
type EventeRegimentStockColumns struct {
	Id              string //
	OrgId           string // 主办id
	RegimentId      string // 团id
	ProductId       string // 产品id
	ProductTitle    string // 产品名称
	ProductSubId    string // 二级id  例如场次id
	ProductSubTitle string // 二级名称  例如场次名称
	ProductRelId    string // 三级id  例如 票品 规格id
	ProductRelTitle string // 三级名称 例如 票品 规格名称
	ProductType     string // 产品类型（event 票务 goods 商品）
	Money           string // 实售价格
	ActivityNum     string // 拼团库存
	IsJoin          string // 是否参加 拼团 1 - 参加 0 - 未参加
	CreateDate      string // 创建时间
	UpdateDate      string // 修改时间
}

// eventeRegimentStockColumns holds the columns for table evente_regiment_stock.
var eventeRegimentStockColumns = EventeRegimentStockColumns{
	Id:              "id",
	OrgId:           "org_id",
	RegimentId:      "regiment_id",
	ProductId:       "product_id",
	ProductTitle:    "product_title",
	ProductSubId:    "product_sub_id",
	ProductSubTitle: "product_sub_title",
	ProductRelId:    "product_rel_id",
	ProductRelTitle: "product_rel_title",
	ProductType:     "product_type",
	Money:           "money",
	ActivityNum:     "activity_num",
	IsJoin:          "is_join",
	CreateDate:      "create_date",
	UpdateDate:      "update_date",
}

// NewEventeRegimentStockDao creates and returns a new DAO object for table data access.
func NewEventeRegimentStockDao() *EventeRegimentStockDao {
	return &EventeRegimentStockDao{
		group:   "default",
		table:   "evente_regiment_stock",
		columns: eventeRegimentStockColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EventeRegimentStockDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EventeRegimentStockDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EventeRegimentStockDao) Columns() EventeRegimentStockColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EventeRegimentStockDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EventeRegimentStockDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EventeRegimentStockDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}