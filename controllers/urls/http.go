package urls

import (
	"fish-hunter/businesses/urls"

	"github.com/gofiber/fiber/v2"
)

type UrlController struct {
	UrlUseCase urls.UseCase
}

func NewUrlController(urlUseCase urls.UseCase) *UrlController {
	return &UrlController{
		UrlUseCase: urlUseCase,
	}
}

func (u *UrlController) GetAll(c *fiber.Ctx) error {
	urls, err := u.UrlUseCase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(urls)
}

func (u *UrlController) FetchUrl(c *fiber.Ctx) error {
	source := c.Params("source")

	urls, err := u.UrlUseCase.FetchUrl(source)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(urls)
}

func (u *UrlController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	url, err := u.UrlUseCase.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(url)
}