package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CardId          int       `gorm:"type:int;not null"`
	TerminalIdIn    uuid.UUID `gorm:"type:uuid"`
	TerminalIn      Terminal  `gorm:"foreignKey:TerminalIdIn;references:Id"`
	TerminalIdOut   uuid.UUID `gorm:"type:uuid"`
	TerminalOut     Terminal  `gorm:"foreignKey:TerminalIdOut;references:Id"`
	CheckinTime     time.Time `gorm:"type:timestamp;not null"`
	CheckoutTime    time.Time `gorm:"type:timestamp"`
	Fare            float64   `gorm:"type:decimal(10,2)"`
	BalanceBefore   float64   `gorm:"type:decimal(10,2)"`
	BalanceAfter    float64   `gorm:"type:decimal(10,2)"`
	CreatedAt       time.Time `gorm:"type:timestamp"`
}
