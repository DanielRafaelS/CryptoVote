package repository_test

import (
	"context"
	"cryptovote/database"
	"cryptovote/database/repository"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	database.StartDatabase("host=localhost port=5431 user=postgres password=postgres dbname=cryptovote sslmode=disable")
	os.Exit(m.Run())
}

func TestCrytoVoteRepository_List(t *testing.T) {
	result, err := repository.RepositoryVote.List(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
}
