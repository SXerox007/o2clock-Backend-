package register

import (
	"context"
	"net/http"
	"o2clock/api-proto/onboarding/register"
	mdb "o2clock/collection/allusers"
	"o2clock/constants/appconstant"
	dbsettings "o2clock/settings/db"
	pdb "o2clock/table/allusers"

	"google.golang.org/grpc"
)

type Server struct {
}

func RegisterUserService(srv *grpc.Server) {
	regsiterpb.RegisterRegisterServiceServer(srv, &Server{})
}

func (*Server) RegisterUserService(ctx context.Context, req *regsiterpb.RegisterUserRequest) (*regsiterpb.RegisterUserResponse, error) {
	var err error
	var token string
	if dbsettings.IsEnableMongoDb() {
		token, err = mdb.CreateUser(req)
	}
	if dbsettings.IsEnablePostgres() {
		token, err = pdb.CreateUser(req)
	}
	if err == nil {
		//success
		return &regsiterpb.RegisterUserResponse{
			Message:     appconstant.MSG_SUCCESS,
			Code:        http.StatusOK,
			AccessToken: token,
		}, nil
	}

	// return &regsiterpb.RegisterUserResponse{
	// 	Message:     appconstant.MSG_FAILURE,
	// 	Code:        http.StatusNoContent,
	// 	AccessToken: "",
	// }, err
	return nil, err
}
