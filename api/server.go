package api

import (
	"github.com/gin-gonic/gin"
	yogo "yogo.io/go-tripPlanner-backend/db"
)

// Servers HTTP requests
type Server struct {
	store  yogo.Store
	router *gin.Engine
}

//creates new HTTP server and setup routing
func NewServer(store yogo.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add routes to router
	router.POST("/api/users/create", server.createUser)
	router.GET("/api/users/:id", server.getUser)
	router.GET("/api/users", server.listUsers)
	server.router = router
	return server
}

//run HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
