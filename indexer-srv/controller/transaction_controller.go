package controller

import (
	"app.io/data/model"
	"app.io/pkg/exception"
	"app.io/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TransactionController struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService *service.TransactionService) TransactionController {
	return TransactionController{TransactionService: *transactionService}
}

func (controller *TransactionController) Route(app *fiber.App) {
	app.Post("/api/transactions", controller.Create)
	app.Get("/api/transactions", controller.List)
}

func (controller *TransactionController) Create(c *fiber.Ctx) error {
	var request model.CreateTransactionRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.TransactionService.Create(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *TransactionController) List(c *fiber.Ctx) error {
	responses := controller.TransactionService.List()
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}
