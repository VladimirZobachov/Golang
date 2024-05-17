package service

import (
	"fmt"
	"gorm.io/gorm"
	"hostess-service/internal/model"
)

type DepartmentService interface {
	GetAllWorkTimeByDepartment(departmentId int64, dayOfWeek string) ([]model.WorkTimeAPI, error)
	GetDepartmentSettings(departmentId int64) (*model.DepartmentSettings, error)
}

type departmentService struct {
	db *gorm.DB
}

func (s *departmentService) GetDepartmentSettings(departmentId int64) (*model.DepartmentSettings, error) {
	settings := &model.DepartmentSettings{}

	var department model.Department
	if err := s.db.Where("department_id = ?", departmentId).First(&department).Error; err != nil {
		return nil, fmt.Errorf("get department settings: %v", err)
	}
	settings.DepartmentID = department.Id
	settings.DepartmentName = department.Name

	workTimes, err := s.GetAllWorkTimeByDepartment(departmentId, "")
	if err != nil {
		return nil, fmt.Errorf("get all work time by department: %v", err)
	}

	settings.WorkTimes = workTimes

	authorService := NewAuthorService(s.db)
	authors, err := authorService.GetAllAuthors()
	if err != nil {
		return nil, fmt.Errorf("get authors: %v", err)
	}
	settings.Hostesses = authors

	return settings, nil
}

func NewDepartmentService(db *gorm.DB) DepartmentService {
	return &departmentService{db: db}
}

func (s *departmentService) GetAllWorkTimeByDepartment(departmentId int64, dayOfWeek string) ([]model.WorkTimeAPI, error) {
	var workTimes []model.WorkTime
	query := s.db.Where("department_id = ?", departmentId)
	if dayOfWeek != "" {
		query = query.Where("day_of_week = ?", dayOfWeek)
	}
	result := query.Find(&workTimes)
	workTimesAPI := make([]model.WorkTimeAPI, 0, len(workTimes))
	for _, wt := range workTimes {
		workTime := model.WorkTimeAPI{
			DayOfWeek: wt.DayOfWeek,
			OpenTime:  wt.OpenTime,
			CloseTime: wt.CloseTime,
		}
		workTimesAPI = append(workTimesAPI, workTime)
	}

	if result.Error != nil {
		return nil, result.Error
	}
	return workTimesAPI, nil
}
