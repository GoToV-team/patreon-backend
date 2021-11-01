package server

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"patreon/internal/app/delivery/http/handler_factory"
	"patreon/internal/app/repository/repository_factory"
	"patreon/internal/app/usecase/usecase_factory"

	"golang.org/x/crypto/acme/autocert"

	_ "patreon/docs"
	"patreon/internal/app"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	config      *app.Config
	logger      *log.Logger
	connections app.ExpectedConnections
}

func New(config *app.Config, connections app.ExpectedConnections, logger *log.Logger) *Server {
	return &Server{
		config:      config,
		logger:      logger,
		connections: connections,
	}
}

func (s *Server) checkConnection() error {
	if err := s.connections.SqlConnection.Ping(); err != nil {
		return fmt.Errorf("Can't check connection to sql with error %v ", err)
	}

	s.logger.Info("Success check connection to sql db")

	connSession, err := s.connections.SessionRedisPool.Dial()
	if err != nil {
		return fmt.Errorf("Can't check connection to redis with error: %v ", err)
	}

	s.logger.Info("Success check connection to redis")

	err = connSession.Close()
	if err != nil {
		return fmt.Errorf("Can't close connection to redis with error: %v ", err)
	}

	connAccess, err := s.connections.SessionRedisPool.Dial()
	if err != nil {
		return fmt.Errorf("Can't check connection to redis with error: %v ", err)
	}

	s.logger.Info("Success check connection to redis")

	err = connAccess.Close()
	if err != nil {
		return fmt.Errorf("Can't close connection to redis with error: %v ", err)
	}

	return nil
}

// @title Patreon
// @version 1.0
// @description Server for Patreon application.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /api/v1/

// @x-extension-openapi {"example": "value on a json format"}

func (s *Server) Start(config *app.Config) error {
	if err := s.checkConnection(); err != nil {
		return err
	}

	router := mux.NewRouter()
	routerApi := router.PathPrefix("/api/v1/").Subrouter()
	routerApi.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	fileServer := http.FileServer(http.Dir(config.MediaDir + "/"))
	routerApi.PathPrefix("/" + app.LoadFileUrl).Handler(http.StripPrefix("/api/v1/"+app.LoadFileUrl, fileServer))

	repositoryFactory := repository_factory.NewRepositoryFactory(s.logger, s.connections)
	usecaseFactory := usecase_factory.NewUsecaseFactory(repositoryFactory)
	factory := handler_factory.NewFactory(s.logger, router, &config.Cors, usecaseFactory)
	hs := factory.GetHandleUrls()

	for apiUrl, h := range *hs {
		h.Connect(routerApi.Path(apiUrl))
	}

	s.logger.Info("starting server")

	server := &http.Server{
		Addr: config.BindHttpAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			targetUrl := url.URL{Scheme: "https", Host: r.Host, Path: r.URL.Path, RawQuery: r.URL.RawQuery}
			log.Infof("Redirect from %s, to %s", r.URL.String(), targetUrl.String())
			http.Redirect(w, r, targetUrl.String(), http.StatusPermanentRedirect)
		}),
	}
	if config.IsProduction {
		m := &autocert.Manager{
			Cache:      autocert.DirCache("patreon-secrt"),
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(config.Domen, "www."+config.Domen),
		}
		server = &http.Server{
			Addr:      config.BindHttpsAddr,
			TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
			Handler:   routerApi,
		}
		return server.ListenAndServeTLS("", "")
	} else {
		server.Handler = routerApi
		return server.ListenAndServe()
		//return http.ListenAndServe(config.BindAddr, routerApi)
	}
}
