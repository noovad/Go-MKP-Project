package model

type Terminal struct {
	Id       int    `gorm:"type:int;primary_key"`
	Name     string `gorm:"type:varchar(100);unique"`
	Location string `gorm:"type:varchar(255)"`
	Status   string `gorm:"type:varchar(50)"`
}
