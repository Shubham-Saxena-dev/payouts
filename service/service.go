package service

/*
This is a service class. Controller sends the request to this class where it gets validated and sent to store into the database
*/

import (
	"gopkg.in/validator.v2"
	"takeHomeTest/errorHandlers"
	"takeHomeTest/models"
	"takeHomeTest/repository"
	"takeHomeTest/utility"
)

type Service interface {
	CreatePayout(item []models.Item) ([]models.Payout, error)
}

type service struct {
	repo         repository.Repository
	errorHandler errorHandlers.ErrorHandlers
}

func NewService(repo repository.Repository, errorHandler errorHandlers.ErrorHandlers) Service {
	return &service{repo: repo,
		errorHandler: errorHandler,
	}
}

/*
This method validates the request body and sends the request to create payout
and store to mysql database
It returns the created payouts along with the error, if any
*/

func (s *service) CreatePayout(item []models.Item) ([]models.Payout, error) {

	if err := validator.Validate(item); err != nil {
		return nil, err
	}

	payouts, err := utility.CreatePayout(item)
	if err != nil {
		s.errorHandler.FailOnError(err, "Failed to create payouts")
	}
	err = s.repo.StorePayouts(payouts)
	if err != nil {
		s.errorHandler.FailOnError(err, "Failed to store payouts in the database")
	}
	return payouts, err
}
