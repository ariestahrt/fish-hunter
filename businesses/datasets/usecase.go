package datasets

import (
	"fish-hunter/util"
	"fish-hunter/util/datasetutil"
	"fish-hunter/util/s3"
	"os"
	"time"
)

type DatasetsUseCase struct {
	DatasetsRepository Repository
	S3                 s3.AWS_S3
	datasetUtil			datasetutil.DatasetUtil
}

func NewDatasetUseCase(datasetsRepository Repository, s3 s3.AWS_S3, datasetUtil datasetutil.DatasetUtil) UseCase {
	return &DatasetsUseCase{
		DatasetsRepository: datasetsRepository,
		S3:	s3,
		datasetUtil: datasetUtil,
	}
}

func (u *DatasetsUseCase) Status(status string) ([]Domain, error) {
	return u.DatasetsRepository.Status(status)
}

func (u *DatasetsUseCase) GetByID(id string) (Domain, error) {
	return u.DatasetsRepository.GetByID(id)
}

func (u *DatasetsUseCase) Validate(domain Domain) (Domain, error) {
	return u.DatasetsRepository.Validate(domain)
}

func (u *DatasetsUseCase) TopBrands() (map[string]interface{}, error) {
	return u.DatasetsRepository.TopBrands()
}

func (u *DatasetsUseCase) Download(id string) (string, error) {
	// get dataset by id
	dataset, err := u.DatasetsRepository.GetByID(id)
	if err != nil {
		return "", err
	}

	// Download from s3
	file7z := util.GetConfig("APP_PATH") + "files/" + dataset.Ref_Url.Hex() + ".7z"
	file, err := os.Create(file7z)
	if err != nil {
		return "", err
	}
	defer file.Close()

	err = u.S3.DownloadFile(file, dataset.FolderPath + ".7z")
	if err != nil {
		return "", err
	}

	// Unzip
	err = u.datasetUtil.Extract7Zip(file7z, util.GetConfig("7Z_PASSWORD"))
	if err != nil {
		return "", err
	}

	// Compress
	folder_to_compress := util.GetConfig("APP_PATH") + "files/datasets/" + dataset.Ref_Url.Hex()
	err = u.datasetUtil.Compress7Zip(folder_to_compress)
	if err != nil {
		return "", err
	}
	// Remove folder
	os.RemoveAll(folder_to_compress)

	// Prune
	go func() {
		// Sleep 2 minute
		time.Sleep(2 * time.Minute)
		os.RemoveAll(folder_to_compress + ".7z")
	}()

	return folder_to_compress + ".7z", nil
}

func (u *DatasetsUseCase) Activate(id string) (string, error) {
	// get dataset by id
	dataset, err := u.DatasetsRepository.GetByID(id)
	if err != nil {
		return "", err
	}

	// Download from s3
	file7z := util.GetConfig("APP_PATH") + "files/" + dataset.Ref_Url.Hex() + ".7z"
	file, err := os.Create(file7z)
	if err != nil {
		return "", err
	}
	defer file.Close()

	err = u.S3.DownloadFile(file, dataset.FolderPath + ".7z")
	if err != nil {
		return "", err
	}

	// Unzip
	u.datasetUtil.Extract7Zip(file7z, util.GetConfig("7Z_PASSWORD"))

	go u.datasetUtil.TimedPruneDirectory("files/"+dataset.FolderPath, 30)
	return "/datasets/view/" + dataset.Ref_Url.Hex() + "/index.html", nil
}