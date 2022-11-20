// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"golang8090/internal/dao/internal"
)

// internalEventeReductionProductEventDao is internal type for wrapping internal DAO implements.
type internalEventeReductionProductEventDao = *internal.EventeReductionProductEventDao

// eventeReductionProductEventDao is the data access object for table evente_reduction_product_event.
// You can define custom methods on it to extend its functionality as you wish.
type eventeReductionProductEventDao struct {
	internalEventeReductionProductEventDao
}

var (
	// EventeReductionProductEvent is globally public accessible object for table evente_reduction_product_event operations.
	EventeReductionProductEvent = eventeReductionProductEventDao{
		internal.NewEventeReductionProductEventDao(),
	}
)

// Fill with you ideas below.
