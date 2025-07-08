package model

import "github.com/google/uuid"

type Terminal struct {
	Id       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name     string    `gorm:"type:varchar(100);unique;not null"`
	Location string    `gorm:"type:varchar(255);not null"`
	Status   string    `gorm:"type:varchar(50);not null"`
}
