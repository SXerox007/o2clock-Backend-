package chat

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"o2clock/api-proto/home/chat"
	"o2clock/constants/appconstant"
	"sync"

	mdb "o2clock/collection/chat"
	pdb "o2clock/table/chat"

	dbsettings "o2clock/settings/db"

	"google.golang.org/grpc"
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
	outbox := make(chan chatpb.ChatMessage, 100)
	go ListenToClient(stream, outbox)
	for {
		select {
		case outMsg := <-outbox:
			//broadcast msg to all the group members

		case inMsg := <-clients["sd"].Ch:
			//send msg to a single particular group

			stream.Send(&inMsg)
		}
	}
	return nil
}

// ListenToClient listens on the incoming stream for any messages. It adds those messages to the channel.
func ListenToClient(stream chatpb.ChatRoom_ChatServer, messages chan<- chatpb.ChatMessage) {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
		}
		if err != nil {
		} else {
			log.Printf("[ListenToClient] Client ", msg.GetMessage())
			messages <- *msg
		}
	}
}

// AddClient adds a new client n to the server.
func AddClient(n string) {

	lock.Lock()
	defer lock.Unlock()

	c := &Client{
		Name:      n,
		Ch:        make(chan chatpb.ChatMessage, 100),
		WaitGroup: &sync.WaitGroup{},
	}

	log.Print("[AddClient]: Registered client " + n)
	clients[n] = c
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

// AddGroup adds a new group to the server.
func AddGroup(n string) {

	lock.Lock()
	defer lock.Unlock()

	g := &Group{
		Name:      n,
		Ch:        make(chan chatpb.ChatMessage, 100),
		WaitGroup: &sync.WaitGroup{},
	}

	log.Print("[AddGroup]: Added group ")
	groups[n] = g
	groups[n].WaitGroup.Add(1)
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
		log.Print("[RemoveClient]: Removed client " + name)
		if InGroup(name) {
			RemoveClientFromGroup(name)
		} else {
			log.Print("[RemoveClient]: " + name + " was not in any groups.")
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

	log.Println("[AddClientToGroup] Added " + c + " to " + g)
}

// RemoveClientFromGroup will remove a client from a specific group. It will also
// delete a group if the client is the last one leaving it.
// It returns an error.
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
