package appconstant

//log file Error Messages
const (
	LOG_TXT_LOG_LEVEL   = "loglevel"
	LOG_TXT_INTEGER_VAL = "an integer value (0-4)"
	LOG_FILE_NAME       = "brain_logs.txt"
	LOG_SERVER_START    = "Starting Server: localhost:50051"
)

//status and messages
const (
	MSG_SUCCESS             = "Success"
	MSG_FAILURE             = "Error"
	PSWD_SUCCESS            = "Password Success  :)"
	PSWD_SUCCESS_MSG        = "Your change password is updated with success. Please login into app with the new password."
	MSG_SOMETING_WENT_WRONG = "Something Went Wrong"
)

//utils
const (
	MEM_ALLOC       = "Alloc: "
	MEM_TOTAL_ALLOC = "TotalAlloc: "
	MEM_SYS         = "Sys: "
	NUM_GC          = "NumGC: "
	LOOKUPS         = "Lookups: "
	MALLOCS         = "Mallocs: "
	ALPHA_NUM       = "Alpha-Num"
	ALPHA           = "Alpha"
	NUM             = "Num"
	DIC_ALPHA_NUM   = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	DIC_ALPHA       = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	DIC_NUM         = "0123456789"
)

// slack basic queries
const (
	SRVINFO = "srv -h"
)

// email constants
const (
	TXT_HTML                  = "text/html"
	TXT_PLAIN                 = "text/plain"
	EMAIL_TO                  = "To"
	EMAIL_FROM                = "From"
	EMAIL_SUBJECT             = "Subject"
	MIME_VERSION              = "MIME-Version"
	VERSION                   = "1.0"
	CONTENT_TYPE              = "Content-Type"
	CONTENT_TRANSFER_ENCODING = "Content-Transfer-Encoding"
	QUOTED_PRINT              = "quoted-printable"
	CONTENT_DISPOSITION       = "Content-Disposition"
	INLINE                    = "inline"
)

// status set for chat
const (
	ONLINE  = 1
	OFFLINE = 2
	EXIT    = 3
)

// all events
const (
	EVENT_FORGOT_PSWD = 5001
)

// all states
const (
	STATE_LINK_GENERATE = 1
	STATE_LINK_INACTIVE = 0
	STATE_LINK_ERROR    = 2
)

// URL
const (
	BASE_URL = "https://92de1297.ngrok.io"
)
