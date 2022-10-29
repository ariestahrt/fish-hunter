package cron

import appjwt "fish-hunter/util/jwt"

type CronUseCase struct {
	cronRepository Repository
}

func NewCronUseCase(cronRepository Repository) UseCase {
	return &CronUseCase{
		cronRepository: cronRepository,
	}
}

func (c *CronUseCase) CleanUpToken() {
	appjwt.CleanExpiredToken()
}