package server

import (
	"github.com/bohdanabadi/Traffic-Simulation/api/api/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

var env string

type Server struct {
	engine *gin.Engine
	config Config
}

func NewServer(cfg Config) *Server {
	// Here you could add a switch to start the server with TLS or without depending on the configuration
	env = os.Getenv("ENV")
	if env == "dev" || env == "" {
		env = "development"

	}

	srv := &Server{
		engine: gin.Default(),
		config: cfg,
	}
	corsConfig := cors.DefaultConfig()
	//switch env {
	//case "development":
	corsConfig.AllowOrigins = []string{cfg.ServerDev.CrossOrigin}
	//case "production":
	//	corsConfig.AllowOrigins = []string{cfg.ServerProd.CrossOrigin}
	//default:
	//	log.Fatalf("Invalid environment: %s", env)
	//
	//}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	srv.engine.Use(cors.New(corsConfig))
	return srv
}

func (srv *Server) StartServer() error {
	var err error
	v1 := srv.engine.Group("/v1")
	{
		v1.GET("/fe", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "Hi!, Component data not found. Time to useState!"})
		})
		v1.GET("points/random-pair", handler.GetPotentialJourneyPoints)
		v1.POST("journeys", handler.CreateJourney)
		v1.GET("journeys", handler.GetJourney)
		v1.PATCH("journeys/status", handler.UpdateJourneyStatus)
		v1.GET("ws/simulation/path", handler.HandleSimulationConnection)
		v1.GET("ws/fe/path", handler.HandleFrontendConnection)
	}

	switch env {
	case "development":
		err = srv.engine.Run(srv.config.ServerDev.Host + ":" + srv.config.ServerDev.Port)
	case "production":
		err = srv.engine.Run(srv.config.ServerProd.Host + ":" + srv.config.ServerProd.Port)
	default:
		log.Fatalf("Invalid environment: %s", env)
	}
	return err
}
