package logout

import (
	"context"
	"net/http"
	"o2clock/api-proto/home/logout"
	"o2clock/constants/appconstant"

	mdb "o2clock/collection/logout"
	dbsettings "o2clock/settings/db"
	pdb "o2clock/table/logout"

	"google.golang.org/grpc"
)

type Server struct {
}

func RegisterLogoutService(srv *grpc.Server) {
	logoutpb.RegisterLogoutServiceServer(srv, &Server{})
}

func (*Server) LogoutUserService(ctx context.Context, req *logoutpb.LogoutRequest) (*logoutpb.LogoutResponse, error) {
	var err error
	if dbsettings.IsEnableMongoDb() {
		err = mdb.LogoutUser(req)
	}
	if dbsettings.IsEnablePostgres() {
		err = pdb.LogoutUser(req)
	}
	if err == nil {
		//success
		return &logoutpb.LogoutResponse{
			Message: appconstant.MSG_SUCCESS,
			Code:    http.StatusOK,
		}, nil
	}
	return nil, err
}
