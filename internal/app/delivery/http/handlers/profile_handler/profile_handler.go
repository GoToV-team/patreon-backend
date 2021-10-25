package profile_handler

import (
	"io"
	"net/http"
	"patreon/internal/app"
	bh "patreon/internal/app/delivery/http/handlers/base_handler"
	"patreon/internal/app/delivery/http/handlers/handler_errors"
	"patreon/internal/app/models"
	"patreon/internal/app/sessions"
	"patreon/internal/app/sessions/middleware"
	usecase_user "patreon/internal/app/usecase/user"

	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

type ProfileHandler struct {
	sessionManager sessions.SessionsManager
	userUsecase    usecase_user.Usecase
	bh.BaseHandler
}

func NewProfileHandler(log *logrus.Logger, router *mux.Router, cors *app.CorsConfig,
	sManager sessions.SessionsManager, ucUser usecase_user.Usecase) *ProfileHandler {
	h := &ProfileHandler{
		sessionManager: sManager,
		userUsecase:    ucUser,
		BaseHandler:    *bh.NewBaseHandler(log, router, cors),
	}
	h.AddMethod(http.MethodGet, h.GET,
		middleware.NewSessionMiddleware(h.sessionManager, log).Check,
	)
	return h
}

// GET Profile
// @Summary get information from user for user
// @Description get nickname and avatar for user
// @Accept  json
// @Produce json
// @Success 201 {object} models.ProfileResponse "Successfully get user"
// @Failure 401 "User are not authorized"
// @Failure 404 {object} models.ErrResponse "user with this id not found"
// @Failure 500 {object} models.ErrResponse "can not do bd operation"
// @Failure 500 {object} models.ErrResponse "can not get info from context"
// @Router /user [GET]
func (h *ProfileHandler) GET(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			h.Log(r).Error(err)
		}
	}(r.Body)

	userID := r.Context().Value("user_id")
	if userID == nil {
		h.Log(r).Error("can not get user_id from context")
		h.Error(w, r, http.StatusInternalServerError, handler_errors.ContextError)
		return
	}

	u, err := h.userUsecase.GetProfile(userID.(int64))
	if err != nil {
		h.UsecaseError(w, r, err, codeByError)
		return
	}

	h.Log(r).Debugf("get user %s", u)
	h.Respond(w, r, http.StatusOK, models.Profile{Nickname: u.Nickname, Avatar: u.Avatar})
}
