package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"patreon/internal/models"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type LogoutTestSuit struct {
	SuiteTestStore
}

func TestTestStore2(t *testing.T) {
	suite.Run(t, new(LogoutTestSuit))
}

func (s *LogoutTestSuit) TestServeHTTP_WithSession() {
	uniqID := "1"
	test := TestTable{
		name:              "with cookies",
		data:              &models.User{},
		expectedMockTimes: 1,
		expectedCode:      http.StatusOK,
	}

	recorder := httptest.NewRecorder()
	handler := NewLogoutHandler()
	logger := logrus.New()
	str := bytes.Buffer{}
	logger.SetOutput(&str)

	handler.SetLogger(logger)
	handler.SetSessionManager(s.mockSessionsManager)

	b := bytes.Buffer{}
	err := json.NewEncoder(&b).Encode(test.data)

	require.NoError(s.T(), err)
	ctx := context.WithValue(context.Background(), "uniq_id", uniqID)
	reader, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/login", &b)

	s.mockSessionsManager.EXPECT().Delete(uniqID).Times(test.expectedMockTimes).Return(nil)
	handler.ServeHTTP(recorder, reader)
	assert.Equal(s.T(), test.expectedCode, recorder.Code)
}

func (s *LogoutTestSuit) TestServeHTTP_WithoutCookies() {
	uniqID := "1"
	test := TestTable{
		name:              "without cookies",
		data:              &models.User{},
		expectedMockTimes: 0,
		expectedCode:      http.StatusInternalServerError,
	}

	recorder := httptest.NewRecorder()
	handler := NewLogoutHandler()
	logger := logrus.New()
	str := bytes.Buffer{}
	logger.SetOutput(&str)

	handler.SetLogger(logger)
	handler.SetSessionManager(s.mockSessionsManager)

	b := bytes.Buffer{}
	err := json.NewEncoder(&b).Encode(test.data)

	require.NoError(s.T(), err)
	reader, _ := http.NewRequest(http.MethodPost, "/login", &b)

	s.mockSessionsManager.EXPECT().Delete(uniqID).Times(test.expectedMockTimes).Return(nil)
	handler.ServeHTTP(recorder, reader)
	assert.Equal(s.T(), test.expectedCode, recorder.Code)
}

func (s *LogoutTestSuit) TestServeHTTP_ErrorSessions() {
	uniqID := "1"
	test := TestTable{
		name:              "without cookies",
		data:              &models.User{},
		expectedMockTimes: 1,
		expectedCode:      http.StatusInternalServerError,
	}

	recorder := httptest.NewRecorder()
	handler := NewLogoutHandler()
	logger := logrus.New()
	str := bytes.Buffer{}
	logger.SetOutput(&str)

	handler.SetLogger(logger)
	handler.SetSessionManager(s.mockSessionsManager)

	b := bytes.Buffer{}
	err := json.NewEncoder(&b).Encode(test.data)

	require.NoError(s.T(), err)
	ctx := context.WithValue(context.Background(), "uniq_id", uniqID)
	reader, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/login", &b)

	s.mockSessionsManager.EXPECT().Delete(uniqID).Times(test.expectedMockTimes).Return(errors.New(""))
	handler.ServeHTTP(recorder, reader)
	assert.Equal(s.T(), test.expectedCode, recorder.Code)
}
