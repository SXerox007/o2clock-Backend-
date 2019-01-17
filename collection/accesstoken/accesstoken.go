package accesstoken

import (
	"context"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"o2clock/constants/collections"
	"o2clock/constants/errormsg"
	"o2clock/core/jwtauth"
	"o2clock/db/mongodb"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mongodb/mongo-go-driver/bson"
	objectid "github.com/mongodb/mongo-go-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	VERSION = "v1.0"
)

type AccessToken struct {
	ID          objectid.ObjectID `bson:"_id,omitempty"`
	UserId      objectid.ObjectID `bson:"user_id,omitempty"`
	AccessToken string            `bson:"access_token,omitempty"`
	Version     string            `bson:"version"`
	CaptureTime time.Time         `bson:"capture_time"`
}

func CreateUserAccessToken(userid objectid.ObjectID, username string) (string, error) {
	access_token, err := jwtauth.GenerateToken(createKey(), username, userid.String())
	if err == nil {
		data := AccessToken{
			UserId:      userid,
			AccessToken: access_token,
			Version:     VERSION,
			CaptureTime: time.Now(),
		}
		res, err := mongodb.CreateCollection(collections.COLLECTIONS_ACCESS_TOKEN).InsertOne(context.Background(), data)
		if err != nil {
			return "", status.Errorf(
				codes.Internal,
				fmt.Sprintln(errormsg.ERR_MSG_INTERNAL, err))
		}
		_, ok := res.InsertedID.(objectid.ObjectID)
		if !ok {
			return "", status.Errorf(
				codes.Internal,
				fmt.Sprintln(errormsg.ERR_INTERNAL_OID, ok))
		}
		return access_token, nil
	}
	return "", err
}

func createKey() *rsa.PrivateKey {
	privateKey, err := ioutil.ReadFile("secure-keys/o2clock.rsa")
	if err != nil {
		return nil
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil
	}
	return key
}

func CheckAccessToken(accessToken string) error {
	filter := bson.M{collections.PARAM_ACCESS_TOKEN: accessToken}
	err := mongodb.CreateCollection(collections.COLLECTIONS_ACCESS_TOKEN).FindOne(context.Background(), filter).Decode(&AccessToken{})
	if err != nil {
		return status.Errorf(
			codes.PermissionDenied,
			fmt.Sprintln(errormsg.ERR_MSG_INVALID_ACCESS_TOKEN))
	}
	return nil
}

/**
*
* Get all the data through access token
*
**/
func GetAllAccessTokenInfo(accessToken string) (*AccessToken, error) {
	data := &AccessToken{}
	filter := bson.M{collections.PARAM_ACCESS_TOKEN: accessToken}
	res := mongodb.CreateCollection(collections.COLLECTIONS_ACCESS_TOKEN).FindOne(context.Background(), filter)
	//for single data decode
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.Aborted,
			fmt.Sprintln(errormsg.ERR_MSG_DATA_CANT_DECODE, err))
	}
	return data, nil
}

func UpdateAccessToken(userid objectid.ObjectID, username string) (string, error) {
	data := &AccessToken{}
	filter := bson.M{collections.PARAM_USER_ID: userid}
	res := mongodb.CreateCollection(collections.COLLECTIONS_ACCESS_TOKEN).FindOne(context.Background(), filter)
	//for single data decode
	if err := res.Decode(data); err != nil {
		return "", status.Errorf(
			codes.Aborted,
			fmt.Sprintln(errormsg.ERR_MSG_DATA_CANT_DECODE, err))
	}
	access_token, err := jwtauth.GenerateToken(createKey(), username, userid.String())
	//update host name on particular sequence
	filter = bson.M{collections.PARAM_ID: data.ID}
	data.AccessToken = access_token

	_, err = mongodb.CreateCollection(collections.COLLECTIONS_ACCESS_TOKEN).ReplaceOne(context.Background(), filter, data)

	if err != nil {
		return "", status.Errorf(
			codes.Aborted,
			fmt.Sprintln(errormsg.ERR_MSG_INTERNAL_SERVER, err))
	}
	return access_token, nil
}
