package gapi

import (
	"log"

	db "github.com/annguyen34/simple-bank/db/sqlc"
	"github.com/annguyen34/simple-bank/pb"
	"github.com/annguyen34/simple-bank/token"
	"github.com/annguyen34/simple-bank/util"
	"github.com/annguyen34/simple-bank/worker"
)

// Server serve gRPC requests
type Server struct {
	// Server configuration
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("cannot create token maker:", err)
	}

	server := &Server{
		config:          config,
		tokenMaker:      tokenMaker,
		store:           store,
		taskDistributor: taskDistributor}

	return server, err
}
