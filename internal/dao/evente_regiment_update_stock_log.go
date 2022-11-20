// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"golang8090/internal/dao/internal"
)

// internalEventeRegimentUpdateStockLogDao is internal type for wrapping internal DAO implements.
type internalEventeRegimentUpdateStockLogDao = *internal.EventeRegimentUpdateStockLogDao

// eventeRegimentUpdateStockLogDao is the data access object for table evente_regiment_update_stock_log.
// You can define custom methods on it to extend its functionality as you wish.
type eventeRegimentUpdateStockLogDao struct {
	internalEventeRegimentUpdateStockLogDao
}

var (
	// EventeRegimentUpdateStockLog is globally public accessible object for table evente_regiment_update_stock_log operations.
	EventeRegimentUpdateStockLog = eventeRegimentUpdateStockLogDao{
		internal.NewEventeRegimentUpdateStockLogDao(),
	}
)

// Fill with you ideas below.
