package urls_test

import (
	"errors"
	"fish-hunter/businesses/urls"
	"fish-hunter/businesses/urls/mocks"
	sourceHelper "fish-hunter/helpers/source"
	"testing"
	"time"

	_scrapper "fish-hunter/util/scrapper"
	_scrapperMock "fish-hunter/util/scrapper/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	urlRepository mocks.Repository
	urlService	urls.UseCase
	urlDomain urls.Domain
	mockUrlScrapper _scrapperMock.Scrapper
	urlScrapper _scrapper.Scrapper
)

/*
type Domain struct {
    Id       		primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Ref_Source		primitive.ObjectID `json:"-"`
    Url 			string             `json:"url,omitempty" validate:"required"`
    Status  string			 `json:"status,omitempty"`
    Source_Url 		string             `json:"source_url,omitempty" validate:"required" bson:"source_url,omitempty"` 
    Source_Name 	string             `json:"source_name,omitempty" validate:"required" bson:"source_name,omitempty"` 
    CreatedAt		primitive.DateTime `json:"created_at,omitempty"`
    UpdatedAt		primitive.DateTime `json:"updated_at,omitempty"`
	DeletedAt		primitive.DateTime `json:"deleted_at,omitempty"`
}
*/

func TestMain(m *testing.M){
	urlScrapper = _scrapper.NewUrlScrapper()
	urlService = urls.NewUrlUseCase(&urlRepository, &mockUrlScrapper)
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
	m.Run()
}

// Test Get ALl
func TestGetAll(t *testing.T){
	t.Run("Test Case 1 | Get All", func(t *testing.T){
		urlRepository.On("GetAll").Return([]urls.Domain{urlDomain}, nil).Once()
		result, err := urlService.GetAll()

		assert.Equal(t, []urls.Domain{urlDomain}, result)
		assert.Nil(t, err)
	})
}

// Test Get By ID
func TestGetByID(t *testing.T){
	t.Run("Test Case 1 | Get By ID", func(t *testing.T){
		urlRepository.On("GetByID", urlDomain.Id.Hex()).Return(urlDomain, nil).Once()
		result, err := urlService.GetByID(urlDomain.Id.Hex())

		assert.Equal(t, urlDomain, result)
		assert.Nil(t, err)
	})
}

// Test Fetch Url
func TestFetchUrl(t *testing.T){
	t.Run("Test Case 1 | OK Fetch Url", func(t *testing.T){
		url_list, _ := urlScrapper.GetPhishUrl("openphish")

		// Mock Get Phish Url
		mockUrlScrapper.On("GetPhishUrl", "openphish").Return(url_list, nil).Once()

		src := sourceHelper.GetSourceInformation("openphish")

		urlDomainArr := []urls.Domain{}

		// Mock Save every url
		for _, _url := range url_list {
			mock_result := urls.Domain{
				Id: primitive.NewObjectID(),
				Ref_Source: src.Id,
				Url: _url,
				Status: "queued",
				Source_Url: "https://openphish.com",
				Source_Name: "OpenPhish",
				CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
				UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
				DeletedAt: primitive.NewDateTimeFromTime(time.Time{}),
			}

			urlRepository.On("Save", mock.Anything).Return(mock_result, nil).Once()

			// Append to array
			urlDomainArr = append(urlDomainArr, mock_result)
		}

		result, err := urlService.FetchUrl("openphish")

		assert.Equal(t, urlDomainArr, result)
		assert.Nil(t, err)
	})

	// Scrapper error
	t.Run("Test Case 2 | Scrapper Error", func(t *testing.T){
		// Mock Get Phish Url
		mockUrlScrapper.On("GetPhishUrl", "openphish").Return([]string{}, errors.New("")).Once()

		result, err := urlService.FetchUrl("openphish")

		assert.Equal(t, []urls.Domain{}, result)
		assert.NotNil(t, err)
	})
}