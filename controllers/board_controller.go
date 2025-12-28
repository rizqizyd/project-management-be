package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/rizqizyd/project-management-be/models"
	"github.com/rizqizyd/project-management-be/services"
	"github.com/rizqizyd/project-management-be/utils"
)

type BoardController struct {
	service services.BoardService
}

func NewBoardController(s services.BoardService) *BoardController {
	return &BoardController{service: s}
}

func (c *BoardController) CreateBoard(ctx *fiber.Ctx) error {
	var userID uuid.UUID
	var err error

	board := new(models.Board)
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if err := ctx.BodyParser(board); err != nil {
		return utils.BadRequest(ctx, "Invalid request body", err.Error())
	}

	userID, err = uuid.Parse(claims["public_id"].(string))
	if err != nil {
		return utils.BadRequest(ctx, "Invalid request body", err.Error())
	}
	board.OwnerPublicID = userID

	if err := c.service.Create(board); err != nil {
		return utils.BadRequest(ctx, "Board creation failed", err.Error())
	}

	return utils.Success(ctx, "Board created successfully", board)
}
