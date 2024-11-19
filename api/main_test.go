package api

import (
	"os"
	"testing"
	"time"

	db "github.com/annguyen34/simple-bank/db/sqlc"
	"github.com/annguyen34/simple-bank/util"
	"github.com/gin-gonic/gin"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	if err != nil {
		t.Fatal("cannot create server:", err)
	}

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())

}
