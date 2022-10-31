package datasets

import (
	"bytes"
	"fish-hunter/businesses/datasets"
	"fish-hunter/controllers/datasets/requests"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type DatasetController struct {
	DatasetUseCase datasets.UseCase
}

func NewDatasetController(datasetUseCase datasets.UseCase) *DatasetController {
	return &DatasetController{
		DatasetUseCase: datasetUseCase,
	}
}

func (u *DatasetController) Status(c *fiber.Ctx) error {
	status := c.Params("status")

	datasets, err := u.DatasetUseCase.Status(status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(datasets)
}

func (u *DatasetController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	dataset, err := u.DatasetUseCase.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dataset)
}

func (u *DatasetController) Activate(c *fiber.Ctx) error {
	id := c.Params("id")

	view_path, err := u.DatasetUseCase.Activate(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Delete for
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Dataset activated",
		"path":	 view_path,
	})
}

func (u *DatasetController) View(c *fiber.Ctx) error {
	id := c.Params("id")
	file_request := c.Params("*")

	// Make sure not path traversal
	if strings.Contains(file_request, "../") || strings.Contains(file_request, "..\\") {
		return c.Status(fiber.StatusForbidden).SendString("Mau ngapain!")
	}

	file_path := "files/datasets/" + id + "/" + file_request

	// Check if file is image or video
	imgExt := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg"}
	videoExt := []string{".mp4", ".avi", ".mkv", ".mov"}

	for _, ext := range append(imgExt, videoExt...) {
		if strings.HasSuffix(file_path, ext) {
			return c.SendFile(file_path)
		}
	}

	// Read File
	file_byte, err := os.ReadFile(file_path)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("File not found")
	}

	// Add content-type
	if strings.HasSuffix(file_path, ".css") {
		c.Set("Content-Type", "text/css")
	} else if strings.HasSuffix(file_path, ".js") {
		c.Set("Content-Type", "text/javascript")
	} else if strings.HasSuffix(file_path, ".json") {
		c.Set("Content-Type", "application/json")
	} else if strings.HasSuffix(file_path, ".xml") {
		c.Set("Content-Type", "application/xml")
	} else if strings.HasSuffix(file_path, ".txt") {
		c.Set("Content-Type", "text/plain")
	} else if strings.HasSuffix(file_path, ".html") {
		c.Set("Content-Type", "text/html")
	}

	return c.Status(fiber.StatusOK).SendStream(bytes.NewReader(file_byte))
}

func (u *DatasetController) Validate(c *fiber.Ctx) error {
	id := c.Params("id")
	userInput := requests.DatasetValidateRequest{}

	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	dataset, err := u.DatasetUseCase.Validate(*userInput.ToDomain(id))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dataset)
}

func (u *DatasetController) TopBrands(c *fiber.Ctx) error {
	topBrands, err := u.DatasetUseCase.TopBrands()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(topBrands)
}

func (u *DatasetController) Download(c *fiber.Ctx) error {
	id := c.Params("id")

	file, err := u.DatasetUseCase.Download(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Get dataset information
	dataset, _ := u.DatasetUseCase.GetByID(id)
	brands_joined := strings.Join(dataset.Brands, "-")

	// Add content disposition
	c.Set("Content-Disposition", "attachment; filename="+brands_joined+".7z")

	return c.Status(fiber.StatusOK).SendFile(file)
}