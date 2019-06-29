package forgotpassword

import (
	"context"
	"net/http"
	"o2clock/api-proto/onboarding/forgotpassword"
	"o2clock/constants/appconstant"

	mdb "o2clock/collection/forgotpassword"
	dbsettings "o2clock/settings/db"
	pdb "o2clock/table/forgotpassword"

	"google.golang.org/grpc"
)

type Server struct {
}

func RegisterForgotPasswordService(srv *grpc.Server) {
	forgotpasswordpb.RegisterForgotPasswordServiceServer(srv, &Server{})
}

func (*Server) ForgotPassowrdUserService(ctx context.Context, req *forgotpasswordpb.ForgotPasswordRequest) (*forgotpasswordpb.ForgotPasswordResponse, error) {
	var err error
	if dbsettings.IsEnableMongoDb() {
		err = mdb.UserForgotPassword(req.Email)
	}
	if dbsettings.IsEnablePostgres() {
		err = pdb.UserForgotPassword(req.Email)
	}
	if err == nil {
		//success
		return &forgotpasswordpb.ForgotPasswordResponse{
			Message: appconstant.MSG_SUCCESS,
			Code:    http.StatusOK,
		}, nil
	}
	return nil, err
}
