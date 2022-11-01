package stats

type UseCase interface {
	GetStatistics() (map[string]interface{}, error)
	GetLastWeekStatistics() (map[string]interface{}, error)
}