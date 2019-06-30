package allusers

import (
	"context"
	"fmt"
	"log"
	"o2clock/api-proto/home/chat"
	"o2clock/api-proto/onboarding/register"
	"o2clock/collection/accesstoken"
	"o2clock/constants/collections"
	"o2clock/constants/errormsg"
	"o2clock/db/mongodb"
	"o2clock/utils/pswdmanager"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	objectid "github.com/mongodb/mongo-go-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	VERSION = "v1.0"
)

type Users struct {
	ID          objectid.ObjectID `bson:"_id,omitempty"`
	Phone       string            `bson:"phone"`
	FirstName   string            `bson:"first_name"`
	LastName    string            `bson:"last_name"`
	CompanyName string            `bson:"company_name"`
	CountryCode string            `bson:"country_code"`
	UserName    string            `bson:"user_name"`
	Password    string            `bson:"password"`
	Email       string            `bson:"email"`
	Lat         float64           `bson:"lat"`
	Lan         float64           `bson:"lan"`
	Version     string            `bson:"version"`
	CaptureTime time.Time         `bson:"capture_time"`
}

/**
*
* Create the new User
*
**/
func CreateUser(req *regsiterpb.RegisterUserRequest) (string, error) {
	if err := validationCheck(req); err != nil {
		return "", err
	}
	data := Users{
		Phone:       req.GetPhone(),
		FirstName:   req.GetFirstName(),
		LastName:    req.GetLastName(),
		CompanyName: req.GetCompanyName(),
		CountryCode: req.GetCountryCode(),
		UserName:    req.GetUserName(),
		Password:    pswdmanager.GetPswdHash([]byte(req.GetPassword())),
		Email:       req.GetEmail(),
		Lat:         req.GetLocation().GetLat(),
		Lan:         req.GetLocation().GetLan(),
		Version:     VERSION,
		CaptureTime: time.Now(),
	}
	res, err := mongodb.CreateCollection(collections.COLLECTIONS_ALL_USERS).InsertOne(context.Background(), data)
	if err != nil {
		return "", status.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_MSG_INTERNAL, err))
	}

	oid, ok := res.InsertedID.(objectid.ObjectID)
	if !ok {
		return "", status.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_INTERNAL_OID, ok))
	}
	return accesstoken.CreateUserAccessToken(oid, req.GetUserName())
}

/**
*
* Get all the users
*
**/
func GetAllUsers(req *chatpb.CommonRequest) ([]*chatpb.User, error) {
	if err := accesstoken.CheckAccessToken(req.GetAccessToken()); err == nil {
		res, err := mongodb.CreateCollection(collections.COLLECTIONS_ALL_USERS).Find(context.Background(), nil)
		if err != nil {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintln(errormsg.ERR_MSG_INTERNAL, err))
		} else {
			var items []*chatpb.User
			for res.Next(nil) {
				item := Users{}
				if err := res.Decode(&item); err != nil {
					return nil, status.Errorf(
						codes.Aborted,
						fmt.Sprintln(errormsg.ERR_MSG_DATA_CANT_DECODE, err))
				}
				//items = append(items, &item)
				items = append(items, &chatpb.User{
					UserId:   item.ID.String(),
					UserName: item.FirstName + " " + item.LastName + "\n" + item.Email,
				})
			}
			return items, nil

		}
	} else {
		return nil, err
	}
}

/**
*
* Get login user info
*
**/
func GetUserInfo(req *chatpb.CommonRequest) (*chatpb.User, error) {
	data, err := accesstoken.GetAllAccessTokenInfo(req.GetAccessToken())
	if err != nil {
		return nil, err
	} else {
		user, _ := GetUserInfoById(data.UserId)
		return &chatpb.User{
			UserId:   user.ID.String(),
			UserName: user.FirstName + " " + user.LastName,
		}, nil
	}
}

/**
*
* Get the user info by id
*
**/
func GetUserInfoById(userId objectid.ObjectID) (*Users, error) {
	data := &Users{}
	filter := bson.M{collections.PARAM_ID: userId}
	res := mongodb.CreateCollection(collections.COLLECTIONS_ALL_USERS).FindOne(context.Background(), filter)
	//for single data decode
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.Aborted,
			fmt.Sprintln(errormsg.ERR_MSG_DATA_CANT_DECODE, err))
	}
	return data, nil
}

func validationCheck(req *regsiterpb.RegisterUserRequest) error {
	filter := bson.M{collections.PARAM_USER_NAME: req.GetUserName()}
	err := mongodb.CreateCollection(collections.COLLECTIONS_ALL_USERS).FindOne(context.Background(), filter).Decode(&Users{})
	log.Println("Error", err)
	if err == nil {
		return status.Errorf(
			codes.Aborted,
			fmt.Sprintln(errormsg.ERR_MSG_USERNAME_ALREADY_EXIST))
	}

	filter = bson.M{collections.PARAM_EMAIL: req.GetEmail()}
	err = mongodb.CreateCollection(collections.COLLECTIONS_ALL_USERS).FindOne(context.Background(), filter).Decode(&Users{})
	if err == nil {
		return status.Errorf(
			codes.Aborted,
			fmt.Sprintln(errormsg.ERR_MSG_EMAIL_ALREADY_EXIST))
	}
	return nil
}

// validate user email
func ValidateUserEmail(email string) (*Users, error) {
	data := &Users{}
	filter := bson.M{collections.PARAM_EMAIL: email}
	err := mongodb.CreateCollection(collections.COLLECTIONS_ALL_USERS).FindOne(context.Background(), filter).Decode(data)
	if err != nil {
		return data, status.Errorf(
			codes.Aborted,
			fmt.Sprintln(errormsg.ERR_MSG_EMAIL_NOT_FOUND))
	}
	return data, nil
}
