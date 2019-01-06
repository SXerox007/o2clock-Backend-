package accesstoken

import (
	"crypto/rsa"
	"database/sql"
	"fmt"
	"io/ioutil"
	"o2clock/constants/errormsg"
	"o2clock/core/jwtauth"
	db "o2clock/db/postgres"
	"time"

	"o2clock/api-proto/onboarding/accesstoken"

	jwt "github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccessToken struct {
	TableName   struct{}  `sql:"all_tokens" json:"-"`
	UserId      string    `param:"user_id"`
	AccessToken string    `param:"access_token"`
	Version     string    `param:"version"`
	CaptureTime time.Time `param:"capture_time"`
}

const (
	VERSION                    = "v1.0"
	SQL_STATEMENT_CREATE_TOKEN = `
	INSERT INTO all_tokens (user_id, access_token, version,capture_time)
	VALUES ($1, $2, $3,$4)
	RETURNING id`
	SQL_STATEMENT_CHECK_TOKEN  = `SELECT * FROM all_tokens WHERE access_token=$1;`
	SQL_STATEMENT_UPDATE_TOKEN = `
	UPDATE all_tokens
	SET access_token = $2
	WHERE id = $1;`
)

func CreateUserAccessToken(userid string, username string) (string, error) {
	access_token, err := jwtauth.GenerateToken(createKey(), username, userid)
	if err == nil {
		sqlStatement := SQL_STATEMENT_CREATE_TOKEN
		var id string
		err := db.GetClient().QueryRow(sqlStatement, userid, access_token, VERSION, time.Now()).Scan(&id)
		if err != nil {
			return "", status.Errorf(
				codes.Internal,
				fmt.Sprintln(errormsg.ERR_MSG_INTERNAL_SERVER, err))
		}
		return access_token, nil
	} else {
		return "", status.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_MSG_INTERNAL_SERVER, err))
	}
}

func createKey() *rsa.PrivateKey {
	privateKey, err := ioutil.ReadFile("secure-keys/o2clock.rsa")
	if err != nil {
		return nil
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil
	}
	return key
}

func CheckAccessToken(req *accesstokenpb.AccessTokenRequest) error {
	sqlStatement := SQL_STATEMENT_CHECK_TOKEN
	var accessToken AccessToken
	row := db.GetClient().QueryRow(sqlStatement, req.GetAccessToken())
	switch err := row.Scan(&accessToken.UserId); err {
	case sql.ErrNoRows:
		return status.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_MSG_INVALID_ACCESS_TOKEN, err))
	case nil:
		return nil
	default:
		return status.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_MSG_INTERNAL_SERVER, err))
	}
}

func UpdateAccessToken(userid string, username string) (string, error) {
	access_token, err := jwtauth.GenerateToken(createKey(), username, userid)
	if err != nil {
		return "", status.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_MSG_INTERNAL_SERVER, err))
	}
	sqlStatement := SQL_STATEMENT_UPDATE_TOKEN
	res, err := db.GetClient().Exec(sqlStatement, 5, access_token)
	if err != nil {
		return "", err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return "", err
	}
	return "", nil
}
