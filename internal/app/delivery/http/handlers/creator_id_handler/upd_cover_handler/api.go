package upd_cover_creator_handler

import (
	"net/http"
	"patreon/internal/app/delivery/http/handlers/base_handler"
	"patreon/internal/app/delivery/http/handlers/handler_errors"
	"patreon/internal/app/repository"
	repository_os "patreon/internal/microservices/files/files/repository/files/os"
	"patreon/pkg/utils"

	log "github.com/sirupsen/logrus"
)

var codeByError = base_handler.CodeMap{
	repository.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, log.ErrorLevel},
	repository.NotFound: {
		http.StatusUnprocessableEntity, handler_errors.IncorrectCreatorId, log.WarnLevel},
	repository_os.ErrorCreate: {
		http.StatusInternalServerError, handler_errors.InternalError, log.ErrorLevel},
	repository_os.ErrorCreate: {
		http.StatusInternalServerError, handler_errors.InternalError, log.ErrorLevel},
	utils.ConvertErr: {
		http.StatusInternalServerError, handler_errors.InternalError, log.ErrorLevel},
	utils.UnknownExtOfFileName: {
		http.StatusInternalServerError, handler_errors.InternalError, log.ErrorLevel},
}
