package repository

import (
	"context"
	"cryptovote/database"
	"cryptovote/model"
	"errors"
)

var RepositoryVote crytoVoteRepository

type crytoVoteRepository struct{}

func (r crytoVoteRepository) List(ctx context.Context) ([]model.CryptoVote, error) {
	var list []model.CryptoVote

	rows, err := database.Db.QueryContext(ctx, "SELECT name, is_upvote FROM votes")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var m model.CryptoVote

		if err := rows.Scan(&m.Name, &m.IsUpvote); err != nil {
			return nil, err
		}

		list = append(list, m)
	}

	return list, nil
}

func (r crytoVoteRepository) Insert(vote *model.CryptoVote) error {
	result, err := database.Db.Exec("INSERT INTO votes (name, is_upvote) VALUES ($1, $2)", vote.Name, vote.IsUpvote)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows < 1 {
		return errors.New("failed to insert value on table")
	}

	return nil
}
