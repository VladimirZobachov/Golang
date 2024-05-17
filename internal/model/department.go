package model

import "time"

type Department struct {
	Id           int64  `gorm:"primary_key;AUTO_INCREMENT"`
	Name         string `gorm:"name"`
	DepartmentID int64  `gorm:"department_id"`
}

type DepartmentSettings struct {
	DepartmentID   int64         `json:"department_id"`
	DepartmentName string        `json:"department_name"`
	WorkTimes      []WorkTimeAPI `json:"work_time"`
	Hostesses      []*Author     `json:"hostesses"`
}

type WorkTime struct {
	Id           int64  `gorm:"primary_key;AUTO_INCREMENT"`
	DepartmentId int64  `gorm:"department_id"`
	DayOfWeek    string `gorm:"day_of_week"`
	OpenTime     string `gorm:"open_time"`
	CloseTime    string `gorm:"close_time"`
}

type WorkTimeAPI struct {
	DayOfWeek string `gorm:"day_of_week" json:"day_of_week"`
	OpenTime  string `gorm:"open_time" json:"open_time"`
	CloseTime string `gorm:"close_time " json:"close_time"`
}

type Holiday struct {
	Id           int64     `gorm:"primary_key;AUTO_INCREMENT"`
	DepartmentId int64     `gorm:"department_id"`
	Date         time.Time `gorm:"date"`
	Name         string    `gorm:"name"`
}
