package datasets

import (
	"fish-hunter/businesses/datasets"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dataset struct {
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

func FromDomain(domain datasets.Domain) Dataset {
	return Dataset{
		Id:         domain.Id,
		Ref_Url:    domain.Ref_Url,
		Ref_Job:    domain.Ref_Job,
		DateScrapped: domain.DateScrapped,
		HttpStatus: domain.HttpStatus,
		Domain: domain.Domain,
		AssetsDownloaded: domain.AssetsDownloaded,
		ContentLength: domain.ContentLength,
		Url: domain.Url,
		Categories: domain.Categories,
		Brands: domain.Brands,
		DatasetPath: domain.DatasetPath,
		HtmldomPath: domain.HtmldomPath,
		ScrappedFrom: domain.ScrappedFrom,
		UrlscanUuid: domain.UrlscanUuid,
		Status: domain.Status,
		ScreenshotPath: domain.ScreenshotPath,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
	}
}

func FromDomainArray(domain []datasets.Domain) []Dataset {
	var result []Dataset
	for _, v := range domain {
		result = append(result, FromDomain(v))
	}
	return result
}

func (rec *Dataset) ToDomain() datasets.Domain {
	return datasets.Domain{
		Id:         rec.Id,
		Ref_Url:    rec.Ref_Url,
		Ref_Job:    rec.Ref_Job,
		DateScrapped: rec.DateScrapped,
		HttpStatus: rec.HttpStatus,
		Domain: rec.Domain,
		AssetsDownloaded: rec.AssetsDownloaded,
		ContentLength: rec.ContentLength,
		Url: rec.Url,
		Categories: rec.Categories,
		Brands: rec.Brands,
		DatasetPath: rec.DatasetPath,
		HtmldomPath: rec.HtmldomPath,
		ScrappedFrom: rec.ScrappedFrom,
		UrlscanUuid: rec.UrlscanUuid,
		Status: rec.Status,
		ScreenshotPath: rec.ScreenshotPath,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
		DeletedAt:   rec.DeletedAt,
	}
}

func ToDomainArray(rec *[]Dataset) []datasets.Domain {
	var result []datasets.Domain
	for _, v := range *rec {
		result = append(result, v.ToDomain())
	}
	return result
}