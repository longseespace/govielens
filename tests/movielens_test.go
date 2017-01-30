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
	err := badClient.Login("test@nomadreact.com", "xxx")
	assert.NotNil(t, err)
}

func TestLogin(t *testing.T) {
	client = movielens.NewClient(nil)
	err := client.Login("test@nomadreact.com", "ddpgvoQuTCCs4VouVQedneuU")
	assert.Nil(t, err)
}

func TestGetMe(t *testing.T) {
	my, err := client.GetMe()
	assert.Nil(t, err)
	assert.Equal(t, "test@nomadreact.com", my.Email)
}

func TestExplore(t *testing.T) {
	params := map[string]string{
		"hasRated": "yes",
	}
	_, err := client.Explore(params)
	assert.Nil(t, err)
}
