package repositories

import (
	"github.com/saftssauro/taskplayer/domain/entities"
	"github.com/saftssauro/taskplayer/domain/providers"
)

type iReportsRepository interface {
	Create(report entities.Report) (entities.Report, error)
	List(userId string) ([]entities.Report, error)
}

type ReportsRepository struct {
	iReportsRepository
	remoteProvider providers.RemoteProvider
}

func (reportsRepository ReportsRepository) New(remoteProvider providers.RemoteProvider) *ReportsRepository {
	return &ReportsRepository{
		remoteProvider: remoteProvider,
	}
}
