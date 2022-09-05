package model

import "errors"

var cryptoRequireds = []string{
	"KLV",
	"BITCOIN",
	"ETHEREUM",
	"SOLANA",
}

type CryptoVote struct {
	Name     string `json:"name"`
	IsUpvote bool   `json:"isUpvote"`
}

func NewVote(name string, isUpvote bool) *CryptoVote {
	return &CryptoVote{
		name, isUpvote,
	}
}

func (v *CryptoVote) IsValid() error {
	if v.Name == "" {
		return errors.New("crypto name is required")
	}

	if len(v.Name) > 50 {
		return errors.New("crypto name is invalid")
	}

	isValid := false

	for _, validCoin := range cryptoRequireds {
		if validCoin == v.Name {
			isValid = true
		}
	}

	if !isValid {
		return errors.New("crypto name is not in available range")
	}

	return nil
}
