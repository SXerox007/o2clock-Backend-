package logout

import (
	"o2clock/api-proto/home/logout"
	db "o2clock/db/postgres"
)

const (
	SQL_STATEMENT_LOGOUT = `
	UPDATE all_tokens
	SET access_token = $2
	WHERE id = $1;`
)

func LogoutUser(req *logoutpb.LogoutRequest) error {
	sqlStatement := SQL_STATEMENT_LOGOUT
	res, err := db.GetClient().Exec(sqlStatement, 5, nil)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
