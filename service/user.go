package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"userservice/data"

	"github.com/Smart-Pot/pkg/adapter/amqp"
	"github.com/Smart-Pot/pkg/tool/crypto"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	logger   log.Logger
	producer amqp.Producer
}

type Service interface {
	GetUsersPublic(ctx context.Context, userIDList []string) (result []*data.UserPublicData, err error)
	Get(ctx context.Context, userID string) (*data.UserPublicData, error)
	Create(ctx context.Context, newUser data.User) error
	Update(ctx context.Context, userID string, updatedUser data.User) error
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

func (s service) Create(ctx context.Context, newUser data.User) error {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "Create",
			"param:newComment", newUser,
			"took", time.Since(beginTime))
	}(time.Now())

	if err := data.CreateUser(ctx, newUser); err != nil {
		return err
	}

	// Hash user id for verification mail
	h, err := crypto.Encrypt(newUser.ID)
	if err != nil {
		return err
	}

	r := struct {
		Hash  string `json:"hash"`
		Email string `json:"email"`
	}{
		Hash:  h,
		Email: newUser.Email,
	}

	b, err := json.Marshal(r)

	fmt.Println("SENDING", string(b))
	if err = s.producer.Produce(b); err != nil {
		return err
	}

	return nil
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
