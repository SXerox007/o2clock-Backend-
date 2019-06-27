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
)

// status set for chat
const (
	ONLINE  = 1
	OFFLINE = 2
	EXIT    = 3
)
