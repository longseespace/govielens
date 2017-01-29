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
	client = movielens.NewClient(nil)
	_, err := client.Login("test@nomadreact.com", "xxx")
	assert.NotNil(t, err)
}

func TestLogin(t *testing.T) {
	client = movielens.NewClient(nil)
	_, err := client.Login("test@nomadreact.com", "ddpgvoQuTCCs4VouVQedneuU")
	assert.Nil(t, err)
}
