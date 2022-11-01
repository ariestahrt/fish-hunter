package stats

import (
	"fish-hunter/businesses/stats"

	"github.com/gofiber/fiber/v2"
)

type StatsController struct {
	StatsUseCase stats.UseCase
}

func NewStatController(statsUseCase stats.UseCase) *StatsController {
	return &StatsController{
		StatsUseCase: statsUseCase,
	}
}

func (s *StatsController) GetStatistics(c *fiber.Ctx) error {
	stats, err := s.StatsUseCase.GetStatistics()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(stats)
}

func (s *StatsController) GetLastWeekStatistics(c *fiber.Ctx) error {
	stats, err := s.StatsUseCase.GetLastWeekStatistics()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(stats)
}