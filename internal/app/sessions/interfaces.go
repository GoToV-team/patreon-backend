package sessions

import "patreon/internal/app/sessions/models"

type IRepository interface {
	Set(session *models.Session) error
	GetUserId(key string) (string, error)
	Del(session *models.Session) error
}

type ISessionsManager interface {
	CheckSession(uniqID string) (models.Result, error)
	CreateSession(userID int64) (models.Result, error)
	DeleteSession(uniqID string) error
}
