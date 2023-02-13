package server

import (
	"fmt"

	"github.com/SAMCHEE/webserver-go/middleware"
	"github.com/SAMCHEE/webserver-go/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

type Port uint32

func (p *Port) string() string { return fmt.Sprintf(":%d", *p) }

type Server struct {
	port Port
	app  *fiber.App
}

func fiberConfig() fiber.Config {
	appName := "my webserver"
	return fiber.Config{
		AppName:      appName,
		ServerHeader: appName,
		// html template를 사용하는경우
		Views: engine(),
	}
}

func NewServer(port uint32) *Server {
	s := &Server{}
	s.port = Port(port)
	s.app = fiber.New(fiberConfig())
	return s
}

func (s *Server) set() {
	// static file
	s.app.Static("/static", "./static")
}

func (s *Server) middlewares() {
	s.app.Use("/", compress.New(), middleware.BindSiteConfig)
	// my middlewares
	// valid api-key
	// ..
}

func (s *Server) routes() {
	route.HandleIndex(s.app.Group("/"))
	// route.HandleIndex(s.app.Group("/store")) // example   /store
}

func (s *Server) Run() error {
	s.set()
	s.middlewares()
	s.routes()
	return s.app.Listen(s.port.string())
}
