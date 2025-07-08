package repository

import (
	"errors"
	"go-gin-project/helper"
	"go-gin-project/model"

	"gorm.io/gorm"
)

type TerminalRepository interface {
	Save(terminal model.Terminal) (model.Terminal, error)
	FindAll() ([]model.Terminal, error)
	FindById(terminalId string) (terminal model.Terminal, err error)
	Update(terminal model.Terminal) (model.Terminal, error)
	Delete(terminalId int) error
}

func NewTerminalRepositoryImpl(Db *gorm.DB) TerminalRepository {
	return &TerminalRepositoryImpl{Db: Db}
}

type TerminalRepositoryImpl struct {
	Db *gorm.DB
}

func (t *TerminalRepositoryImpl) Save(terminal model.Terminal) (model.Terminal, error) {
	result := t.Db.Create(&terminal)
	if result.Error != nil {
		return model.Terminal{}, result.Error
	}
	return terminal, nil
}

func (t *TerminalRepositoryImpl) FindAll() ([]model.Terminal, error) {
	var terminals []model.Terminal
	result := t.Db.Find(&terminals)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return []model.Terminal{}, nil
		}
		return nil, result.Error
	}
	return terminals, nil
}

func (t *TerminalRepositoryImpl) FindById(terminalId string) (terminalModel model.Terminal, err error) {
	var terminal model.Terminal
	result := t.Db.First(&terminal, terminalId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.Terminal{}, helper.ErrNotFound
		}
		return model.Terminal{}, result.Error
	}
	return terminal, nil
}

func (t *TerminalRepositoryImpl) Update(terminals model.Terminal) (model.Terminal, error) {
	result := t.Db.Model(&terminals).Updates(terminals)
	if result.Error != nil {
		return model.Terminal{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.Terminal{}, helper.ErrNotFound
	}

	var updatedTerminal model.Terminal
	err := t.Db.Where("id = ?", terminals.Id).First(&updatedTerminal).Error
	if err != nil {
		return model.Terminal{}, err
	}

	return updatedTerminal, nil
}

func (t *TerminalRepositoryImpl) Delete(terminalsId int) error {
	deleteResult := t.Db.Delete(&model.Terminal{}, terminalsId)
	if deleteResult.Error != nil {
		if errors.Is(deleteResult.Error, gorm.ErrForeignKeyViolated) {
			return gorm.ErrForeignKeyViolated
		}
		return deleteResult.Error
	}

	if deleteResult.RowsAffected == 0 {
		return helper.ErrNotFound
	}

	return nil
}
