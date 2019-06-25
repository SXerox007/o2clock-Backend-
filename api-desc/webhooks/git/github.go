package git

import (
	"context"
	"net/http"
	"o2clock/api-proto/webhooks/git"
	mdb "o2clock/collection/webhooks/git"
	"o2clock/constants/appconstant"
	dbsettings "o2clock/settings/db"
	pdb "o2clock/table/webhooks/git"

	"google.golang.org/grpc"
)

type Server struct {
}

func RegisterGithubService(srv *grpc.Server) {
	githubpb.RegisterGithubWebhookServicesServer(srv, &Server{})
}

func (*Server) FetchGithubPushCommitInfo(ctx context.Context, req *githubpb.GithubPushWebhookRequest) (*githubpb.CommonResponse, error) {
	var err error
	if dbsettings.IsEnableMongoDb() {
		err = mdb.SaveGithubPushWebhookInfo(req)
	}

	if dbsettings.IsEnablePostgres() {
		err = pdb.SaveGithubPushWebhookInfo(req)
	}

	if err == nil {
		return &githubpb.CommonResponse{
			Message: appconstant.MSG_SUCCESS,
			Code:    http.StatusOK,
		}, nil

	}
	return nil, err
}
