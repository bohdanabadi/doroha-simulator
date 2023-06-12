package config

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type Server struct {
	engine *gin.Engine
	config Config
}

func NewServer(cfg Config) *Server {
	srv := &Server{
		engine: gin.Default(),
		config: cfg,
	}
	return srv
}

func (srv *Server) StartServer() error {
	var err error

	srv.engine.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hi!, Keep your Gin up")
	})

	// Here you could add a switch to start the server with TLS or without depending on the configuration
	env := os.Getenv("ENV")
	if env == "dev" || env == "" {
		env = "development"
	}
	switch env {
	case "development":
		err = srv.engine.Run(srv.config.ServerDev.Host + ":" + srv.config.ServerDev.Port)
	case "production":
		err = srv.engine.RunTLS(srv.config.ServerProd.Host+":"+srv.config.ServerProd.Port, srv.config.CertFile, srv.config.KeyFile)
	default:
		log.Fatalf("Invalid environment: %s", env)
	}
	return err
}
