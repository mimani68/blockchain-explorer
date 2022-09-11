package controller

import (
	"strconv"

	"app.io/data/model"
	"app.io/service"
	"github.com/gofiber/fiber/v2"
)

type StatsController struct {
	StatsService service.StatsService
}

func NewStatsController(statsService *service.StatsService) StatsController {
	return StatsController{
		StatsService: *statsService,
	}
}

func (controller *StatsController) Route(app *fiber.App) {
	app.Get("/stats", controller.Total)
	app.Get("/stats/:start-:end", controller.Custom)
}

func (controller *StatsController) Total(c *fiber.Ctx) error {
	tx, err := controller.StatsService.TotalStatsReport()
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

func (controller *StatsController) Custom(c *fiber.Ctx) error {
	start, err := strconv.Atoi(c.Params("start"))
	if err != nil {
		if err != nil {
			return c.JSON(model.WebResponse{
				Code:   400,
				Status: "NOK",
			})
		}
	}
	end, err := strconv.Atoi(c.Params("end"))
	if err != nil {
		if err != nil {
			return c.JSON(model.WebResponse{
				Code:   400,
				Status: "NOK",
			})
		}
	}
	tx, err := controller.StatsService.CustomStatsReport(start, end)
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
