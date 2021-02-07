package endpoints

import "github.com/Smart-Pot/pkg/adapter/amqp"

func MakeNewUserConsumer() (amqp.Consumer, error) {
	return amqp.MakeConsumer("newUser1", "NewUser")
}
