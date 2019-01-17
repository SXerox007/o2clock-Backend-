package chat

import (
	"o2clock/api-proto/home/chat"
	"o2clock/collection/allusers"
)

/**
*
* Get all the users
*
**/
func GetAllUsers(req *chatpb.CommonRequest) ([]*chatpb.User, error) {
	return allusers.GetAllUsers(req)
}
