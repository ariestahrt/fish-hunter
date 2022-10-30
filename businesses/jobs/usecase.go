package jobs

type JobsUseCase struct {
	JobsRepository Repository
}

func NewJobUseCase(jobsRepository Repository) UseCase {
	return &JobsUseCase{
		JobsRepository: jobsRepository,
	}
}

func (u *JobsUseCase) GetAll() ([]Domain, error) {
	return u.JobsRepository.GetAll()
}

func (u *JobsUseCase) GetByID(id string) (Domain, error) {
	return u.JobsRepository.GetByID(id)
}