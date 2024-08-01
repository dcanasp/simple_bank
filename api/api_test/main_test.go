package api_test

import (
	"fmt"
	"net/http"
	"os"
	"simple_bank/api"
	"simple_bank/db/myDb"
	"simple_bank/token"
	"simple_bank/util"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store myDb.Store) *api.Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := api.NewServer(config, store)
	require.NoError(t, err)

	return server
}

func createAndSetAuthToken(t *testing.T, request *http.Request, tokenMaker token.Maker, username string) {
	if len(username) == 0 {
		return
	}

	token, err := tokenMaker.CreateToken(username, time.Minute)
	require.NoError(t, err)

	authorizationHeader := fmt.Sprintf("%s %s", api.AuthorizationTypeBearer, token)
	request.Header.Set(api.AuthorizationHeaderKey, authorizationHeader)
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
