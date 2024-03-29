package controller

import (
	"github.com/takez0o/honestwork-api/api/repository"
)

type RatingController struct {
	Address string
}

type RatingIndexer struct {
	IndexName string
}

func NewRatingController(address string) *RatingController {
	return &RatingController{
		Address: address,
	}
}

func NewRatingIndexer(index_name string) *RatingIndexer {
	return &RatingIndexer{
		IndexName: index_name,
	}
}

func (s *RatingController) GetRating() (string, error) {
	rating, err := repository.StringRead("rating:" + s.Address)
	if err != nil {
		return "", err
	}
	return rating, nil
}
