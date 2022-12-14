// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"golang8090/internal/dao/internal"
)

// internalEventeRegimentExtendDao is internal type for wrapping internal DAO implements.
type internalEventeRegimentExtendDao = *internal.EventeRegimentExtendDao

// eventeRegimentExtendDao is the data access object for table evente_regiment_extend.
// You can define custom methods on it to extend its functionality as you wish.
type eventeRegimentExtendDao struct {
	internalEventeRegimentExtendDao
}

var (
	// EventeRegimentExtend is globally public accessible object for table evente_regiment_extend operations.
	EventeRegimentExtend = eventeRegimentExtendDao{
		internal.NewEventeRegimentExtendDao(),
	}
)

// Fill with you ideas below.
