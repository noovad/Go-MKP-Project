package model

import "time"

type Transaction struct {
	Id            int       `gorm:"type:int;primary_key"`
	CardId        int       `gorm:"type:int"`
	TerminalIdIn  int       `gorm:"type:int"`
	TerminalIn    Terminal  `gorm:"foreignKey:TerminalIdIn;references:Id"`
	TerminalIdOut int       `gorm:"type:int"`
	TerminalOut   Terminal  `gorm:"foreignKey:TerminalIdOut;references:Id"`
	CheckinTime   time.Time `gorm:"type:timestamp"`
	CheckoutTime  time.Time `gorm:"type:timestamp"`
	Fare          float64   `gorm:"type:decimal(10,2)"`
	BalanceBefore float64   `gorm:"type:decimal(10,2)"`
	BalanceAfter  float64   `gorm:"type:decimal(10,2)"`
	CreatedAt     time.Time `gorm:"type:timestamp"`
}
