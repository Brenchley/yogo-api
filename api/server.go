package api

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	yogo "yogo.io/go-tripPlanner-backend/db"
	"yogo.io/go-tripPlanner-backend/token"
	"yogo.io/go-tripPlanner-backend/util"
)

// Servers HTTP requests
type Server struct {
	config     util.Config
	store      yogo.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

//creates new HTTP server and setup routing
func NewServer(config util.Config, store yogo.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	// tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	router.Use(cors.Default())

	// add routes to router
	router.POST("/api/users/login", server.loginUser)
	router.POST("/api/users/create", server.createUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.GET("/api/users/:id", server.getUser)
	authRoutes.GET("/api/users", server.listUsers)

	server.router = router
}

//run HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
