package samples

type SamplesUseCase struct {
	SamplesRepository Repository
}

func NewSampleUseCase(samplesRepository Repository) UseCase {
	return &SamplesUseCase{
		SamplesRepository: samplesRepository,
	}
}

func (u *SamplesUseCase) GetAll() ([]Domain, error) {
	return u.SamplesRepository.GetAll()
}

func (u *SamplesUseCase) GetByID(id string) (Domain, error) {
	return u.SamplesRepository.GetByID(id)
}

func (u *SamplesUseCase) Update(id string, sampleDomain *Domain) (Domain, error) {
	return u.SamplesRepository.Update(id, sampleDomain)
}