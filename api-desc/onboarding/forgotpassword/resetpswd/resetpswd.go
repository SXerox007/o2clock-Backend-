package resetpswd

import (
	"context"
	"net/http"
	"o2clock/api-proto/onboarding/forgotpassword/resetpswd"
	"o2clock/constants/appconstant"

	mdb "o2clock/collection/forgotpassword/resetpswd"
	dbsettings "o2clock/settings/db"
	pdb "o2clock/table/forgotpassword/resetpswd"

	"google.golang.org/grpc"
)

type Server struct {
}

func RegisterResetPasswordService(srv *grpc.Server) {
	resetpswdpb.RegisterResetPasswordServiceServer(srv, &Server{})
}

func (*Server) ResetPassowrdUserService(ctx context.Context, req *resetpswdpb.ResetPasswordRequest) (*resetpswdpb.ResetPasswordResponse, error) {
	var err error
	if dbsettings.IsEnableMongoDb() {
		err = mdb.UserResetPassword(req)
	}
	if dbsettings.IsEnablePostgres() {
		err = pdb.UserResetPassword(req)
	}
	if err == nil {
		//success
		return &resetpswdpb.ResetPasswordResponse{
			Message: appconstant.MSG_SUCCESS,
			Code:    http.StatusOK,
		}, nil
	}
	return nil, err
}
