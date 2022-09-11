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
	app.Get("/tx", controller.LastTransaction)
	app.Get("/tx/:hash", controller.GetSingleTransaction)
}

func (controller *TransactionController) LastTransaction(c *fiber.Ctx) error {
	tx, err := controller.TransactionService.LastTransaction()
	if err != nil {
		return c.JSON(model.WebResponse{
			Code:   400,
			Status: "NOK",
		})
	}
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   tx,
	})
}

func (controller *TransactionController) GetSingleTransaction(c *fiber.Ctx) error {
	tx, err := controller.TransactionService.GettransactionByHash(c.Params("hash"))
	if err != nil {
		return c.JSON(model.WebResponse{
			Code:   400,
			Status: "NOK",
		})
	}
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   tx,
	})
}
