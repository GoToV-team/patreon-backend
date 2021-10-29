package usecase_factory

import (
	repCsrf "patreon/internal/app/csrf/repository/jwt"
	repAccess "patreon/internal/app/repository/access"
	repoAwrds "patreon/internal/app/repository/awards"
	repCreator "patreon/internal/app/repository/creator"
	useSubscr "patreon/internal/app/repository/subscribers"
	repUser "patreon/internal/app/repository/user"

	"patreon/internal/app/sessions"
)

type RepositoryFactory interface {
	GetUserRepository() repUser.Repository
	GetCreatorRepository() repCreator.Repository
	GetAwardsRepository() repoAwrds.Repository
	GetCsrfRepository() repCsrf.Repository
	GetSessionRepository() sessions.SessionRepository
	GetAccessRepository() repAccess.Repository
	GetSubscribersRepository() useSubscr.Repository
}
