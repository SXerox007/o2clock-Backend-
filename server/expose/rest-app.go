package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"o2clock/api-proto/home"
	"o2clock/api-proto/home/chat"
	"o2clock/api-proto/home/logout"
	"o2clock/api-proto/onboarding/accesstoken"
	"o2clock/api-proto/onboarding/forgotpassword"
	"o2clock/api-proto/onboarding/login"

	"o2clock/api-proto/onboarding/register"
	"o2clock/api-proto/webhooks/git"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	authpoint = flag.String("auth_end_points", "localhost:50051", "expose end point of oAuth")
)

func ExposePoint(address string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	err := regsiterpb.RegisterRegisterServiceHandlerFromEndpoint(ctx, mux, *authpoint, dialOpts)
	err = accesstokenpb.RegisterAccessTokenServiceHandlerFromEndpoint(ctx, mux, *authpoint, dialOpts)
	err = logoutpb.RegisterLogoutServiceHandlerFromEndpoint(ctx, mux, *authpoint, dialOpts)
	err = loginpb.RegisterLoginServiceHandlerFromEndpoint(ctx, mux, *authpoint, dialOpts)
	err = homepb.RegisterVerifyServiceHandlerFromEndpoint(ctx, mux, *authpoint, dialOpts)
	err = chatpb.RegisterChatRoomHandlerFromEndpoint(ctx, mux, *authpoint, dialOpts)
	err = githubpb.RegisterGithubWebhookServicesHandlerFromEndpoint(ctx, mux, *authpoint, dialOpts)
	err = forgotpasswordpb.RegisterForgotPasswordServiceHandlerFromEndpoint(ctx, mux, *authpoint, dialOpts)
	if err != nil {
		return err
	}
	log.Println("Starting Endpoint Exposed Server: localhost:5051")
	http.ListenAndServe(address, mux)
	return nil
}

func main() {
	Init()
}

func Init() {
	if err := ExposePoint(":5051"); err != nil {
		log.Fatal("Error in serve", err)
	}
}
