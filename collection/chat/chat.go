package chat

import (
	"log"
	"o2clock/api-proto/home/chat"
	"o2clock/collection/allusers"
)

/**
*
* Get all the users
*
**/
func GetAllUsers(req *chatpb.CommonRequest) ([]*chatpb.User, error) {
	log.Println("Access Token:", req.GetAccessToken())
	return allusers.GetAllUsers(req)
}
