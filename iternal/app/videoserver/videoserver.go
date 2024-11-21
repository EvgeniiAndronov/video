package videoserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"video/iternal/app/store"
)

type VideoServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
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

	if err := s.configureStore(); err != nil {
		return err
	}

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

func (s *VideoServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st

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
