package api

import (
	db "github.com/annguyen34/simple-bank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	// Server configuration
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// Add routes
	router.GET("/", server.example)
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	router.POST("/transfers", server.createTransfer)
	router.POST("/users", server.createUser)
	server.router = router
	return server
}

func (server *Server) example(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
