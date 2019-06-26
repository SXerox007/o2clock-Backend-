package git

import (
	"fmt"
	"o2clock/api-proto/webhooks/git"
	"o2clock/constants/errormsg"
	db "o2clock/db/postgres"
	"o2clock/utils/log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	VERSION                            = "v1.0"
	SQL_STATEMENT_INSERT_GITHUB_RECORD = `
	Insert into all_github_track (commit_id,message,url,type,version)
	values($1,$2,$3,$4,$5)`
)

type Github struct {
	TableName   struct{}  `sql:"all_github_track" json:"-"`
	CommitId    string    `sql:"commit_id" param:"commit_id"`
	Message     string    `sql:"message" param:"message"`
	Url         string    `sql:"url" param:"url"`
	Type        string    `sql:"type" param:"type"`
	Version     string    `sql:"version" param:"version"`
	CaptureTime time.Time `sql:"capture_time" param:"capture_time"`
}

func SaveGithubPushWebhookInfo(req *githubpb.GithubPushWebhookRequest) error {
	sqlStatement := SQL_STATEMENT_INSERT_GITHUB_RECORD
	for i := 0; i < len(req.Commits); i++ {
		err := db.GetClient().QueryRow(sqlStatement, req.Commits[i].GetId, req.Commits[i].GetMessage(), req.Commits[i].GetUrl(), "push", VERSION)
		if err != nil {
			log.Error.Println(errormsg.ERR_MSG_INTERNAL_SERVER, err)
			return status.Errorf(
				codes.Internal,
				fmt.Sprintln(errormsg.ERR_MSG_INTERNAL_SERVER, err))
		}
	}
	return nil
}
