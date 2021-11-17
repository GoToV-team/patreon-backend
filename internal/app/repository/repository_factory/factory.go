package repository_factory

import (
	"patreon/internal/app"
	repCsrf "patreon/internal/app/csrf/repository/jwt"
	repositoryAccess "patreon/internal/app/repository/access"
	repoAwrds "patreon/internal/app/repository/awards"
	repAwardsPsql "patreon/internal/app/repository/awards/postgresql"
	repCreator "patreon/internal/app/repository/creator"
	repCreatorPsql "patreon/internal/app/repository/creator/postgresql"
	repoFiles "patreon/internal/app/repository/files"
	repoFilesOs "patreon/internal/app/repository/files/os"
	repoInfo "patreon/internal/app/repository/info"
	repInfoPsql "patreon/internal/app/repository/info/postgresql"
	repoLikes "patreon/internal/app/repository/likes"
	repLikesPsql "patreon/internal/app/repository/likes/postgresql"
	repoPayments "patreon/internal/app/repository/payments"
	repoPaymentsPsql "patreon/internal/app/repository/payments/postgresql"
	repoPosts "patreon/internal/app/repository/posts"
	repPostsPsql "patreon/internal/app/repository/posts/postgresql"
	repoPostsData "patreon/internal/app/repository/posts_data"
	repoPostsDataPsql "patreon/internal/app/repository/posts_data/postgresql"
	repoSubscribers "patreon/internal/app/repository/subscribers"
	repUser "patreon/internal/app/repository/user"
	repUserPsql "patreon/internal/app/repository/user/postgresql"

	"github.com/sirupsen/logrus"
)

type RepositoryFactory struct {
	expectedConnections   app.ExpectedConnections
	logger                *logrus.Logger
	userRepository        repUser.Repository
	creatorRepository     repCreator.Repository
	awardsRepository      repoAwrds.Repository
	postsRepository       repoPosts.Repository
	likesRepository       repoLikes.Repository
	PostDataRepository    repoPostsData.Repository
	csrfRepository        repCsrf.Repository
	accessRepository      repositoryAccess.Repository
	subscribersRepository repoSubscribers.Repository
	FilesRepository       repoFiles.Repository
	paymentsRepository    repoPayments.Repository
	infoRepository        repoInfo.Repository
}

func NewRepositoryFactory(logger *logrus.Logger, expectedConnections app.ExpectedConnections) *RepositoryFactory {
	return &RepositoryFactory{
		expectedConnections: expectedConnections,
		logger:              logger,
	}
}

func (f *RepositoryFactory) GetUserRepository() repUser.Repository {
	if f.userRepository == nil {
		f.userRepository = repUserPsql.NewUserRepository(f.expectedConnections.SqlConnection)
	}
	return f.userRepository
}

func (f *RepositoryFactory) GetCreatorRepository() repCreator.Repository {
	if f.creatorRepository == nil {
		f.creatorRepository = repCreatorPsql.NewCreatorRepository(f.expectedConnections.SqlConnection)
	}
	return f.creatorRepository
}

func (f *RepositoryFactory) GetCsrfRepository() repCsrf.Repository {
	if f.csrfRepository == nil {
		f.csrfRepository = repCsrf.NewJwtRepository()
	}
	return f.csrfRepository
}

func (f *RepositoryFactory) GetAwardsRepository() repoAwrds.Repository {
	if f.awardsRepository == nil {
		f.awardsRepository = repAwardsPsql.NewAwardsRepository(f.expectedConnections.SqlConnection)
	}
	return f.awardsRepository
}
func (f *RepositoryFactory) GetAccessRepository() repositoryAccess.Repository {
	if f.accessRepository == nil {
		f.accessRepository = repositoryAccess.NewRedisRepository(f.expectedConnections.AccessRedisPool, f.logger)
	}
	return f.accessRepository
}
func (f *RepositoryFactory) GetSubscribersRepository() repoSubscribers.Repository {
	if f.subscribersRepository == nil {
		f.subscribersRepository = repoSubscribers.NewSubscribersRepository(f.expectedConnections.SqlConnection)
	}
	return f.subscribersRepository
}

func (f *RepositoryFactory) GetPostsRepository() repoPosts.Repository {
	if f.postsRepository == nil {
		f.postsRepository = repPostsPsql.NewPostsRepository(f.expectedConnections.SqlConnection)
	}
	return f.postsRepository
}

func (f *RepositoryFactory) GetLikesRepository() repoLikes.Repository {
	if f.likesRepository == nil {
		f.likesRepository = repLikesPsql.NewLikesRepository(f.expectedConnections.SqlConnection)
	}
	return f.likesRepository
}

func (f *RepositoryFactory) GetPostsDataRepository() repoPostsData.Repository {
	if f.PostDataRepository == nil {
		f.PostDataRepository = repoPostsDataPsql.NewPostsDataRepository(f.expectedConnections.SqlConnection)
	}
	return f.PostDataRepository
}

func (f *RepositoryFactory) GetFilesRepository() repoFiles.Repository {
	if f.FilesRepository == nil {
		f.FilesRepository = repoFilesOs.NewFileRepository(f.expectedConnections.PathFiles)
	}
	return f.FilesRepository
}

func (f *RepositoryFactory) GetPaymentsRepository() repoPayments.Repository {
	if f.paymentsRepository == nil {
		f.paymentsRepository = repoPaymentsPsql.NewPaymentsRepository(f.expectedConnections.SqlConnection)
	}
	return f.paymentsRepository
}

func (f *RepositoryFactory) GetInfoRepository() repoInfo.Repository {
	if f.infoRepository == nil {
		f.infoRepository = repInfoPsql.NewInfoRepository(f.expectedConnections.SqlConnection)
	}
	return f.infoRepository
}
