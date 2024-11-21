package videoserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type VideoServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *VideoServer {
	return &VideoServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *VideoServer) Start() error {
	if err := s.configLogger(); err != nil {
		return err
	}

	s.configRouter()

	s.logger.Info("starting video server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *VideoServer) configLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *VideoServer) configRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *VideoServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world!")
	}
}
