package transport

import (
	"context"
	"encoding/json"
	"userservice/data"
	"userservice/service"

	"github.com/Smart-Pot/pkg/adapter/amqp"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func MakeNewUserConsumerTask(service service.Service, consumer amqp.Consumer, log log.Logger) func() {
	return func() {
		for {
			newUserJSON := consumer.Consume()
			var newUser data.User
			json.Unmarshal(newUserJSON, &newUser)
			if err := service.Create(context.TODO(), newUser); err != nil {
				level.Error(log).Log(err)
			}
		}
	}
}
