package model

import "gorm.io/gorm"

func Migration(db *gorm.DB) error {
	if err := db.AutoMigrate(&Terminal{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Transaction{}); err != nil {
		return err
	}

	return nil
}
