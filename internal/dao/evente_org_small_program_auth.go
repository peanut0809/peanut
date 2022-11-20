// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"golang8090/internal/dao/internal"
)

// internalEventeOrgSmallProgramAuthDao is internal type for wrapping internal DAO implements.
type internalEventeOrgSmallProgramAuthDao = *internal.EventeOrgSmallProgramAuthDao

// eventeOrgSmallProgramAuthDao is the data access object for table evente_org_small_program_auth.
// You can define custom methods on it to extend its functionality as you wish.
type eventeOrgSmallProgramAuthDao struct {
	internalEventeOrgSmallProgramAuthDao
}

var (
	// EventeOrgSmallProgramAuth is globally public accessible object for table evente_org_small_program_auth operations.
	EventeOrgSmallProgramAuth = eventeOrgSmallProgramAuthDao{
		internal.NewEventeOrgSmallProgramAuthDao(),
	}
)

// Fill with you ideas below.