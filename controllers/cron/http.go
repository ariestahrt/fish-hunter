package cron

import (
	"fish-hunter/businesses/cron"

	"github.com/gofiber/fiber/v2"
)

type CronController struct {
	cronUseCase cron.UseCase
}

func NewCronController(cronUseCase cron.UseCase) *CronController {
	return &CronController{
		cronUseCase,
	}
}

func (ctrl *CronController) CleanUpToken(c *fiber.Ctx) error {
	ctrl.cronUseCase.CleanUpToken()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}