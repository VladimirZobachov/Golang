package model

import (
	"time"
)

type Reservation struct {
	ID               int       `gorm:"primary_key" json:"id"`
	CreatedAt        time.Time `gorm:"created_at" json:"-"`
	UpdatedAt        time.Time `gorm:"updated_at" json:"-"`
	ShiftDate        string    `gorm:"column:shift_date" json:"shift_date"`
	DepartmentID     int       `gorm:"column:department_id" json:"department_id"`
	TableID          int64     `gorm:"column:table_id" json:"table_id"`
	TimeStart        int64     `gorm:"column:time_start" json:"time_start"`
	TimeFinish       int64     `gorm:"column:time_finish" json:"time_finish"`
	Status           string    `gorm:"column:status" json:"status"`
	GuestName        string    `gorm:"column:guest_name" json:"guest_name"`
	GuestEmail       string    `gorm:"column:guest_email" json:"guest_email"`
	GuestTel         string    `gorm:"column:guest_tel" json:"guest_tel"`
	GuestComment     string    `gorm:"column:guest_comment" json:"guest_comment"`
	GuestCount       int64     `gorm:"column:guest_count" json:"guest_count"`
	ReserveComment   string    `gorm:"column:reserve_comment" json:"reserve_comment"`
	Author           string    `gorm:"column:author" json:"author"`
	Source           string    `gorm:"column:source" json:"source"`
	Tags             string    `gorm:"column:tags" json:"tags"`
	ConfirmationType string    `gorm:"column:confirmation_type" json:"confirmation_type"`
}

type ReservationResponse struct {
	Success              bool          `json:"success"`
	ErrorMessage         string        `json:"error_message"`
	ConflictReservations []Reservation `json:"conflict_reservations"`
	Reservation          *Reservation  `json:"reservation"`
}

type ReservationResponseAll struct {
	Success      bool          `json:"success"`
	ErrorMessage string        `json:"error_message"`
	Reservations []Reservation `json:"reservations"`
}
