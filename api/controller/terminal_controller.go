package controller

import (
	"errors"
	"go-gin-project/api/service"
	"go-gin-project/data"
	"go-gin-project/helper"
	"go-gin-project/helper/responsejson"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TerminalController struct {
	terminalService service.TerminalService
}

func NewTerminalController(service service.TerminalService) *TerminalController {
	return &TerminalController{
		terminalService: service,
	}
}

func (c *TerminalController) Create(ctx *gin.Context) {
	createTerminalRequest := data.TerminalRequest{}
	err := ctx.ShouldBindJSON(&createTerminalRequest)
	if err != nil {
		responsejson.BadRequest(ctx, err, "Invalid request body")
		return
	}

	terminalResponse, err := c.terminalService.Create(createTerminalRequest)
	if err != nil {
		if errors.Is(err, helper.ErrFailedValidation) {
			responsejson.BadRequest(ctx, err, "Validation error")
			return
		}
		if errors.Is(err, helper.ErrUniqueViolation) {
			responsejson.Conflict(ctx, err, "Unique constraint violation")
			return
		}
		if errors.Is(err, helper.ErrForeignKeyViolation) {
			responsejson.BadRequest(ctx, err, "Foreign key violation")
			return
		}
		responsejson.InternalServerError(ctx, err, "Failed to create terminal")
		return
	}
	responsejson.Created(ctx, terminalResponse, "Terminal created successfully")
}

func (c *TerminalController) FindAll(ctx *gin.Context) {
	terminalResponse, err := c.terminalService.FindAll()
	if err != nil {
		responsejson.InternalServerError(ctx, err, "Failed to retrieve terminals")
		return
	}
	responsejson.Success(ctx, terminalResponse, "Terminals retrieved successfully")
}

func (c *TerminalController) FindById(ctx *gin.Context) {
	terminalId := ctx.Param("terminalId")

	terminalResponse, err := c.terminalService.FindById(terminalId)
	if err != nil {
		if errors.Is(err, helper.ErrNotFound) {
			responsejson.NotFound(ctx, "Terminal not found")
			return
		}
		responsejson.InternalServerError(ctx, err, "Failed to retrieve terminal")
		return
	}
	responsejson.Success(ctx, terminalResponse, "Terminal retrieved successfully")
}

func (c *TerminalController) Update(ctx *gin.Context) {
	terminalId := ctx.Param("terminalId")
	updateTerminalRequest := data.TerminalRequest{}
	err := ctx.ShouldBindJSON(&updateTerminalRequest)
	if err != nil {
		responsejson.BadRequest(ctx, err, "Invalid request body")
		return
	}

	terminalResponse, err := c.terminalService.Update(terminalId, updateTerminalRequest)
	if err != nil {
		if errors.Is(err, helper.ErrNotFound) {
			responsejson.NotFound(ctx, "Terminal not found")
			return
		}
		if errors.Is(err, helper.ErrFailedValidation) {
			responsejson.BadRequest(ctx, err, "Validation error")
			return
		}
		if errors.Is(err, helper.ErrUniqueViolation) {
			responsejson.Conflict(ctx, err, "Unique constraint violation")
			return
		}
		if errors.Is(err, helper.ErrForeignKeyViolation) {
			responsejson.BadRequest(ctx, err, "Foreign key violation")
			return
		}
		responsejson.InternalServerError(ctx, err, "Failed to update terminal")
		return
	}

	responsejson.Success(ctx, terminalResponse, "Terminal updated successfully")
}

func (c *TerminalController) Delete(ctx *gin.Context) {
	terminalId := ctx.Param("terminalId")
	id, err := strconv.Atoi(terminalId)
	if err != nil {
		responsejson.BadRequest(ctx, err, "Invalid terminal ID")
		return
	}

	if err := c.terminalService.Delete(id); err != nil {
		if errors.Is(err, helper.ErrNotFound) {
			responsejson.NotFound(ctx, "Terminal not found")
			return
		}
		if errors.Is(err, helper.ErrForeignKeyViolation) {
			responsejson.BadRequest(ctx, err, "Cannot delete terminal due to foreign key constraint")
			return
		}
		responsejson.InternalServerError(ctx, err, "Failed to delete terminal")
		return
	}

	responsejson.Success(ctx, nil, "Terminal deleted successfully")
}
