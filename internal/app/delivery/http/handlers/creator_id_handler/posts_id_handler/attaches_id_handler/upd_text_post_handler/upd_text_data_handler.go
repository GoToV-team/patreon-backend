package upd_text_data_handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	csrf_middleware "patreon/internal/app/csrf/middleware"
	repository_jwt "patreon/internal/app/csrf/repository/jwt"
	usecase_csrf "patreon/internal/app/csrf/usecase"
	bh "patreon/internal/app/delivery/http/handlers/base_handler"
	"patreon/internal/app/delivery/http/handlers/handler_errors"
	"patreon/internal/app/delivery/http/models"
	"patreon/internal/app/middleware"
	models_db "patreon/internal/app/models"
	"patreon/internal/app/sessions"
	sessionMid "patreon/internal/app/sessions/middleware"
	useAttaches "patreon/internal/app/usecase/attaches"
	usePosts "patreon/internal/app/usecase/posts"

	"github.com/sirupsen/logrus"
)

type AttachesUpdateTextHandler struct {
	attachesUsecase useAttaches.Usecase
	bh.BaseHandler
}

func NewAttachesUpdateTextHandler(log *logrus.Logger,
	ucAttaches useAttaches.Usecase, ucPosts usePosts.Usecase,
	manager sessions.SessionsManager) *AttachesUpdateTextHandler {
	h := &AttachesUpdateTextHandler{
		BaseHandler:      *bh.NewBaseHandler(log),
		attachesUsecase: ucAttaches,
	}
	sessionMiddleware := sessionMid.NewSessionMiddleware(manager, log)
	h.AddMiddleware(sessionMiddleware.Check, csrf_middleware.NewCsrfMiddleware(log, usecase_csrf.
		NewCsrfUsecase(repository_jwt.NewJwtRepository())).CheckCsrfToken,
		middleware.NewCreatorsMiddleware(log).CheckAllowUser,
		middleware.NewPostsMiddleware(log, ucPosts).CheckCorrectPost,
		middleware.NewAttachesMiddleware(log, ucAttaches).CheckCorrectAttach)
	h.AddMethod(http.MethodPut, h.PUT)
	return h
}

// PUT update text to post
// @Summary update text to post
// @Accept  json
// @Param attach_text body http_models.RequestText true "Request body for text"
// @Success 200
// @Failure 500 {object} http_models.ErrResponse "can not do bd operation", "server error"
// @Failure 422 {object} http_models.ErrResponse "this post id not know"
// @Failure 404 {object} http_models.ErrResponse "attach with this id not found"
// @Failure 400 {object} http_models.ErrResponse "invalid parameters", "invalid data type", "invalid body in request"
// @Failure 403 {object} http_models.ErrResponse "for this user forbidden change creator", "this post not belongs this creators", "csrf token is invalid, get new token"
// @Failure 401 "user are not authorized"
// @Router /creators/{:creator_id}/posts/{:post_id}/{:attach_id}/update/text [PUT]
func (h *AttachesUpdateTextHandler) PUT(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			h.Log(r).Error(err)
		}
	}(r.Body)

	var attachId int64
	var ok bool

	if attachId, ok = h.GetInt64FromParam(w, r, "attach_id"); !ok {
		return
	}

	if len(mux.Vars(r)) > 3 {
		h.Log(r).Warnf("Too many parametres %v", mux.Vars(r))
		h.Error(w, r, http.StatusBadRequest, handler_errors.InvalidParameters)
		return
	}

	req := &http_models.RequestText{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(req); err != nil {
		h.Log(r).Warnf("can not parse request %s", err)
		h.Error(w, r, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return
	}

	err := h.attachesUsecase.UpdateText(&models_db.AttachWithoutLevel{ID: attachId, Value: req.Text})
	if err != nil {
		h.UsecaseError(w, r, err, codeByErrorPUT)
		return
	}

	w.WriteHeader(http.StatusOK)
}
