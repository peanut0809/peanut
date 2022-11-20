// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"golang8090/internal/dao/internal"
)

// internalEventeECardAvoidAppointmentDao is internal type for wrapping internal DAO implements.
type internalEventeECardAvoidAppointmentDao = *internal.EventeECardAvoidAppointmentDao

// eventeECardAvoidAppointmentDao is the data access object for table evente_e_card_avoid_appointment.
// You can define custom methods on it to extend its functionality as you wish.
type eventeECardAvoidAppointmentDao struct {
	internalEventeECardAvoidAppointmentDao
}

var (
	// EventeECardAvoidAppointment is globally public accessible object for table evente_e_card_avoid_appointment operations.
	EventeECardAvoidAppointment = eventeECardAvoidAppointmentDao{
		internal.NewEventeECardAvoidAppointmentDao(),
	}
)

// Fill with you ideas below.