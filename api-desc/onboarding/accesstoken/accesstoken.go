package accesstoken

import (
	"context"
	"net/http"
	"o2clock/api-proto/onboarding/accesstoken"
	mdb "o2clock/collection/accesstoken"
	"o2clock/constants/appconstant"
	dbsettings "o2clock/settings/db"
	pdb "o2clock/table/accesstoken"

	"google.golang.org/grpc"
)

type Server struct {
}

func RegisterAccessTokenService(srv *grpc.Server) {
	accesstokenpb.RegisterAccessTokenServiceServer(srv, &Server{})
}

func (*Server) CheckAccessTokenService(ctx context.Context, req *accesstokenpb.AccessTokenRequest) (*accesstokenpb.AccessTokenResponse, error) {
	var err error
	if dbsettings.IsEnableMongoDb() {
		err = mdb.CheckAccessToken(req.GetAccessToken())
	}
	if dbsettings.IsEnablePostgres() {
		err = pdb.CheckAccessToken(req.GetAccessToken())

	}
	if err == nil {
		//success
		return &accesstokenpb.AccessTokenResponse{
			Message: appconstant.MSG_SUCCESS,
			Code:    http.StatusOK,
		}, nil
	}
	return nil, err
}
