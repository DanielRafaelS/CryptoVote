package model_test

import (
	"cryptovote/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCryptoVote_IsValid_Should_Works(t *testing.T) {
	m := model.CryptoVote{
		Name:     "KLV",
		IsUpvote: true,
	}

	err := m.IsValid()
	require.NoError(t, err)
}

func TestCryptoVote_IsInvalid_Should_Not_Works(t *testing.T) {
	m := model.CryptoVote{
		Name:     "",
		IsUpvote: true,
	}

	err := m.IsValid()
	require.Error(t, err)
}
