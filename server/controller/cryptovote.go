package controller

import (
	"cryptovote/database/repository"
	"cryptovote/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

var CryptoController cryptoVoteController

type cryptoVoteController struct{}

func (cv cryptoVoteController) GetVotes(c *fiber.Ctx) error {
	result, err := repository.RepositoryVote.List(c.Context())
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(err)
	}

	if result == nil {
		c.Status(http.StatusNoContent)
	}

	return c.Status(http.StatusOK).JSON(model.ApiResponse{
		Code: http.StatusOK,
		Data: result,
	})
}
func (cv cryptoVoteController) SetVote(c *fiber.Ctx) error {
	var req struct {
		Name     string `json:"name"`
		IsUpvote bool   `json:"isUpvote"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	vote := model.NewVote(req.Name, req.IsUpvote)

	if err := vote.IsValid(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fmt.Println(fmt.Sprintf("vote for %s in %t has been registed ", req.Name, req.IsUpvote))

	if err := repository.RepositoryVote.Insert(vote); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON("vote insered sucessfully")
}
