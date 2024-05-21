package department

import (
	"fmt"

	"hostess-service/internal/model"
)

type service struct {
	repo departmentRepo
}

func New(r departmentRepo) *service {
	return &service{repo: r}
}

type departmentRepo interface {
	GetAllWorkTimeByDepartment(departmentId int64, dayOfWeek string) ([]model.WorkTimeAPI, error)
	GetAllAuthors() ([]*model.Author, error)
	GetDepartment(departmentId int64) (*model.Department, error)
}

func (s *service) GetDepartmentSettings(departmentId int64) (*model.DepartmentSettings, error) {
	department, _ := s.repo.GetDepartment(departmentId)

	workTimes, err := s.repo.GetAllWorkTimeByDepartment(departmentId, "")
	if err != nil {
		return nil, fmt.Errorf("get all work time by department: %v", err)
	}

	authors, err := s.repo.GetAllAuthors()
	if err != nil {
		return nil, fmt.Errorf("get authors: %v", err)
	}

	return &model.DepartmentSettings{
		DepartmentID:   departmentId,
		DepartmentName: department.Name,
		WorkTimes:      workTimes,
		Hostesses:      authors,
	}, nil
}

func (s *service) GetAllWorkTimeByDepartment(departmentId int64, dayOfWeek string) ([]model.WorkTimeAPI, error) {
	return s.repo.GetAllWorkTimeByDepartment(departmentId, dayOfWeek)
}
