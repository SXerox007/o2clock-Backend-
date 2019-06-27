package errormsg

//console log error msg
const (
	LOG_TXT_DUMMY     = "Start trace Important Error"
	ERR_MONGO         = "Can't connect to mongoDB: "
	ERR_SERVER_LISTEN = "Error in server listen: "
	ERR_SERVER_SERVE  = "Error in Serve the Server: "
	ERR_SERVER_EXIT   = "Exit the server with "
)

//Erro msg for users
const (
	ERR_MSG_USERNAME_ALREADY_EXIST = "Username already exists please choose another one."
	ERR_MSG_EMAIL_ALREADY_EXIST    = "Email id  already exists. If you forgot your password click on forgot password."
	ERR_MSG_INTERNAL               = "Internal MongoDb Error: "
	ERR_INTERNAL_OID               = "Internal oid Error: "
	ERR_MSG_INTERNAL_SERVER        = "Internal Server Error:"
	ERR_MSG_INVALID_ACCESS_TOKEN   = "Invalid Access Token"
	ERR_MSG_DATA_CANT_DECODE       = "Data Can't be decoded"
	ERR_MSG_INVALID_CREDS          = "Invalid Creds"
	ERR_MSG_PSWD_DECODE            = "Verify Pswd Error: "
	ERR_FACE_NOT_REC               = "Can't recognize the face"
	ERR_NOT_A_SINGLE_FACE          = "Not a single face on the image"
	ERR_NO_ROWS_Users              = "Can't get the Users"
	ERR_NOT_FOUND                  = "No Data found"
	ERR_STATUS                     = "Some error in update status"
	ERR_SLACK                      = "Error Slack: "
)
