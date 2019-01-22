package chat

import (
	"context"
	"fmt"
	"log"
	"o2clock/api-proto/home/chat"
	"o2clock/collection/allusers"
	"o2clock/constants/collections"
	"o2clock/constants/errormsg"
	"o2clock/db/mongodb"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	objectid "github.com/mongodb/mongo-go-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	VERSION = "v1.0"
)

type P2PChat struct {
	ID          objectid.ObjectID `bson:"_id,omitempty"`
	SenderId    string            `bson:"sender_id"`
	ReciverId   string            `bson:"reciver_id"`
	SenderName  string            `bson:"sender_name"`
	ReciverName string            `bson:"reciver_name"`
	ChatId      objectid.ObjectID `bson:"chat_id,omitempty"`
	Version     string            `bson:"version"`
	CaptureTime time.Time         `bson:"capture_time"`
}

type SendAndReciveP2PChat struct {
	ID            objectid.ObjectID `bson:"_id,omitempty"`
	SenderId      string            `bson:"sender_id"`
	ReciverId     string            `bson:"reciver_id"`
	ChatId        objectid.ObjectID `bson:"chat_id,omitempty"`
	Message       string            `bson:"message"`
	IsMessageRead bool              `bson:"is_message_read"`
	Version       string            `bson:"version"`
	CaptureTime   time.Time         `bson:"capture_time"`
}

/**
*
* Get all the users
*
**/
func GetAllUsers(req *chatpb.CommonRequest) ([]*chatpb.User, error) {
	return allusers.GetAllUsers(req)
}

/**
*
* Get loged in used info
*
**/
func GetUserInfo(req *chatpb.CommonRequest) (*chatpb.User, error) {
	return allusers.GetUserInfo(req)
}

/**
*
* Start the Person to person chat
*
**/
func StartP2PChat(req *chatpb.P2PChatRequest) (string, error) {
	chatid, err := P2PReciverAndUserValidation(req)
	if err == nil {
		return chatid, err
	}

	log.Println("Error", err)

	data := P2PChat{
		SenderId:    req.GetUserInfo().GetUserId(),
		SenderName:  req.GetUserInfo().GetUserName(),
		ReciverId:   req.GetReciverInfo().GetUserId(),
		ReciverName: req.GetReciverInfo().GetUserName(),
		Version:     VERSION,
		CaptureTime: time.Now(),
	}
	res, err := mongodb.CreateCollection(collections.COLLECTIONS_ALL_P2P_CHATS).InsertOne(context.Background(), data)
	if err != nil {
		return "", status.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_MSG_INTERNAL, err))
	}

	id, ok := res.InsertedID.(objectid.ObjectID)
	if !ok {
		return "", status.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_INTERNAL_OID, ok))
	}
	return id.String(), nil
}

/**
*
*  P2P chat validation
*
**/
func P2PReciverAndUserValidation(req *chatpb.P2PChatRequest) (string, error) {
	data := &P2PChat{}
	filter := bson.M{collections.PARAM_SENDER_ID: req.GetUserInfo().GetUserId(), collections.PARAM_RECIVER_ID: req.GetReciverInfo().GetUserId()}
	res := mongodb.CreateCollection(collections.COLLECTIONS_ALL_P2P_CHATS).FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		filter := bson.M{collections.PARAM_SENDER_ID: req.GetReciverInfo().GetUserId(), collections.PARAM_RECIVER_ID: req.GetUserInfo().GetUserId()}
		res := mongodb.CreateCollection(collections.COLLECTIONS_ALL_P2P_CHATS).FindOne(context.Background(), filter)
		if err := res.Decode(data); err != nil {
			return "", status.Errorf(
				codes.Aborted,
				fmt.Sprintln(errormsg.ERR_MSG_DATA_CANT_DECODE))
		} else {
			return data.ID.String(), nil
		}
	} else {
		return data.ID.String(), nil
	}
}
