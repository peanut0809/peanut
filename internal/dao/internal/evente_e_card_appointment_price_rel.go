// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EventeECardAppointmentPriceRelDao is the data access object for table evente_e_card_appointment_price_rel.
type EventeECardAppointmentPriceRelDao struct {
	table   string                                // table is the underlying table name of the DAO.
	group   string                                // group is the database configuration group name of current DAO.
	columns EventeECardAppointmentPriceRelColumns // columns contains all the column names of Table for convenient usage.
}

// EventeECardAppointmentPriceRelColumns defines and stores column names for table evente_e_card_appointment_price_rel.
type EventeECardAppointmentPriceRelColumns struct {
	Id            string // 自增id
	IncrementId   string // 主单号
	AppointmentId string // 预约主表id
	OrgId         string //
	UserId        string //
	PriceId       string // 票品
	PriceName     string //
	Price         string //
	SeatsAttr     string // 座位信息(json 结构化数据)
	Number        string // 票品数量
	CreateDate    string // (预约时间)创建时间
	UpdateDate    string // 修改时间
}

// eventeECardAppointmentPriceRelColumns holds the columns for table evente_e_card_appointment_price_rel.
var eventeECardAppointmentPriceRelColumns = EventeECardAppointmentPriceRelColumns{
	Id:            "id",
	IncrementId:   "increment_id",
	AppointmentId: "appointment_id",
	OrgId:         "org_id",
	UserId:        "user_id",
	PriceId:       "price_id",
	PriceName:     "price_name",
	Price:         "price",
	SeatsAttr:     "seats_attr",
	Number:        "number",
	CreateDate:    "create_date",
	UpdateDate:    "update_date",
}

// NewEventeECardAppointmentPriceRelDao creates and returns a new DAO object for table data access.
func NewEventeECardAppointmentPriceRelDao() *EventeECardAppointmentPriceRelDao {
	return &EventeECardAppointmentPriceRelDao{
		group:   "default",
		table:   "evente_e_card_appointment_price_rel",
		columns: eventeECardAppointmentPriceRelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EventeECardAppointmentPriceRelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EventeECardAppointmentPriceRelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EventeECardAppointmentPriceRelDao) Columns() EventeECardAppointmentPriceRelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EventeECardAppointmentPriceRelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EventeECardAppointmentPriceRelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EventeECardAppointmentPriceRelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
