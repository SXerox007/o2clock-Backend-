package resetpswd

import (
	"net/http"
	"o2clock/collection/forgotpassword/resetpswd"
)

// reset password handler
func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	resetpswd.UserResetPassword("http://01ee3bcd.ngrok.io"+r.RequestURI, w, r)
}
