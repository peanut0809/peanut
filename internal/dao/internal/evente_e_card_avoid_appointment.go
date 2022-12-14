// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EventeECardAvoidAppointmentDao is the data access object for table evente_e_card_avoid_appointment.
type EventeECardAvoidAppointmentDao struct {
	table   string                             // table is the underlying table name of the DAO.
	group   string                             // group is the database configuration group name of current DAO.
	columns EventeECardAvoidAppointmentColumns // columns contains all the column names of Table for convenient usage.
}

// EventeECardAvoidAppointmentColumns defines and stores column names for table evente_e_card_avoid_appointment.
type EventeECardAvoidAppointmentColumns struct {
	Id                  string // 自增id
	UserId              string // C端用户id
	OrgId               string // 主办id
	CardId              string // E通卡id
	CardSaleId          string // E通卡销售表id
	PicId               string // 活动图片id
	EventeId            string // 活动id
	EventeName          string // 活动名称
	ScreeningsId        string // 场次id
	ScreeningsName      string // 场次名称
	ScreeningsStartDate string // 场次开始时间
	ScreeningsEndDate   string // 场次结束时间
	VerificationName    string // 核销人-主办名称
	CreateDate          string // 创建时间
	UpdateDate          string // 修改时间
}

// eventeECardAvoidAppointmentColumns holds the columns for table evente_e_card_avoid_appointment.
var eventeECardAvoidAppointmentColumns = EventeECardAvoidAppointmentColumns{
	Id:                  "id",
	UserId:              "user_id",
	OrgId:               "org_id",
	CardId:              "card_id",
	CardSaleId:          "card_sale_id",
	PicId:               "pic_id",
	EventeId:            "evente_id",
	EventeName:          "evente_name",
	ScreeningsId:        "screenings_id",
	ScreeningsName:      "screenings_name",
	ScreeningsStartDate: "screenings_start_date",
	ScreeningsEndDate:   "screenings_end_date",
	VerificationName:    "verification_name",
	CreateDate:          "create_date",
	UpdateDate:          "update_date",
}

// NewEventeECardAvoidAppointmentDao creates and returns a new DAO object for table data access.
func NewEventeECardAvoidAppointmentDao() *EventeECardAvoidAppointmentDao {
	return &EventeECardAvoidAppointmentDao{
		group:   "default",
		table:   "evente_e_card_avoid_appointment",
		columns: eventeECardAvoidAppointmentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EventeECardAvoidAppointmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EventeECardAvoidAppointmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EventeECardAvoidAppointmentDao) Columns() EventeECardAvoidAppointmentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EventeECardAvoidAppointmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EventeECardAvoidAppointmentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EventeECardAvoidAppointmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
