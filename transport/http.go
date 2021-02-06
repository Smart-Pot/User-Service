package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"userservice/data"
	"userservice/endpoints"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

const userIDTag = "x-user-id"

func MakeHTTPHandlers(e endpoints.Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter().PathPrefix("/user").Subrouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path("/single/{id}").Handler(httptransport.NewServer(
		e.Get,
		decodeUserHTTPRequest,
		encodeHTTPResponse,
		options...,
	))

	r.Methods("GET").Path("/public").Handler(httptransport.NewServer(
		e.GetUsersPublic,
		decodeUserPublicDataHTTPRequest,
		encodeHTTPResponse,
		options...,
	))

	r.Methods("PUT").Path("/update").Handler(httptransport.NewServer(
		e.Update,
		decodeUpdateUserHTTPRequest,
		encodeHTTPResponse,
		options...,
	))

	r.Methods("POST").Path("/create").Handler(httptransport.NewServer(
		e.Create,
		decodeNewUserHTTPRequest,
		encodeHTTPResponse,
		options...,
	))
	return r
}

func encodeHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeUserHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {

	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		// Handler error
	}
	return endpoints.UserRequest{
		ID: id,
	}, nil
}
func decodeUserPublicDataHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.UserPublicDataRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func decodeUpdateUserHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var updatedUser data.User

	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		return nil, err
	}

	return endpoints.UpdateUserRequest{
		ID:          r.Header.Get(userIDTag),
		UpdatedUser: updatedUser,
	}, nil
}

func decodeNewUserHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var newUser data.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		return nil, err
	}
	return endpoints.NewUserRequest{
		NewUser: newUser,
	}, nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
