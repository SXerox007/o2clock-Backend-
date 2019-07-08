package login

import (
	"context"
	"fmt"
	"o2clock/api-proto/onboarding/login"
	"o2clock/collection/accesstoken"
	"o2clock/collection/allusers"
	"o2clock/constants/collections"
	"o2clock/constants/errormsg"
	"o2clock/db/mongodb"
	"o2clock/utils/pswdmanager"

	objectid "github.com/mongodb/mongo-go-driver/bson/primitive"

	"github.com/mongodb/mongo-go-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LoginUser(req *loginpb.LoginRequest) (string, error) {
	id, username, err := ValidationUserExist(req)
	if err != nil {
		return "", err
	}
	return accesstoken.UpdateAccessToken(id, username)
}

func ValidationUserExist(req *loginpb.LoginRequest) (objectid.ObjectID, string, error) {
	data := &allusers.Users{}
	filter := bson.M{collections.PARAM_USER_NAME: req.GetUsernameEmail()}
	res := mongodb.CreateCollection(collections.COLLECTIONS_ALL_USERS).FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		filter := bson.M{collections.PARAM_EMAIL: req.GetUsernameEmail()}
		res := mongodb.CreateCollection(collections.COLLECTIONS_ALL_USERS).FindOne(context.Background(), filter)
		if err := res.Decode(data); err != nil {
			return data.ID, "", status.Errorf(
				codes.Aborted,
				fmt.Sprintln(errormsg.ERR_MSG_INVALID_CREDS))
		}
		if pswdmanager.VerifyPassword(data.Password, []byte(req.GetPassword())) {
			return data.ID, data.UserName, nil
		} else {
			return data.ID, "", status.Errorf(
				codes.Internal,
				fmt.Sprintln(errormsg.ERR_MSG_PSWD_DECODE))
		}
	}
	if pswdmanager.VerifyPassword(data.Password, []byte(req.GetPassword())) {
		return data.ID, data.UserName, nil
	} else {
		return data.ID, "", status.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_MSG_PSWD_DECODE))
	}
}
