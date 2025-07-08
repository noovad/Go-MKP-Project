package service

import (
	"go-gin-project/api/repository"
	"go-gin-project/data"
	"go-gin-project/helper"
	"go-gin-project/model"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type TerminalService interface {
	Create(terminal data.TerminalRequest) (data.TerminalResponse, error)
	FindAll() ([]data.TerminalResponse, error)
	FindById(terminalId uuid.UUID) (data.TerminalResponse, error)
	Update(terminalId uuid.UUID, terminal data.TerminalRequest) (data.TerminalResponse, error)
	Delete(terminalId uuid.UUID) error
}

func NewTerminalServiceImpl(terminalRepository repository.TerminalRepository, validate *validator.Validate) TerminalService {
	return &TerminalServiceImpl{
		TerminalRepository: terminalRepository,
		Validate:           validate,
	}
}

type TerminalServiceImpl struct {
	TerminalRepository repository.TerminalRepository
	Validate           *validator.Validate
}

func (t *TerminalServiceImpl) Create(terminal data.TerminalRequest) (data.TerminalResponse, error) {
	err := t.Validate.Struct(terminal)
	if err != nil {
		return data.TerminalResponse{}, helper.WrapValidation(err)
	}

	terminalModel := model.Terminal{
		Name:     terminal.Name,
		Location: terminal.Location,
		Status:   terminal.Status,
	}

	createdTerminal, err := t.TerminalRepository.Save(terminalModel)
	if err != nil {
		return data.TerminalResponse{}, helper.HandlePostgresError(err)
	}

	terminalResponse := data.TerminalResponse{
		Id:       createdTerminal.Id,
		Name:     createdTerminal.Name,
		Location: createdTerminal.Location,
		Status:   createdTerminal.Status,
	}

	return terminalResponse, nil
}

func (t *TerminalServiceImpl) FindAll() ([]data.TerminalResponse, error) {
	result, err := t.TerminalRepository.FindAll()
	if err != nil {
		return nil, helper.HandlePostgresError(err)
	}

	var terminals []data.TerminalResponse
	for _, value := range result {
		terminal := data.TerminalResponse{
			Id:       value.Id,
			Name:     value.Name,
			Location: value.Location,
			Status:   value.Status,
		}
		terminals = append(terminals, terminal)
	}

	return terminals, nil
}

func (t *TerminalServiceImpl) FindById(terminalId uuid.UUID) (data.TerminalResponse, error) {
	terminalData, err := t.TerminalRepository.FindById(terminalId)
	if err != nil {
		return data.TerminalResponse{}, helper.HandlePostgresError(err)
	}

	terminalResponse := data.TerminalResponse{
		Id:       terminalData.Id,
		Name:     terminalData.Name,
		Location: terminalData.Location,
		Status:   terminalData.Status,
	}

	return terminalResponse, nil
}

func (t *TerminalServiceImpl) Update(terminalId uuid.UUID, terminal data.TerminalRequest) (data.TerminalResponse, error) {
	err := t.Validate.Struct(terminal)
	if err != nil {
		return data.TerminalResponse{}, helper.WrapValidation(err)
	}

	terminalData, err := t.TerminalRepository.FindById(terminalId)
	if err != nil {
		return data.TerminalResponse{}, helper.HandlePostgresError(err)
	}

	if terminalData.Id == uuid.Nil {
		return data.TerminalResponse{}, helper.ErrNotFound
	}

	terminalData.Name = terminal.Name
	terminalData.Location = terminal.Location
	terminalData.Status = terminal.Status

	updatedTerminal, err := t.TerminalRepository.Update(terminalData)
	if err != nil {
		return data.TerminalResponse{}, helper.HandlePostgresError(err)
	}

	terminalResponse := data.TerminalResponse{
		Id:       updatedTerminal.Id,
		Name:     updatedTerminal.Name,
		Location: updatedTerminal.Location,
		Status:   updatedTerminal.Status,
	}

	return terminalResponse, nil
}

func (t *TerminalServiceImpl) Delete(terminalId uuid.UUID) error {
	err := t.TerminalRepository.Delete(terminalId)
	if err != nil {
		return helper.HandlePostgresError(err)
	}
	return nil
}
