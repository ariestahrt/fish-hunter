package stats

import (
	"fish-hunter/businesses/datasets"
	"fish-hunter/businesses/jobs"
	"fish-hunter/businesses/urls"
	"time"
)

type StatsUseCase struct {
	DatasetsRepository datasets.Repository
	JobsRepository     jobs.Repository
	UrlsRepository     urls.Repository
}

func NewStatUseCase(datasetsRepository datasets.Repository, jobsRepository jobs.Repository, urlsRepository urls.Repository) UseCase {
	return &StatsUseCase{
		DatasetsRepository: datasetsRepository,
		JobsRepository:     jobsRepository,
		UrlsRepository:     urlsRepository,
	}
}

func (s *StatsUseCase) GetStatistics() (map[string]interface{}, error) {
	
	// Total Urls
	totalUrls, _ := s.UrlsRepository.CountTotal()

	// Total Datasets
	totalDatasets, _ := s.DatasetsRepository.CountTotal()
	
	// Total Validated Datasets
	totalValidatedDatasets, _ := s.DatasetsRepository.CountTotalValid()

	// Total Jobs
	totalJobs, _ := s.JobsRepository.CountTotal()

	return map[string]interface{}{
		"urls": totalUrls,
		"datasets": totalDatasets,
		"valid_datasets": totalValidatedDatasets,
		"jobs": totalJobs,
	}, nil
}

func (s *StatsUseCase) GetLastWeekStatistics() (map[string]interface{}, error) {
	resp := map[string]interface{}{}
	resp["date"] = []string{}
	resp["total_url"] = []int{}
	resp["total_job"] = []int{}
	resp["total_dataset"] = []int{}

	t := time.Now().UTC()
	t = t.Truncate(24 * time.Hour)
	t = t.Add(time.Hour * 24 * 1)
	
	for i := 0; i < 7; i++ {
		t_next := t
		t = t.Add(time.Hour * 24 * -1)
		resp["date"] = append(resp["date"].([]string), t.Format("02-01-2006"))

		// Total Urls
		totalUrls, _ := s.UrlsRepository.GetTotalBetweenDates(t, t_next)
		resp["total_url"] = append(resp["total_url"].([]int), int(totalUrls))

		// Total Datasets
		totalDatasets, _ := s.DatasetsRepository.GetTotalBetweenDates(t, t_next)
		resp["total_dataset"] = append(resp["total_dataset"].([]int), int(totalDatasets))

		// Total Jobs
		totalJobs, _ := s.JobsRepository.GetTotalBetweenDates(t, t_next)
		resp["total_job"] = append(resp["total_job"].([]int), int(totalJobs))
	}

	return resp, nil
}