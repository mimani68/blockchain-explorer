package controller

import (
	"app.io/data/model"
	"app.io/service"
	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService *service.TransactionService) TransactionController {
	return TransactionController{
		TransactionService: *transactionService,
	}
}

func (controller *TransactionController) Route(app *fiber.App) {
	app.Get("/tx", controller.ListOfTransactions)
	app.Get("/tx/:hash", controller.GetSingleTransaction)
}

func (controller *TransactionController) ListOfTransactions(c *fiber.Ctx) error {
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "",
	})
}

func (controller *TransactionController) GetSingleTransaction(c *fiber.Ctx) error {
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "",
	})
}
