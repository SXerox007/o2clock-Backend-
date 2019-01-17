package chat

import (
	"o2clock/api-proto/home/chat"
	"o2clock/table/allusers"
)

/**
*
* Get all the user from db
*
**/
func GetAllUsers(req *chatpb.CommonRequest) ([]*chatpb.User, error) {
	return allusers.GetAllUsers(req)
}
