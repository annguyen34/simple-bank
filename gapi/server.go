package gapi

import (
	"log"

	db "github.com/annguyen34/simple-bank/db/sqlc"
	"github.com/annguyen34/simple-bank/pb"
	"github.com/annguyen34/simple-bank/token"
	"github.com/annguyen34/simple-bank/util"
)

// Server serve gRPC requests
type Server struct {
	// Server configuration
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
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

	return server, err
}
