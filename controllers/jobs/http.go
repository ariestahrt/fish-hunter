package jobs

import (
	"fish-hunter/businesses/jobs"

	"github.com/gofiber/fiber/v2"
)

type JobController struct {
	JobUseCase jobs.UseCase
}

func NewJobController(jobUseCase jobs.UseCase) *JobController {
	return &JobController{
		JobUseCase: jobUseCase,
	}
}

func (u *JobController) GetAll(c *fiber.Ctx) error {
	jobs, err := u.JobUseCase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(jobs)
}

func (u *JobController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	job, err := u.JobUseCase.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(job)
}