package endpoints

import (
	"context"
	"userservice/service"

	"github.com/go-kit/kit/endpoint"
)


func makeGetUsersPublicEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserPublicDataRequest)
		result, err := s.GetUsersPublic(ctx, req.IDList)
		response := UserPublicDataResponse{Users: result, Success: 1, Message: "Public data found!"}
		if err != nil {
			return nil,err
		}
		return response, nil
	}
}

func makeGetEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserRequest)
		result, err := s.Get(ctx, req.ID)
		response := UserResponse{User: result, Success: 1, Message: "User found!"}
		if err != nil {
			return nil,err
		}
		return response, nil
	}
}



func makeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		req.UpdatedUser.ID = req.ID
		err := s.Update(ctx,  req.UpdatedUser)
		response := UserResponse{User: nil, Success: 1, Message: "User updated!"}
		if err != nil {
			return nil,err
		}
		return response, nil
	}
}
