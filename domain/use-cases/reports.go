package usecases

import (
	"github.com/saftssauro/taskplayer/domain/entities"
	"github.com/saftssauro/taskplayer/domain/repositories"
)

type ReportsUseCases struct {
	reportsRepository repositories.ReportsRepository
}

func (reportsUseCases ReportsUseCases) New(reportsRepository repositories.ReportsRepository) *ReportsUseCases {
	return &ReportsUseCases{
		reportsRepository: reportsRepository,
	}
}

func (reportsUseCases *ReportsUseCases) List(userId string) ([]entities.Report, error) {
	reports, err := reportsUseCases.reportsRepository.List(userId)

	return reports, err
}

func (reportsUseCases *ReportsUseCases) Create(report entities.Report) (entities.Report, error) {
	reports, err := reportsUseCases.reportsRepository.Create(report)

	return reports, err
}
