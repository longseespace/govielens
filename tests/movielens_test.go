package tests

import (
	"testing"

	"github.com/longseespace/govielens/movielens"
	"github.com/stretchr/testify/assert"
)

var (
	client *movielens.Client
)

func TestBadLogin(t *testing.T) {
	badClient := movielens.NewClient(nil)
	_, err := badClient.Login("test@nomadreact.com", "xxx")
	assert.NotNil(t, err)
}

func TestLogin(t *testing.T) {
	client = movielens.NewClient(nil)
	_, err := client.Login("test@nomadreact.com", "ddpgvoQuTCCs4VouVQedneuU")
	assert.Nil(t, err)
}

func TestGetUser(t *testing.T) {
	user, err := client.GetMe()
	assert.Nil(t, err)
	assert.Equal(t, "test@nomadreact.com", user.Email)
}
