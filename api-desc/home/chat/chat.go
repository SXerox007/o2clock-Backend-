package chat

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"o2clock/api-proto/home/chat"
	"o2clock/constants/appconstant"
	"o2clock/constants/errormsg"
	"sync"

	mdb "o2clock/collection/chat"
	pdb "o2clock/table/chat"

	dbsettings "o2clock/settings/db"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
}

// Group info
type Group struct {
	Name      string
	Ch        chan chatpb.ChatMessage
	Clients   []string
	WaitGroup *sync.WaitGroup
}

//client info
type Client struct {
	Name      string
	Groups    []string
	Ch        chan chatpb.ChatMessage
	WaitGroup *sync.WaitGroup
}

var lock = &sync.RWMutex{}
var clients = make(map[string]*Client)
var groups = make(map[string]*Group)

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
		data, err = pdb.GetAllUsers(req)
	}
	RegisterAllClients(data)
	RegisterP2PChats()
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

// Get signin user details
func (*Server) GetUserDetails(ctx context.Context, req *chatpb.CommonRequest) (*chatpb.User, error) {
	var err error
	var data *chatpb.User
	if dbsettings.IsEnableMongoDb() {
		data, err = mdb.GetUserInfo(req)
	}
	if dbsettings.IsEnablePostgres() {
		//data, err = pdb.GetAllUsers(req)
	}
	if err == nil {
		//success
		return data, nil
	}
	return nil, err
}

// chat stream for both side bi-directional streaming
func (*Server) Chat(stream chatpb.ChatRoom_ChatServer) error {
	msg, err := stream.Recv()
	if err != nil {
		return err
	}
	//outbox := make(chan chatpb.ChatMessage, 100)
	//go ListenToClient(stream, outbox)
	// log.Println("All Clients:", clients)
	// log.Println("All Groups:", groups)
	// for {
	// 	select {
	// 	case outMsg := <-outbox:
	// 		fmt.Println("Here at case 1:", outMsg)
	//broadcast msg to all the group members
	Broadcast(msg.GetChatId(), *msg)
	//	break
	//case inMsg := <-clients[msg.GetSingleMessage().GetReciverId()].Ch:
	//send msg to a single particular group
	//	fmt.Println("Here at case 2:", inMsg)
	//stream.Send(&inMsg)
	//}
	//	}
	return nil
}

// ListenToClient listens on the incoming stream for any messages. It adds those messages to the channel.
func ListenToClient(stream chatpb.ChatRoom_ChatServer, messages chan<- chatpb.ChatMessage) {
	fmt.Println("Channel print:", messages)

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			//	fmt.Println("End of file")
		} else if err != nil {
			//log.Error.Println("Error when listen to client:", err)
		} else {
			//fmt.Printf("[ListenToClient] Client ", msg.GetMessage())
			mdb.SaveChatMessage(msg)
			messages <- *msg
		}
	}
}

// ClientExists checks if a client exists on the server.
// It returns a bool value.
func ClientExists(n string) bool {

	lock.RLock()
	defer lock.RUnlock()
	for c := range clients {
		if c == n {
			return true
		}
	}

	return false
}

// Register all the users for chats
func RegisterAllClients(arr []*chatpb.User) {
	lock.Lock()
	defer lock.Unlock()
	var i int
	for i = 0; i < len(arr); i++ {
		c := &Client{
			Name:      arr[i].GetUserName(),
			Ch:        make(chan chatpb.ChatMessage, 100),
			WaitGroup: &sync.WaitGroup{},
		}
		//fmt.Print("[RegisterAllClient]: Registered client " + arr[i].GetUserName())
		clients[arr[i].GetUserId()] = c
	}
}

/**
*
* Register all the P2P Chats
*
**/
func RegisterP2PChats() error {
	var i int
	data, err := mdb.GetP2PAllChats()
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_MSG_INTERNAL, err))
	}
	lock.Lock()
	defer lock.Unlock()
	for i = 0; i < len(data); i++ {
		g := &Group{
			Name:      data[i].ID.String(),
			Ch:        make(chan chatpb.ChatMessage, 100),
			WaitGroup: &sync.WaitGroup{},
		}

		//fmt.Print("[RegsiterP2PGroups]: P2P Groups Register ", data[i].ID.String())
		groups[data[i].ID.String()] = g
		groups[data[i].ID.String()].WaitGroup.Add(1)
		//add sender and reciver to the groups
		AddClientToGroup(data[i].ReciverId, data[i].ID.String())
		AddClientToGroup(data[i].SenderId, data[i].ID.String())
	}
	return nil
}

// AddGroup adds a new group to the server.
func AddGroup(groupName string) {

	lock.Lock()
	defer lock.Unlock()

	g := &Group{
		Name:      groupName,
		Ch:        make(chan chatpb.ChatMessage, 100),
		WaitGroup: &sync.WaitGroup{},
	}

	//fmt.Print("[AddGroup]: Added group ")
	groups[groupName] = g
	groups[groupName].WaitGroup.Add(1)
}

// GroupExists checks if a group exists on the server.
// It returns a bool value.
func GroupExists(gName string) bool {

	lock.RLock()
	defer lock.RUnlock()
	for g := range groups {
		if g == gName {
			return true
		}
	}

	return false
}

// InGroup checks whether a client is currently in a
// specific group.
// It returns a bool value.
func InGroup(n string) bool {

	for _, g := range groups {
		for _, c := range g.Clients {
			if n == c {
				return true
			}
		}
	}

	return false
}

// RemoveClient will remove a client from the server as well as any
// groups that they are currently in.
// It returns an error.
func RemoveClient(name string) error {

	// TODO: There is some deadlock here when a user attempts to quit
	// 		 the chat app with !exit.

	lock.Lock()
	defer lock.Unlock()

	if ClientExists(name) {
		delete(clients, name)
		//fmt.Print("[RemoveClient]: Removed client " + name)
		if InGroup(name) {
			RemoveClientFromGroup(name)
		} else {
			//fmt.Print("[RemoveClient]: " + name + " was not in any groups.")
			return nil
		}
	}

	return errors.New("[RemoveClient]: Client (" + name + ") doesn't exist")
}

// AddClientToGroup will add a client to a group.
// It doesn't return anything.
func AddClientToGroup(c string, g string) {

	//lock.Lock()
	//defer lock.Unlock()

	groups[g].WaitGroup.Add(1)
	defer groups[g].WaitGroup.Done()

	groups[g].Clients = append(groups[g].Clients, c)
	clients[c].Groups = append(clients[c].Groups, g)

	//fmt.Println("[AddClientToGroup] Added " + c + " to " + g)
}

// RemoveClientFromGroup will remove a client from a specific group. It will also
// delete a group if the client is the last one leaving it.
func RemoveClientFromGroup(n string) error {

	for _, g := range groups {
		for i, c := range g.Clients {
			if n == c {
				c := clients[n].Groups
				// Remove the group from the user.
				for i, _ := range c {
					if n == g.Name {
						c[i] = c[len(c)-1]
						c = c[:len(c)-1]
						clients[n].Groups = c
					}
				}
				if len(g.Clients) == 1 {
					delete(groups, g.Name)
				} else {
					c := g.Clients
					c[i] = c[len(c)-1]
					c = c[:len(c)-1]
					g.Clients = c
				}
				return nil
			}
		}
	}

	return errors.New("no user found in the group list. Something went wrong")
}

// Broadcast takes any messages that need to be sent and sorts them by group. It then
// adds the message to the channel of each member of that group.
// It doesn't return anything.
func Broadcast(gName string, msg chatpb.ChatMessage) {

	lock.Lock()
	defer lock.Unlock()

	for gn := range groups {
		fmt.Printf("[Broadcast]: I found " + gn + ".")
		if gn == gName {
			fmt.Printf("[Broadcast]: Client " + msg.SenderName + " sent " + msg.GetSingleMessage().GetReciverName() + " a message: " + msg.GetMessage())
			for _, c := range groups[gn].Clients {
				fmt.Printf("[Broadcast]: I found " + c + " in gName")
				// if c == msg.Sender && msg.Message == msg.Sender+" left chat!\n" {
				// 	log.Printf("[Broadcast]: ADDING THE KILL MESSAGE TO " + c)
				// 	clients[c].ch <- msg
				// } else
				if c != msg.GetSenderid() {
					fmt.Printf("[Broadcast] Adding the message to " + c + "'s channel.")
					clients[c].Ch <- msg
				}
			}
		}
	}
}

// single chat p2p
func (*Server) StartP2PChat(ctx context.Context, req *chatpb.P2PChatRequest) (*chatpb.P2PChatResponse, error) {
	var err error
	var id string
	if dbsettings.IsEnableMongoDb() {
		id, err = mdb.StartP2PChat(req)
	}
	if dbsettings.IsEnablePostgres() {
		//data, err = pdb.GetAllUsers(req)
	}
	if err == nil {
		return &chatpb.P2PChatResponse{
			ChatId: id,
			CommmonResponse: &chatpb.CommonResponse{
				Message: appconstant.MSG_SUCCESS,
				Code:    http.StatusOK,
			},
		}, nil
	}
	return nil, err
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
