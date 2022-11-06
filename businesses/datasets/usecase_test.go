package datasets_test

import (
	"errors"
	"testing"
	"time"

	"fish-hunter/businesses/datasets"
	"fish-hunter/businesses/datasets/mocks"
	"fish-hunter/util"

	_datasetUtilMock "fish-hunter/util/datasetutil/mocks"
	_s3Mock "fish-hunter/util/s3/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	datasetsRepository mocks.Repository
	s3Mock			 	_s3Mock.AWS_S3
	datasetUtilMock 	_datasetUtilMock.DatasetUtil
	datasetService     datasets.UseCase
	datasetDomain      datasets.Domain
)

func TestMain(m *testing.M) {
	datasetService = datasets.NewDatasetUseCase(&datasetsRepository, &s3Mock, &datasetUtilMock)
	datasetDomain = datasets.Domain{
		Id: primitive.NewObjectID(),
		Ref_Url: primitive.NewObjectID(),
		Ref_Job: primitive.NewObjectID(),
		DateScrapped: primitive.NewDateTimeFromTime(time.Now()),
		HttpStatus: 200,
		Domain: "www.google.com",
		AssetsDownloaded: 0,
		ContentLength: 0,
		Url: "https://www.google.com",
		Categories: []string{"category1", "category2"},
		Brands: []string{"brand1", "brand2"},
		DatasetPath: "dataset_path",
		HtmldomPath: "htmldom_path",
		ScrappedFrom: "scrapped_from",
		UrlscanUuid: "urlscan_uuid",
		Status: "new",
		ScreenshotPath: "screenshot_path",
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
		DeletedAt: primitive.NewDateTimeFromTime(time.Time{}), // null time
	}
	m.Run()
}

// TestGetByID ...
func TestGetByID(t *testing.T) {
	t.Run("Test Case 1 | Valid GetByID", func(t *testing.T) {
		datasetsRepository.On("GetByID", datasetDomain.Id.Hex()).Return(datasetDomain, nil).Once()

		result, err := datasetService.GetByID(datasetDomain.Id.Hex())
		if err != nil {
			t.Errorf("this should not be error, but got %v", err)
		}

		assert.Equal(t, datasetDomain, result)
	})

	t.Run("Test Case 2 | Invalid GetByID", func(t *testing.T) {
		datasetsRepository.On("GetByID", datasetDomain.Id.Hex()).Return(datasets.Domain{}, errors.New("")).Once()

		result, err := datasetService.GetByID(datasetDomain.Id.Hex())
		if err == nil {
			t.Errorf("this should be error, but got %v", err)
		}

		assert.Equal(t, datasets.Domain{}, result)
	})
}

// Get Status
func TestGetStatus(t *testing.T) {
	t.Run("Test Case 1 | Valid GetStatus", func(t *testing.T) {
		datasetsRepository.On("Status", datasetDomain.Status).Return([]datasets.Domain{datasetDomain}, nil).Once()

		result, err := datasetService.Status(datasetDomain.Status)
		if err != nil {
			t.Errorf("this should not be error, but got %v", err)
		}

		assert.Equal(t, []datasets.Domain{datasetDomain}, result)
	})
}

// Test Validate
func TestValidate(t *testing.T) {
	t.Run("Test Case 1 | Valid Validate", func(t *testing.T) {
		datasetsRepository.On("Validate", datasetDomain).Return(datasetDomain, nil).Once()

		result, err := datasetService.Validate(datasetDomain)
		if err != nil {
			t.Errorf("this should not be error, but got %v", err)
		}

		assert.Equal(t, datasetDomain, result)
	})
}

// Test Top Brand
func TestTopBrands(t *testing.T) {
	t.Run("Test Case 1 | Valid TopBrands", func(t *testing.T) {
		ret := map[string]interface{}{
			"brand1": 1,
			"brand2": 1,
		}
		datasetsRepository.On("TopBrands").Return(ret, nil).Once()

		top_brands, err := datasetService.TopBrands()
		if err != nil {
			t.Errorf("this should not be error, but got %v", err)
		}

		assert.Equal(t, ret, top_brands)
	})
}

// Test Download
func TestDownload(t *testing.T) {
	t.Run("Test Case 1 | Valid Download", func(t *testing.T) {
		// Add Mock On GetByID
		datasetsRepository.On("GetByID", datasetDomain.Id.Hex()).Return(datasetDomain, nil).Once()
		
		file7z := util.GetConfig("APP_PATH") + "files/"+datasetDomain.Ref_Url.Hex()+".7z"
		
		// Mock On Download
		datasetsRepository.On("Download", datasetDomain.Id.Hex()).Return(file7z, nil).Once()

		// Mock on S3 Download
		s3Mock.On("DownloadFile", mock.Anything, datasetDomain.DatasetPath+".7z").Return(nil).Once()

		// Mock on Extract
		datasetUtilMock.On("Extract7Zip", file7z, util.GetConfig("7Z_PASSWORD")).Return(nil).Once()

		// MOck on Compress
		folder_to_compress := util.GetConfig("APP_PATH") + "files/datasets/" + datasetDomain.Ref_Url.Hex()
		datasetUtilMock.On("Compress7Zip", folder_to_compress).Return(nil).Once()

		result, err := datasetService.Download(datasetDomain.Id.Hex())
		if err != nil {
			t.Errorf("this should not be error, but got %v", err)
		}

		assert.Equal(t, folder_to_compress + ".7z", result)
	})

	t.Run("Test Case 2 | Failed to get by id", func(t *testing.T) {
		// Add Mock On GetByID
		datasetsRepository.On("GetByID", datasetDomain.Id.Hex()).Return(datasets.Domain{}, errors.New("")).Once()

		result, err := datasetService.Download(datasetDomain.Id.Hex())
		assert.Equal(t, "", result)
		assert.NotNil(t, err)
	})

	t.Run("Test Case 3 | Failed to unzip", func(t *testing.T) {
		// Add Mock On GetByID
		datasetsRepository.On("GetByID", datasetDomain.Id.Hex()).Return(datasetDomain, nil).Once()
		
		file7z := util.GetConfig("APP_PATH") + "files/"+datasetDomain.Ref_Url.Hex()+".7z"
		
		// Mock On Download
		datasetsRepository.On("Download", datasetDomain.Id.Hex()).Return(file7z, nil).Once()

		// Mock on S3 Download
		s3Mock.On("DownloadFile", mock.Anything, datasetDomain.DatasetPath+".7z").Return(nil).Once()

		// Mock on Extract
		datasetUtilMock.On("Extract7Zip", file7z, util.GetConfig("7Z_PASSWORD")).Return(errors.New("")).Once()

		result, err := datasetService.Download(datasetDomain.Id.Hex())
		assert.Equal(t, "", result)
		assert.NotNil(t, err)
	})

	t.Run("Test Case 4 | Failed to compress", func(t *testing.T) {
		// Add Mock On GetByID
		datasetsRepository.On("GetByID", datasetDomain.Id.Hex()).Return(datasetDomain, nil).Once()
		
		file7z := util.GetConfig("APP_PATH") + "files/"+datasetDomain.Ref_Url.Hex()+".7z"
		
		// Mock On Download
		datasetsRepository.On("Download", datasetDomain.Id.Hex()).Return(file7z, nil).Once()

		// Mock on S3 Download
		s3Mock.On("DownloadFile", mock.Anything, datasetDomain.DatasetPath+".7z").Return(nil).Once()

		// Mock on Extract
		datasetUtilMock.On("Extract7Zip", file7z, mock.Anything).Return(nil).Once()

		// MOck on Compress
		folder_to_compress := util.GetConfig("APP_PATH") + "files/datasets/" + datasetDomain.Ref_Url.Hex()
		datasetUtilMock.On("Compress7Zip", folder_to_compress).Return(errors.New("")).Once()

		result, err := datasetService.Download(datasetDomain.Id.Hex())
		assert.Equal(t, "", result)
		assert.NotNil(t, err)
	})

	t.Run("Test Case 5 | Failed to download from s3", func(t *testing.T) {
		// Add Mock On GetByID
		datasetsRepository.On("GetByID", datasetDomain.Id.Hex()).Return(datasetDomain, nil).Once()
		
		file7z := util.GetConfig("APP_PATH") + "files/"+datasetDomain.Ref_Url.Hex()+".7z"
		
		// Mock On Download
		datasetsRepository.On("Download", datasetDomain.Id.Hex()).Return(file7z, nil).Once()

		// Mock on S3 Download
		s3Mock.On("DownloadFile", mock.Anything, datasetDomain.DatasetPath+".7z").Return(errors.New("")).Once()

		result, err := datasetService.Download(datasetDomain.Id.Hex())
		assert.Equal(t, "", result)
		assert.NotNil(t, err)
	})
}

// Test Activate
func TestActivate(t *testing.T) {
	t.Run("Test Case 1 | Valid Activate", func(t *testing.T) {
		// Add Mock On GetByID
		datasetsRepository.On("GetByID", datasetDomain.Id.Hex()).Return(datasetDomain, nil).Once()
		
		// Mock On Activate
		datasetsRepository.On("Activate", datasetDomain.Id.Hex()).Return(nil).Once()

		// Mock On DownloadFile s3
		s3Mock.On("DownloadFile", mock.Anything, datasetDomain.DatasetPath+".7z").Return(nil).Once()

		file7z := util.GetConfig("APP_PATH") + "files/" + datasetDomain.Ref_Url.Hex() + ".7z"
		// Mock On Extract
		datasetUtilMock.On("Extract7Zip", file7z, util.GetConfig("7Z_PASSWORD")).Return(nil).Once()
		
		// Mock On TimedPruneDirectory
		datasetUtilMock.On("TimedPruneDirectory", "files/"+ datasetDomain.DatasetPath, mock.Anything).Return(nil).Once()

		expected := "/datasets/view/" + datasetDomain.Ref_Url.Hex() + "/index.html"
		result, err := datasetService.Activate(datasetDomain.Id.Hex())
		assert.Nil(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Test Case 2 | Failed to get by id", func(t *testing.T) {
		// Add Mock On GetByID
		datasetsRepository.On("GetByID", datasetDomain.Id.Hex()).Return(datasets.Domain{}, errors.New("")).Once()

		result, err := datasetService.Activate(datasetDomain.Id.Hex())
		assert.Equal(t, "", result)
		assert.NotNil(t, err)
	})

	t.Run("Test Case 3 | Failed Download From S3", func(t *testing.T) {
		// Add Mock On GetByID
		datasetsRepository.On("GetByID", datasetDomain.Id.Hex()).Return(datasetDomain, nil).Once()
		
		// Mock On Activate
		datasetsRepository.On("Activate", datasetDomain.Id.Hex()).Return(nil).Once()

		// Mock On DownloadFile s3
		s3Mock.On("DownloadFile", mock.Anything, datasetDomain.DatasetPath+".7z").Return(errors.New("")).Once()

		result, err := datasetService.Activate(datasetDomain.Id.Hex())
		assert.Equal(t, "", result)
		assert.NotNil(t, err)
	})
}