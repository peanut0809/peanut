// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"golang8090/internal/dao/internal"
)

// internalEventeSmallProgramReleaseDao is internal type for wrapping internal DAO implements.
type internalEventeSmallProgramReleaseDao = *internal.EventeSmallProgramReleaseDao

// eventeSmallProgramReleaseDao is the data access object for table evente_small_program_release.
// You can define custom methods on it to extend its functionality as you wish.
type eventeSmallProgramReleaseDao struct {
	internalEventeSmallProgramReleaseDao
}

var (
	// EventeSmallProgramRelease is globally public accessible object for table evente_small_program_release operations.
	EventeSmallProgramRelease = eventeSmallProgramReleaseDao{
		internal.NewEventeSmallProgramReleaseDao(),
	}
)

// Fill with you ideas below.
