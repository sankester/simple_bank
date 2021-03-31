package api

import (
	db "backend-master-class/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// create new server
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// setup router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// Start run http server on a spectific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// format error response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
