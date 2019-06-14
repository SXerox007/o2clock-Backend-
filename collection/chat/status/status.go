package status

import (
	"context"
	"fmt"
	"o2clock/constants/collections"
	"o2clock/constants/errormsg"
	"o2clock/db/mongodb"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	objectid "github.com/mongodb/mongo-go-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	errorStatus "google.golang.org/grpc/status"
)

const (
	VERSION = "v1.0"
)

type AllPersonStatus struct {
	ID          objectid.ObjectID `bson:"_id,omitempty"`
	PersonID    string            `bson:"person_id"`
	Status      int               `bson:"status"`
	Version     string            `bson:"version"`
	CaptureTime time.Time         `bson:"capture_time"`
}

/**
*
* Set the  person to online/offline
*
**/
func SetUsersStatus(status int, userId string) error {
	if ValidateAndUpdateUsersStatus(userId, status) == nil {
		return nil
	}
	data := AllPersonStatus{
		PersonID:    userId,
		Status:      status,
		Version:     VERSION,
		CaptureTime: time.Now(),
	}
	res, err := mongodb.CreateCollection(collections.COLLECTIONS_ALL_USERS_STATUS).InsertOne(context.Background(), data)
	if err != nil {
		return errorStatus.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_MSG_INTERNAL, err))
	}

	_, ok := res.InsertedID.(objectid.ObjectID)
	if !ok {
		return errorStatus.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_INTERNAL_OID, ok))
	}
	return nil
}

/**
*
* validate and update status
*
**/
func ValidateAndUpdateUsersStatus(personId string, status int) error {
	data := &AllPersonStatus{}
	filter := bson.M{collections.PARAM_PERSON_ID: personId}
	res := mongodb.CreateCollection(collections.COLLECTIONS_ALL_USERS_STATUS).FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return err
	} else {
		data.Status = status
		_, err := mongodb.CreateCollection(collections.COLLECTIONS_ALL_USERS_STATUS).ReplaceOne(context.Background(), filter, data)
		if err != nil {
			return errorStatus.Errorf(
				codes.FailedPrecondition,
				fmt.Sprintln(errormsg.ERR_STATUS, err))
		} else {
			return nil
		}
	}
}
