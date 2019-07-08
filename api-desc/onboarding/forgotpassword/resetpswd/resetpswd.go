package resetpswd

import (
	"net/http"
	"o2clock/collection/forgotpassword/resetpswd"
	"o2clock/constants/appconstant"
)

// reset password handler
func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		resetpswd.SetUserNewPassword(w, r)
	} else {
		resetpswd.GetUserResetPassword(appconstant.BASE_URL+r.RequestURI, w, r)
	}
}
