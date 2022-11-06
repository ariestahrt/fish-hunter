package stats_test

import (
	"fish-hunter/businesses/datasets"
	_mockDataset "fish-hunter/businesses/datasets/mocks"
	"fish-hunter/businesses/jobs"
	_mockJob "fish-hunter/businesses/jobs/mocks"
	"fish-hunter/businesses/stats"
	"fish-hunter/businesses/urls"
	_mockUrl "fish-hunter/businesses/urls/mocks"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	mockJobRepository _mockJob.Repository
	mockUrlRepository _mockUrl.Repository
	mockDatasetRepository _mockDataset.Repository

	statService stats.UseCase

	jobDomain jobs.Domain
	datasetDomain datasets.Domain
	urlDomain urls.Domain

	jobDomainArray []jobs.Domain
	datasetDomainArray []datasets.Domain
	urlDomainArray []urls.Domain
)

func TestMain(m *testing.M) {
	statService = stats.NewStatUseCase(&mockDatasetRepository, &mockJobRepository, &mockUrlRepository)
	
	// Random number betweeen 1 and 10
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10) + 1
	for i := 0; i < randNum; i++ {
		jobDomain = jobs.Domain{
			Id:         primitive.NewObjectID(),
			Ref_Url:    primitive.NewObjectID(),
			Url:        "https://www.google.com",
			HTTPStatus: 200,
			SaveStatus: "success",
			Details:    "success",
			Worker:     "worker",
			CreatedAt:  primitive.NewDateTimeFromTime(time.Now()),
			UpdatedAt:  primitive.NewDateTimeFromTime(time.Now()),
			DeletedAt:  primitive.NewDateTimeFromTime(time.Time{}),
		}
		jobDomainArray = append(jobDomainArray, jobDomain)
	}

	randNum = rand.Intn(10) + 1
	for i := 0; i < randNum; i++ {
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
		datasetDomainArray = append(datasetDomainArray, datasetDomain)
	}

	randNum = rand.Intn(10) + 1
	for i := 0; i < randNum; i++ {
		urlDomain = urls.Domain{
			Id: primitive.NewObjectID(),
			Ref_Source: primitive.NewObjectID(),
			Url: "https://www.phish.com",
			Status: "active",
			Source_Url: "https://openphish.com",
			Source_Name: "OpenPhish",
			CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
			UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
			DeletedAt: primitive.NewDateTimeFromTime(time.Now()),
		}
		urlDomainArray = append(urlDomainArray, urlDomain)
	}

	m.Run()
}

// Test GetStatistics
func TestGetStatistics(t *testing.T) {
	t.Run("Test Case 1 | Get Statistics", func(t *testing.T) {
		mockUrlRepository.On("CountTotal").Return(int64(len(urlDomainArray)), nil).Once()
		mockJobRepository.On("CountTotal").Return(int64(len(jobDomainArray)), nil).Once()
		mockDatasetRepository.On("CountTotal").Return(int64(len(datasetDomainArray)), nil).Once()
		mockDatasetRepository.On("CountTotalValid").Return(int64(len(datasetDomainArray)), nil).Once()
		
		result, err := statService.GetStatistics()

		assert.Nil(t, err)
		assert.Equal(t, int64(len(urlDomainArray)), result["urls"])
		assert.Equal(t, int64(len(jobDomainArray)), result["jobs"])
		assert.Equal(t, int64(len(datasetDomainArray)), result["datasets"])
		assert.Equal(t, int64(len(datasetDomainArray)), result["valid_datasets"])

		mockJobRepository.AssertExpectations(t)
		mockUrlRepository.AssertExpectations(t)
		mockDatasetRepository.AssertExpectations(t)
	})
}

// Test GetLastWeekStatistics
func TestGetLastWeekStatistics(t *testing.T) {
	t.Run("Test Case 1 | Get Last Week Statistics", func(t *testing.T) {
		
		for i := 0; i < 7; i++ {
			mockUrlRepository.On("GetTotalBetweenDates", mock.Anything, mock.Anything).Return(int64(len(urlDomainArray)), nil).Once()
			mockJobRepository.On("GetTotalBetweenDates", mock.Anything, mock.Anything).Return(int64(len(jobDomainArray)), nil).Once()
			mockDatasetRepository.On("GetTotalBetweenDates", mock.Anything, mock.Anything).Return(int64(len(datasetDomainArray)), nil).Once()
		}

		result, err := statService.GetLastWeekStatistics()
		assert.Nil(t, err)

		assert.Equal(t, len(result["date"].([]string)), 7)
		assert.Equal(t, len(result["total_url"].([]int)), 7)
		assert.Equal(t, len(result["total_job"].([]int)), 7)
		assert.Equal(t, len(result["total_dataset"].([]int)), 7)
	})
}