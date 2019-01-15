package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"o2clock/api-desc/home"
	"o2clock/api-desc/home/chat"
	"o2clock/api-desc/home/logout"
	"o2clock/api-desc/onboarding/accesstoken"
	"o2clock/api-desc/onboarding/login"
	"o2clock/api-desc/onboarding/register"
	"o2clock/base/server"
	"o2clock/constants/appconstant"
	"o2clock/constants/errormsg"
	"o2clock/db/mongodb"
	db "o2clock/db/postgres"
	dbsettings "o2clock/settings/db"
	"o2clock/utils/log"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Init() {
	logsSetup()
	if dbsettings.IsEnableMongoDb() {
		MongodbSetup()
	}
	if dbsettings.IsEnablePostgres() {
		PgSetup()
	}
	ServerSetup()
}

func initLogs() {
	logLevel := flag.Int(appconstant.LOG_TXT_LOG_LEVEL, 1, appconstant.LOG_TXT_INTEGER_VAL)
	flag.Parse()
	// save logs in log.txt
	log.SetLogLevel(log.Level(*logLevel), appconstant.LOG_FILE_NAME)
	err := errors.New(errormsg.LOG_TXT_DUMMY)
	//to trace the error
	log.Info.Println(err)
}

func logsSetup() {
	initLogs()
}

func PgSetup() {
	db.DBConnecting()
}

func MongodbSetup() {
	if err := mongodb.InitDB(); err != nil {
		//for saving
		log.Error.Println(errormsg.ERR_MONGO, err)
		return
	}
}

func main() {
	Init()
}

//brain setup
func ServerSetup() {
	listner, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Error.Println(errormsg.ERR_SERVER_LISTEN, err)
		return
	}
	//Create the New gRPC Server
	srv := server.CreateNewgRPCServer(false, nil)
	//Register reflection on gRPC server
	reflection.Register(srv)
	//Register All the Services
	rpcServices(srv)

	go func() {
		fmt.Println(appconstant.LOG_SERVER_START)
		if err := srv.Serve(listner); err != nil {
			log.Error.Fatal(errormsg.ERR_SERVER_SERVE, err)
			return
		}
	}()
	//make a channnel that will wait for server to close or interrupt by control^c
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// this will block the code while server
	<-ch
	log.Info.Println(errormsg.ERR_SERVER_EXIT, os.Interrupt)
}

// All the services
func rpcServices(srv *grpc.Server) {
	//Register the User
	register.RegisterUserService(srv)
	//access token verify service
	accesstoken.RegisterAccessTokenService(srv)
	//logout user
	logout.RegisterLogoutService(srv)
	//login user
	login.RegisterLoginUserService(srv)
	//verify user
	home.RegisterVerifyUserService(srv)
	//chat room
	chat.RegisterChatRoomService(srv)

}
