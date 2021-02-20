package service

import (
	"context"
	"time"
	"userservice/data"

	"github.com/Smart-Pot/pkg/adapter/amqp"
	"github.com/Smart-Pot/pkg/common/perrors"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)


var (
	// ErrNotAuthorized codes returned when user not authorized for a task
	ErrNotAuthorized = perrors.New("User is not authorized",401)
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
	if err != nil {
		return nil,perrors.ErrInternalServer
	}
	return result, err
}

func (s service) Get(ctx context.Context, userID string) (result *data.UserPublicData, err error) {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "Get",
			"param:userID", userID,
			"result", result,
			"err", err,
			"took", time.Since(beginTime))
	}(time.Now())
	result, err = data.GetUserByID(ctx, userID)
	if err != nil {
		return nil, perrors.ErrInternalServer
	}
	return result, err
}


func (s service) Update(ctx context.Context,  updatedUser data.User) error {
	var err error
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "Update",
			"param:updatedUser", updatedUser,
			"result:err", err,
			"took", time.Since(beginTime))
	}(time.Now())
	err = data.UpdateUser(ctx, updatedUser)
	if err != nil {
		return perrors.ErrInternalServer
	}
	return nil
}

