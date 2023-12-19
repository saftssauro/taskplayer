package repositories

import "github.com/saftssauro/taskplayer/domain/providers"

type iUsersRepository interface{}

type UsersRepository struct {
	iUsersRepository
	remoteProvider providers.RemoteProvider
}

func (reportsRepository UsersRepository) New(remoteProvider providers.RemoteProvider) *UsersRepository {
	return &UsersRepository{
		remoteProvider: remoteProvider,
	}
}
