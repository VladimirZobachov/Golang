package model

type Table struct {
	ID      int    `gorm:"primary_key"`
	Name    string `gorm:"name"`
	Persons int    `gorm:"persons"`
	PlaceX  int    `gorm:"column:place_x"`
	PlaceY  int    `gorm:"column:place_y"`
	Width   int    `gorm:"column:width"`
	Height  int    `gorm:"column:height"`
	Type    string `gorm:"column:type"`
	HallID  int    `gorm:"column:hall_id"`
}

type TableGoulash struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TableApi struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Persons int    `json:"persons"`
	Place   Place  `json:"place"`
	Size    Size   `json:"size"`
	Type    string `json:"type"`
}

type Place struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
