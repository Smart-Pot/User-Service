package service

import (
	"context"
	"errors"
	"time"
	"userservice/data"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	logger log.Logger
}

type Service interface {
	GetUsersPublic(ctx context.Context, userIDList []string) (result []*data.UserPublicData, err error)
	Get(ctx context.Context, userID string) (*data.UserPublicData, error)
	Create(ctx context.Context, newUser data.User) error
	Update(ctx context.Context, userID string, updatedUser data.User) error
}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
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

func (s service) Create(ctx context.Context, newUser data.User) error {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "Create",
			"param:newComment", newUser,
			"took", time.Since(beginTime))
	}(time.Now())
	return data.CreateUser(ctx, newUser)
}

func (s service) Update(ctx context.Context, userID string, updatedUser data.User) error {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "Update",
			"param:userID", userID,
			"param:updatedUser", updatedUser,
			"took", time.Since(beginTime))
	}(time.Now())
	if updatedUser.ID != userID {
		return errors.New("user cannot update for other users")
	}
	return data.UpdateUser(ctx, updatedUser)
}
