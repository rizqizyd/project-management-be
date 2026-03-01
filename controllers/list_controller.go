package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizqizyd/project-management-be/models"
	"github.com/rizqizyd/project-management-be/services"
	"github.com/rizqizyd/project-management-be/utils"
)

type ListController struct {
	service services.ListService
}

func NewListController(s services.ListService) *ListController {
	return &ListController{service: s}
}

func (c *ListController) CreateList(ctx *fiber.Ctx) error {
	list := new(models.List)
	if err := ctx.BodyParser(list); err != nil {
		return utils.BadRequest(ctx, "Failed to parse request body", err.Error())
	}
	if err := c.service.Create(list); err != nil {
		return utils.BadRequest(ctx, "Failed to create list", err.Error())
	}
	return utils.Success(ctx, "List created successfully", list)
}
