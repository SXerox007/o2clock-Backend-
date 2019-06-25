package git

import (
	"context"
	"fmt"
	"o2clock/api-proto/webhooks/git"
	"o2clock/constants/collections"
	"o2clock/constants/errormsg"
	"o2clock/db/mongodb"
	"time"

	objectid "github.com/mongodb/mongo-go-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	VERSION = "v1.0"
)

type Github struct {
	ID        objectid.ObjectID `bson:"_id,omitempty"`
	CommitId  string            `bson:"commit_id"`
	Message   string            `bson:"message"`
	Url       string            `bson:"url"`
	Type      string            `bson:"type"`
	Version   string            `bson:"version"`
	TimeStamp time.Time         `bson:"timestamp"`
}

func SaveGithubPushWebhookInfo(req *githubpb.GithubPushWebhookRequest) error {
	var data Github
	for i := 0; i < len(req.Commits); i++ {
		data.Version = VERSION
		data.CommitId = req.Commits[i].GetId()
		data.Message = req.Commits[i].GetMessage()
		data.Url = req.Commits[i].GetUrl()
		data.Type = "push"
		data.TimeStamp = time.Now()
	}
	res, err := mongodb.CreateCollection(collections.COLLECTIONS_ALL_GITHUB).InsertOne(context.Background(), data)
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_MSG_INTERNAL, err))
	}

	_, ok := res.InsertedID.(objectid.ObjectID)
	if !ok {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_INTERNAL_OID, ok))
	}

	return nil
}
