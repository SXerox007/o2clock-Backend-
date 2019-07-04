package resetpswd

import (
	"net/http"
	"o2clock/collection/forgotpassword/resetpswd"
)

// reset password handler
func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		resetpswd.SetUserNewPassword(w, r)
	} else {
		resetpswd.GetUserResetPassword("http://01ee3bcd.ngrok.io"+r.RequestURI, w, r)
	}
}
