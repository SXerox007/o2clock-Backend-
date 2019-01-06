package login

import (
	"database/sql"
	"fmt"
	"o2clock/api-proto/onboarding/login"
	"o2clock/constants/errormsg"
	db "o2clock/db/postgres"
	"o2clock/table/accesstoken"
	"o2clock/table/allusers"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	SQL_STATEMENT_FIND_USER_USING_USERNAME = `
	SELECT * FROM all_users WHERE user_name=$1 && password=$2;`
	SQL_STATEMENT_FIND_USER_USING_EMAIL = `
	SELECT * FROM all_users WHERE email=$1 && password=$2;`
)

func LoginUser(req *loginpb.LoginRequest) (string, error) {
	user, err := validateUser(req)
	if err != nil {
		return "", err
	}
	return accesstoken.UpdateAccessToken(user.Id, user.UserName)
}

func validateUser(req *loginpb.LoginRequest) (allusers.Users, error) {
	var user allusers.Users
	sqlStatement := SQL_STATEMENT_FIND_USER_USING_USERNAME
	row := db.GetClient().QueryRow(sqlStatement, req.GetUsernameEmail(), req.GetPassword())
	err := row.Scan(&user.Id, &user.UserName)
	switch err {
	case sql.ErrNoRows:
		sqlStatement := SQL_STATEMENT_FIND_USER_USING_EMAIL
		row := db.GetClient().QueryRow(sqlStatement, req.GetUsernameEmail(), req.GetPassword())
		if err := row.Scan(&user.Id, &user.UserName); err != nil {
			return user, status.Errorf(
				codes.Internal,
				fmt.Sprintln(errormsg.ERR_MSG_INVALID_CREDS, err))
		} else {
			return user, nil
		}
	case nil:
		return user, nil
	default:
		return user, status.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_MSG_INTERNAL_SERVER, err))
	}

}
