package logout_handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"patreon/internal/app/delivery/http/handlers"
	"patreon/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type LogoutTestSuite struct {
	handlers.SuiteTestBaseHandler
}

func (s *LogoutTestSuite) TestServeHTTP_WithSession() {
	uniqID := "1"
	test := handlers.TestTable{
		name:              "with cookies",
		data:              &models.User{},
		expectedMockTimes: 1,
		expectedCode:      http.StatusOK,
	}

	recorder := httptest.NewRecorder()
	handler := NewLogoutHandler(s.logger, s.dataStorage)

	b := bytes.Buffer{}
	err := json.NewEncoder(&b).Encode(test.data)

	require.NoError(s.T(), err)
	ctx := context.WithValue(context.Background(), "uniq_id", uniqID)
	reader, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/logout", &b)

	s.mockSessionsManager.EXPECT().Delete(uniqID).Times(test.expectedMockTimes).Return(nil)
	handler.ServeHTTP(recorder, reader)
	assert.Equal(s.T(), test.expectedCode, recorder.Code)
}

func (s *LogoutTestSuite) TestServeHTTP_WithoutCookies() {
	uniqID := "1"
	test := handlers.TestTable{
		name:              "without cookies",
		data:              &models.User{},
		expectedMockTimes: 0,
		expectedCode:      http.StatusInternalServerError,
	}

	recorder := httptest.NewRecorder()
	handler := NewLogoutHandler(s.logger, s.dataStorage)

	b := bytes.Buffer{}
	err := json.NewEncoder(&b).Encode(test.data)

	require.NoError(s.T(), err)
	reader, _ := http.NewRequest(http.MethodPost, "/logout", &b)

	s.mockSessionsManager.EXPECT().Delete(uniqID).Times(test.expectedMockTimes).Return(nil)
	handler.ServeHTTP(recorder, reader)
	assert.Equal(s.T(), test.expectedCode, recorder.Code)
}

func (s *LogoutTestSuite) TestServeHTTP_ErrorSessions() {
	uniqID := "1"
	test := handlers.TestTable{
		name:              "without cookies",
		data:              &models.User{},
		expectedMockTimes: 1,
		expectedCode:      http.StatusInternalServerError,
	}

	recorder := httptest.NewRecorder()
	handler := NewLogoutHandler(s.logger, s.dataStorage)

	b := bytes.Buffer{}
	err := json.NewEncoder(&b).Encode(test.data)

	require.NoError(s.T(), err)
	ctx := context.WithValue(context.Background(), "uniq_id", uniqID)
	reader, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/logout", &b)

	s.mockSessionsManager.EXPECT().Delete(uniqID).Times(test.expectedMockTimes).Return(errors.New(""))
	handler.ServeHTTP(recorder, reader)
	assert.Equal(s.T(), test.expectedCode, recorder.Code)
}