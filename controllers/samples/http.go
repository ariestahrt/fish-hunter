package samples

import (
	"fish-hunter/businesses/samples"
	"fish-hunter/controllers/samples/requests"

	"github.com/gofiber/fiber/v2"
)

type SampleController struct {
	SampleUseCase samples.UseCase
}

func NewSampleController(sampleUseCase samples.UseCase) *SampleController {
	return &SampleController{
		SampleUseCase: sampleUseCase,
	}
}

func (u *SampleController) GetAll(c *fiber.Ctx) error {
	samples, err := u.SampleUseCase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(samples)
}

func (u *SampleController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	sample, err := u.SampleUseCase.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(sample)
}

func (u *SampleController) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var sampleUpdate requests.UpdateSamples
	if err := c.BodyParser(&sampleUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := sampleUpdate.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	sample, err := u.SampleUseCase.Update(id, sampleUpdate.ToDomain())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(sample)
}