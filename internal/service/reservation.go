package service

import (
	"errors"
	"gorm.io/gorm"
	"hostess-service/internal/model"
)

type ReservationService interface {
	CreateReservation(reservation *model.Reservation) error
	GetReservationByID(id int64) (*model.Reservation, error)
	GetAllReservations(dateFrom int64, dateTo int64) (model.ReservationResponseAll, error)
	UpdateReservation(id int64, reservation *model.Reservation) error
	UpdateReservationStatus(status string, id int64) error
	DeleteReservationByID(id int64) error
	IsAvailableTimeSlot(tableID int64, timeStart int64, timeFinish int64, reservationIDs ...int64) ([]model.Reservation, error)
}

type reservationService struct {
	db *gorm.DB
}

func NewReservationService(db *gorm.DB) ReservationService {
	return &reservationService{db: db}
}

func (s *reservationService) IsAvailableTimeSlot(tableID int64, timeStart int64, timeFinish int64, reservationIDs ...int64) ([]model.Reservation, error) {
	var reservations []model.Reservation
	err := s.db.Model(&model.Reservation{}).
		Where("table_id = ?", tableID).
		Where("time_finish > ?", timeStart).
		Where("time_start < ?", timeFinish).
		Find(&reservations).Error

	return reservations, err
}

func (s *reservationService) CreateReservation(reservation *model.Reservation) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var count int64
		err := s.db.Model(&model.Reservation{}).
			Where("table_id = ?", reservation.TableID).
			Where("time_finish > ?", reservation.TimeStart).
			Where("time_start < ?", reservation.TimeFinish).
			Count(&count).Error
		if err != nil {
			return err
		}

		if count > 0 {
			return errors.New("reservation already exists")
		}

		return tx.Create(reservation).Error
	})
}

func (s *reservationService) GetReservationByID(id int64) (*model.Reservation, error) {
	var reservation model.Reservation
	result := s.db.First(&reservation, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &reservation, nil
}

func (s *reservationService) GetAllReservations(dateFrom int64, dateTo int64) (model.ReservationResponseAll, error) {
	var reservations []model.Reservation
	query := s.db.Model(&model.Reservation{})

	// Apply filters based on provided dates
	if dateFrom != 0 {
		query = query.Where("time_start >= ?", dateFrom)
	}
	if dateTo != 0 {
		query = query.Where("time_finish <= ?", dateTo)
	}

	// Execute the query
	result := query.Find(&reservations)

	// Check for errors and prepare the response
	if result.Error != nil {
		return model.ReservationResponseAll{
			Success:      false,
			ErrorMessage: result.Error.Error(),
			Reservations: nil,
		}, nil // Return nil as error here since we are handling errors via the API response structure
	}

	// Return success response
	return model.ReservationResponseAll{
		Success:      true,
		ErrorMessage: "",
		Reservations: reservations,
	}, nil
}

func (s *reservationService) UpdateReservation(id int64, reservation *model.Reservation) error {
	result := s.db.Model(reservation).Where("id= ?", id).Updates(reservation)
	return result.Error
}

func (s *reservationService) UpdateReservationStatus(status string, id int64) error {
	result := s.db.Model(&model.Reservation{}).Where("id = ?", id).Update("status", status)
	return result.Error
}

func (s *reservationService) DeleteReservationByID(id int64) error {
	result := s.db.Delete(&model.Reservation{}, id)
	return result.Error
}
