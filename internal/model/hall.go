package model

type Hall struct {
	ID     int     `gorm:"primary_key" json:"id"`
	Name   string  `json:"name"`
	Tables []Table `gorm:"foreignkey:HallID" json:"tables"`
}

type HallGoulash struct {
	ID     int            `json:"id"`
	Name   string         `json:"name"`
	Tables []TableGoulash `json:"tables"`
}

type HallApi struct {
	ID     int        `json:"id"`
	Name   string     `json:"name"`
	Tables []TableApi `json:"tables"`
}

type HallResponse struct {
	Success      bool      `json:"success"`
	ErrorMessage string    `json:"error_message"`
	Halls        []HallApi `json:"halls"`
}
