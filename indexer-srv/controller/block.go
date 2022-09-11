package controller

import (
	"strconv"

	"app.io/data/model"
	"app.io/service"
	"github.com/gofiber/fiber/v2"
)

type BlockController struct {
	BlockService service.BlockService
}

func NewBlockController(blockService *service.BlockService) BlockController {
	return BlockController{
		BlockService: *blockService,
	}
}

func (controller *BlockController) Route(app *fiber.App) {
	app.Get("/block", controller.LastBlock)
	app.Get("/block/:number", controller.GetSingleBlock)
}

func (controller *BlockController) LastBlock(c *fiber.Ctx) error {
	tx, err := controller.BlockService.LastBlock()
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

func (controller *BlockController) GetSingleBlock(c *fiber.Ctx) error {
	blockNumber, err := strconv.Atoi(c.Params("number"))
	if err != nil {
		if err != nil {
			return c.JSON(model.WebResponse{
				Code:   400,
				Status: "NOK",
			})
		}
	}
	tx, err := controller.BlockService.GetBlockByNumber(blockNumber)
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
