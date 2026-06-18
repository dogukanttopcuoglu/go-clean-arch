package httpserver

import (
	"github.com/dogukanttopcuoglu/clean-lab/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

const defaultAddr = ":8080"

type Server struct {
	App     *fiber.App
	address string
	log     logger.Logger
}

func New(log logger.Logger, opts ...Option) *Server {
	server := &Server{
		App:     fiber.New(),
		address: defaultAddr,
		log:     log,
	}
	for _, opt := range opts {
		opt(server)
	}

	return server
}

func (s *Server) Start() error {
	s.log.Info("http server listening on " + s.address)
	return s.App.Listen(s.address)
}
