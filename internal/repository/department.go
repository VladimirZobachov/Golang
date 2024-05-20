package repository

import (
	"fmt"

	"hostess-service/internal/model"
)

func (r *repo) GetDepartmentSettings(departmentId int64) (*model.DepartmentSettings, error) {
	settings := &model.DepartmentSettings{}

	var department model.Department
	if err := r.db.Where("department_id = ?", departmentId).First(&department).Error; err != nil {
		return nil, fmt.Errorf("get department settings: %v", err)
	}
	settings.DepartmentID = department.Id
	settings.DepartmentName = department.Name

	workTimes, err := r.GetAllWorkTimeByDepartment(departmentId, "")
	if err != nil {
		return nil, fmt.Errorf("get all work time by department: %v", err)
	}

	settings.WorkTimes = workTimes

	// никаких созданий новых сущностей внутри репозитория, он у тебя единый
	// всю эту логику совмещения разных сущностей надо уносить в отдельный слой - в бизнес логику.
	// пример накидал
	//authorService := NewAuthorRepo(s.db)
	authors, err := r.GetAllAuthors()
	if err != nil {
		return nil, fmt.Errorf("get authors: %v", err)
	}
	settings.Hostesses = authors

	return settings, nil
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
