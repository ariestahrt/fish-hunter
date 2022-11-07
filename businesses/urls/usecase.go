package urls

import (
	sourceHelper "fish-hunter/helpers/source"
	"fish-hunter/util/scrapper"
)

type UrlUseCase struct {
	UrlRepository Repository
	Scrapper scrapper.Scrapper
}

func NewUrlUseCase(urlRepository Repository, scrapper scrapper.Scrapper) UseCase {
	return &UrlUseCase{
		UrlRepository: urlRepository,
		Scrapper: scrapper,
	}
}

func (u *UrlUseCase) GetAll() ([]Domain, error) {
	return u.UrlRepository.GetAll()
}

func (u *UrlUseCase) FetchUrl(source string) ([]Domain, error) {
	// Get url from helper
	urls, err := u.Scrapper.GetPhishUrl(source)
	if err != nil {
		return []Domain{}, err
	}

	// Get source information from database
	src := sourceHelper.GetSourceInformation(source)
	var urlDomains []Domain
	for _, url := range urls {
		// Save to database
		res, _ := u.UrlRepository.Save(Domain{
			Url: url,
			Source_Url: src.Url,
			Source_Name: src.Name,
			Ref_Source: src.Id,
		})
		urlDomains = append(urlDomains, res)
	}

	return urlDomains, nil
}

func (u *UrlUseCase) GetByID(id string) (Domain, error) {
	return u.UrlRepository.GetByID(id)
}