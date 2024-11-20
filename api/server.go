package api

import (
	"log"

	db "github.com/annguyen34/simple-bank/db/sqlc"
	"github.com/annguyen34/simple-bank/token"
	"github.com/annguyen34/simple-bank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	// Server configuration
	config     util.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("cannot create token maker:", err)
	}

	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
		store:      store}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, err
}

func (server *Server) example(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/example", server.example)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)
	authRoutes.POST("/transfers", server.createTransfer)

	router.POST("/login", server.loginUser)
	router.POST("/signup", server.createUser)
	router.POST("/token/renew_access", server.renewAccessToken)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
