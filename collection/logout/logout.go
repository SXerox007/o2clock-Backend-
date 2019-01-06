package logout

import (
	"context"
	"fmt"
	"o2clock/api-proto/home/logout"
	"o2clock/collection/accesstoken"
	"o2clock/constants/collections"
	"o2clock/constants/errormsg"
	"o2clock/db/mongodb"

	"github.com/mongodb/mongo-go-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LogoutUser(req *logoutpb.LogoutRequest) error {
	data := &accesstoken.AccessToken{}
	filter := bson.M{collections.PARAM_ACCESS_TOKEN: req.GetAccessToken()}
	res := mongodb.CreateCollection(collections.COLLECTIONS_ACCESS_TOKEN).FindOne(context.Background(), filter)
	//for single data decode
	if err := res.Decode(data); err != nil {
		return status.Errorf(
			codes.Aborted,
			fmt.Sprintln(errormsg.ERR_MSG_DATA_CANT_DECODE, err))
	}
	//update host name on particular sequence
	filter = bson.M{collections.PARAM_ID: data.ID}
	data.AccessToken = ""

	_, err := mongodb.CreateCollection(collections.COLLECTIONS_ACCESS_TOKEN).ReplaceOne(context.Background(), filter, data)

	if err != nil {
		return status.Errorf(
			codes.Aborted,
			fmt.Sprintln(errormsg.ERR_MSG_INTERNAL_SERVER, err))
	}
	return nil
}
