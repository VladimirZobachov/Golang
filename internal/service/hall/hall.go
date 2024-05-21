package hall

import "hostess-service/internal/model"

type service struct {
	repo  hallRepo
	proxy proxy
}

func New(r hallRepo, p proxy) *service {
	return &service{repo: r, proxy: p}
}

type hallRepo interface {
	GetAllWorkTimeByDepartment(departmentId int64, dayOfWeek string) ([]model.WorkTimeAPI, error)
	GetAllAuthors() ([]*model.Author, error)
	GetDepartment(departmentId int64) (*model.Department, error)
}

type proxy interface {
	ImportHallsFromGoulash() model.HallResponse
}
