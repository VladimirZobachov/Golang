package repository

import (
	"fmt"

	"hostess-service/internal/model"
)

func (r *repo) GetDepartment(departmentId int64) (*model.Department, error) {
	var department model.Department
	if err := r.db.Where("department_id = ?", departmentId).First(&department).Error; err != nil {
		return nil, fmt.Errorf("get department settings: %v", err)
	}

	return &department, nil
}

func (r *repo) GetAllWorkTimeByDepartment(departmentId int64, dayOfWeek string) ([]model.WorkTimeAPI, error) {
	var workTimes []model.WorkTime
	query := r.db.Where("department_id = ?", departmentId)
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
