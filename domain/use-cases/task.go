package usecases

import (
	"github.com/saftssauro/taskplayer/domain/entities"
	"github.com/saftssauro/taskplayer/domain/repositories"
)

type TasksUseCases struct {
	tasksRepository repositories.TasksRepository
}

func (tasksUseCases TasksUseCases) New(tasksRepository repositories.TasksRepository) *TasksUseCases {
	return &TasksUseCases{
		tasksRepository: tasksRepository,
	}
}

func (tasksUseCases *TasksUseCases) List(reportId string) ([]entities.Task, error) {
	tasks, err := tasksUseCases.tasksRepository.List(reportId)

	return tasks, err
}

func (tasksUseCases *TasksUseCases) Create(report entities.Task) (entities.Task, error) {
	tasks, err := tasksUseCases.tasksRepository.Create(report)

	return tasks, err
}
