package chat

import (
	"context"
	"net/http"
	"o2clock/api-proto/home/chat"
	"o2clock/constants/appconstant"

	mdb "o2clock/collection/chat"
	dbsettings "o2clock/settings/db"

	"google.golang.org/grpc"
)

type Server struct {
}

func RegisterChatRoomService(srv *grpc.Server) {
	chatpb.RegisterChatRoomServer(srv, &Server{})
}

/**
*
* Get all the users
*
**/
func (*Server) GetUsersList(ctx context.Context, req *chatpb.CommonRequest) (*chatpb.UserList, error) {
	var err error
	var data []*chatpb.User
	if dbsettings.IsEnableMongoDb() {
		data, err = mdb.GetAllUsers(req)
	}
	if dbsettings.IsEnablePostgres() {
		//err = pdb.LogoutUser(req)
	}
	if err == nil {
		//success
		return &chatpb.UserList{
			Users: data,
			Total: int32(len(data)),
			CommonResponse: &chatpb.CommonResponse{
				Message: appconstant.MSG_SUCCESS,
				Code:    http.StatusOK,
			},
		}, nil
	}
	return nil, err
}

// chat stream for both side bi-directional streaming
func (*Server) Chat(stream chatpb.ChatRoom_ChatServer) error {
	return nil
}

// single chat p2p
func (*Server) StartP2PChat(ctx context.Context, req *chatpb.P2PChatRequest) (*chatpb.CommonResponse, error) {
	return nil, nil
}

// Get all the chats
func (*Server) UserChats(ctx context.Context, req *chatpb.CommonRequest) (*chatpb.AllChats, error) {
	return nil, nil

}

// Create chat group
func (Server) CreateGroups(ctx context.Context, req *chatpb.CreateGroup) (*chatpb.CommonResponse, error) {
	return nil, nil

}

// Join chat group
func (*Server) JoinGroup(ctx context.Context, req *chatpb.Group) (*chatpb.CommonResponse, error) {
	return nil, nil

}

// Delete Chat group
func (*Server) DeleteGroup(ctx context.Context, req *chatpb.Group) (*chatpb.CommonResponse, error) {
	return nil, nil

}

// Get all user in a particular group
func (*Server) GetGroupUserList(ctx context.Context, req *chatpb.Group) (*chatpb.UserList, error) {
	return nil, nil
}

// Get all Group list
func (*Server) GetGroupList(ctx context.Context, req *chatpb.CommonRequest) (*chatpb.GroupList, error) {
	return nil, nil

}

// Add user in a group
func (*Server) AddUserInGroup(ctx context.Context, req *chatpb.AddMember) (*chatpb.CommonResponse, error) {
	return nil, nil

}

// Kick out the user from the group
func (*Server) KickoutUserFromGroup(ctx context.Context, req *chatpb.KickMember) (*chatpb.CommonResponse, error) {
	return nil, nil

}

// Get the chat history
func (*Server) GetChatHistory(ctx context.Context, req *chatpb.ReadHistoryRequest) (*chatpb.ReadHistoryResponse, error) {
	return nil, nil

}

// Leave the group
func (Server) LeaveGroup(ctx context.Context, req *chatpb.LeaveGroupRequest) (*chatpb.CommonResponse, error) {
	return nil, nil

}
