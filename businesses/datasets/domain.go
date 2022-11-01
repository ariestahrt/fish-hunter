package datasets

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Domain struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Ref_Url    primitive.ObjectID `json:"-" bson:"ref_url,omitempty"`
	Ref_Job    primitive.ObjectID `json:"-" bson:"ref_job,omitempty"`
	DateScrapped primitive.DateTime `json:"date_scrapped,omitempty" bson:"date_scrapped,omitempty"`
	HttpStatus int `json:"http_status,omitempty" bson:"http_status,omitempty"`
	Domain string `json:"domain,omitempty" bson:"domain,omitempty"`
	AssetsDownloaded float64 `json:"assets_downloaded,omitempty" bson:"assets_downloaded,omitempty"`
	ContentLength int `json:"content_length,omitempty" bson:"content_length,omitempty"`
	Url string `json:"url,omitempty" bson:"url,omitempty"`
	Categories []string `json:"categories,omitempty" bson:"categories,omitempty"`
	Brands []string `json:"brands,omitempty" bson:"brands,omitempty"`
	DatasetPath string `json:"dataset_path,omitempty" bson:"dataset_path,omitempty"`
	HtmldomPath string `json:"htmldom_path,omitempty" bson:"htmldom_path,omitempty"`
	ScrappedFrom string `json:"scrapped_from,omitempty" bson:"scrapped_from,omitempty"`
	UrlscanUuid string `json:"urlscan_uuid,omitempty" bson:"urlscan_uuid,omitempty"`
	Status string `json:"status,omitempty" bson:"status,omitempty"`
	ScreenshotPath string `json:"screenshot_path,omitempty" bson:"screenshot_path,omitempty"`
	CreatedAt  primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt   primitive.DateTime `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type UseCase interface {
	Status(status string) ([]Domain, error)
	GetByID(id string) (Domain, error)
	Activate(id string) (string, error)
	Validate(domain Domain) (Domain, error)
	Download(id string) (string, error)
	TopBrands() (map[string]interface{}, error)
}

type Repository interface {
	Status(status string) ([]Domain, error)
	GetByID(id string) (Domain, error)
	Validate(domain Domain) (Domain, error)
	TopBrands() (map[string]interface{}, error)
	CountTotal() (int64, error)
	CountTotalValid() (int64, error)
	GetTotalBetweenDates(start time.Time, end time.Time) (int64, error)
}