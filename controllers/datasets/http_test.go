package datasets_test

import (
	_datasets "fish-hunter/businesses/datasets"
	_mockDataset "fish-hunter/businesses/datasets/mocks"

	_datasetUtilMock "fish-hunter/util/datasetutil/mocks"
	_s3Mock "fish-hunter/util/s3/mocks"

	_mockUser "fish-hunter/businesses/users/mocks"

	"fish-hunter/businesses/users"
	"fish-hunter/controllers/datasets"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

var (
	mockDatasetRepository _mockDataset.Repository
	mockUserRepository _mockUser.Repository

	datasetController datasets.DatasetController
	datasetUseCase _datasets.UseCase

	userUseCase users.UseCase

	mockS3 _s3Mock.AWS_S3
	mockDatasetUtil _datasetUtilMock.DatasetUtil
)

func TestMain(m *testing.M) {
	mockDatasetRepository = _mockDataset.Repository{}
	datasetUseCase = _datasets.NewDatasetUseCase(&mockDatasetRepository, &mockS3, &mockDatasetUtil)

	userUseCase = users.NewUserUseCase(&mockUserRepository)
	datasetController = *datasets.NewDatasetController(datasetUseCase, userUseCase)

	// m.Run()
}

// Test Status
func TestStatus(t *testing.T) {
	testCases := []struct {
		name string
		endpoint string
		method string
		status string
		expected string
	}{
		{
			name: "Test Case 1 | Valid Status",
			endpoint: "/api/v1/datasets/new",
			method: "GET",
			expected: "OK",
		},
	}

	f := fiber.New()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Create FastHTTP Requests Context
			
			c := &fasthttp.RequestCtx{}
			c.Conn()
			c.Request = *fasthttp.AcquireRequest()
			c.Request.Header.SetMethod(testCase.method)
			c.Request.SetRequestURI(testCase.endpoint)
			
			// Convert fasthttp to fiber context
			fctx := f.AcquireCtx(c)
	
			// Perform the request
			err := datasetController.Status(fctx)
	
			if err != nil {
				t.Errorf("Error => %s", err.Error())
			}
	
			assert.Nil(t, err)	
		})
	}

}