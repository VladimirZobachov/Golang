package model

type Author struct {
	ID      int64  `gorm:"primary_key;auto_increment" json:"id"`
	Name    string `gorm:"name" json:"name"`
	Surname string `gorm:"surname" json:"surname"`
	Code    int64  `gorm:"code" json:"code"`
}
