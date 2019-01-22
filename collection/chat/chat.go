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

/**
*
* All the Chat messages
*
**/
type AllChats struct {
	ID            objectid.ObjectID `bson:"_id,omitempty"`
	SenderId      string            `bson:"sender_id"`
	SenderName    string            `bson:"sender_name"`
	ReciverId     string            `bson:"reciver_id"`
	ReciverName   string            `bson:"reciver_name"`
	ChatId        string            `bson:"chat_id,omitempty"`
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
* Get all the P2P chats
*
**/
func GetP2PAllChats() ([]P2PChat, error) {
	res, err := mongodb.CreateCollection(collections.COLLECTIONS_ALL_P2P_CHATS).Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintln(errormsg.ERR_NOT_FOUND, err))
	}
	var data []P2PChat
	for res.Next(nil) {
		item := P2PChat{}
		if err := res.Decode(&item); err != nil {
			return nil, status.Errorf(
				codes.Aborted,
				fmt.Sprintln(errormsg.ERR_MSG_DATA_CANT_DECODE, err))
		}
		data = append(data, item)
	}
	return data, nil
}

/**
*
* Save the all Chat Messages
*
**/
func SaveChatMessage(msg *chatpb.ChatMessage) error {
	item := AllChats{
		SenderId:      msg.GetSenderid(),
		SenderName:    msg.GetSenderName(),
		ChatId:        msg.GetChatId(),
		Message:       msg.GetMessage(),
		IsMessageRead: false,
		Version:       VERSION,
		CaptureTime:   time.Now(),
	}
	if msg.GetIsGroupMessage() {

	} else {
		item.ReciverId = msg.GetSingleMessage().GetReciverId()
		item.ReciverName = msg.GetSingleMessage().GetReciverName()
	}

	res, err := mongodb.CreateCollection(msg.GetChatId()).InsertOne(context.Background(), item)
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

/**
*
* Get the chat history by chat id
*
**/
func GetUserChatHistory(chatId string) ([]AllChats, error) {

	res, err := mongodb.CreateCollection(chatId).Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintln(errormsg.ERR_NOT_FOUND, err))
	}
	var data []AllChats
	for res.Next(nil) {
		item := AllChats{}
		if err := res.Decode(&item); err != nil {
			return nil, status.Errorf(
				codes.Aborted,
				fmt.Sprintln(errormsg.ERR_MSG_DATA_CANT_DECODE, err))
		}
		data = append(data, item)
	}
	return data, nil
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
