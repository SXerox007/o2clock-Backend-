package login

import (
	"context"
	"net/http"
	"o2clock/api-proto/onboarding/login"
	"o2clock/constants/appconstant"

	mdb "o2clock/collection/login"
	dbsettings "o2clock/settings/db"
	pdb "o2clock/table/login"

	"google.golang.org/grpc"
)

type Server struct {
}

func RegisterLoginUserService(srv *grpc.Server) {
	loginpb.RegisterLoginServiceServer(srv, &Server{})
}

func (*Server) LoginUserService(ctx context.Context, req *loginpb.LoginRequest) (*loginpb.LoginResponse, error) {
	var err error
	var token string
	if dbsettings.IsEnableMongoDb() {
		token, err = mdb.LoginUser(req)
	}
	if dbsettings.IsEnablePostgres() {
		token, err = pdb.LoginUser(req)
	}
	if err == nil {
		//success
		return &loginpb.LoginResponse{
			Message:     appconstant.MSG_SUCCESS,
			Code:        http.StatusOK,
			AccessToken: token,
		}, nil
	}
	return nil, err
}
