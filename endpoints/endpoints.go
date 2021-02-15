package endpoints

import (
	"userservice/data"
	"userservice/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetUsersPublic endpoint.Endpoint
	Get            endpoint.Endpoint
	Update         endpoint.Endpoint
}

type UserPublicDataResponse struct {
	Users   []*data.UserPublicData
	Success int32
	Message string
}

type UserResponse struct {
	User    *data.UserPublicData
	Success int32
	Message string
}

type UserRequest struct {
	ID string
}

type UpdateUserRequest struct {
	ID          string
	UpdatedUser data.User
}

type UserPublicDataRequest struct {
	IDList []string `json:"idList"`
}


func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		GetUsersPublic: makeGetUsersPublicEndpoint(s),
		Get:            makeGetEndpoint(s),
		Update:         makeUpdateEndpoint(s),
	}
}
