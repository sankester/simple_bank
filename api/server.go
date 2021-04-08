package api

import (
	"github.com/go-playground/validator/v10"
	db "github.com/sankester/simple_bank/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

// create new server
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// binding custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	// setup router
	router.POST("/users", server.createUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.POST("/transfers", server.createTransfer)

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
