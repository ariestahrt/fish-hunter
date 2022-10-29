package urls

import (
	"fish-hunter/helpers/scrapper"
	sourceHelper "fish-hunter/helpers/source"
)

type UrlUseCase struct {
	UrlRepository Repository
}

func NewUrlUseCase(urlRepository Repository) UseCase {
	return &UrlUseCase{
		UrlRepository: urlRepository,
	}
}

func (u *UrlUseCase) GetAll() ([]Domain, error) {
	return u.UrlRepository.GetAll()
}

func (u *UrlUseCase) FetchUrl(source string) ([]Domain, error) {
	// Get url from helper
	urls, err := scrapper.GetPhishUrl(source)
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