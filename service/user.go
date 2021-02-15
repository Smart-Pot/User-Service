package service

import (
	"context"
	"errors"
	"time"
	"userservice/data"

	"github.com/Smart-Pot/pkg/adapter/amqp"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)


var (
	// ErrNotAuthorized codes returned when user not authorized for a task
	ErrNotAuthorized = errors.New("User is not authorized")
	
)

type service struct {
	logger   log.Logger
	producer amqp.Producer
}

type Service interface {
	GetUsersPublic(ctx context.Context, userIDList []string) (result []*data.UserPublicData, err error)
	Get(ctx context.Context, userID string) (*data.UserPublicData, error)
	Update(ctx context.Context,  updatedUser data.User) error
}

func NewService(logger log.Logger, p amqp.Producer) Service {
	return &service{
		logger:   logger,
		producer: p,
	}
}

func (s service) GetUsersPublic(ctx context.Context, userIDList []string) (result []*data.UserPublicData, err error) {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "GetUsersPublicData",
			"param:userIDList", userIDList,
			"result", result,
			"took", time.Since(beginTime))
	}(time.Now())
	result, err = data.GetUsersPublicData(ctx, userIDList)
	return result, err
}

func (s service) Get(ctx context.Context, userID string) (result *data.UserPublicData, err error) {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "Get",
			"param:userID", userID,
			"result", result,
			"took", time.Since(beginTime))
	}(time.Now())
	result, err = data.GetUserByID(ctx, userID)
	return result, err
}


func (s service) Update(ctx context.Context,  updatedUser data.User) error {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "Update",
			"param:updatedUser", updatedUser,
			"took", time.Since(beginTime))
	}(time.Now())
	return data.UpdateUser(ctx, updatedUser)
}

