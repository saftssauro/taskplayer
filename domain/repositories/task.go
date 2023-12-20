package repositories

import (
	"github.com/saftssauro/taskplayer/domain/entities"
	"github.com/saftssauro/taskplayer/domain/providers"
)

type iTasksRepository interface {
	Create(task entities.Task) (entities.Task, error)
	List(reportId string) ([]entities.Task, error)
}

type TasksRepository struct {
	remoteProvider providers.RemoteProvider
	iTasksRepository
}

func (reportsRepository TasksRepository) New(remoteProvider providers.RemoteProvider) *TasksRepository {
	return &TasksRepository{
		remoteProvider: remoteProvider,
	}
}
