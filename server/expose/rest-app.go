package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"mime"
	"net/http"
	"o2clock/api-proto/home"
	"o2clock/api-proto/home/chat"
	"o2clock/api-proto/home/logout"
	"o2clock/api-proto/onboarding/accesstoken"
	"o2clock/api-proto/onboarding/forgotpassword"
	"o2clock/api-proto/onboarding/forgotpassword/resetpswd"
	"o2clock/swagger/pkg/ui/data/swagger"

	"o2clock/api-proto/onboarding/login"

	"o2clock/api-proto/onboarding/register"
	"o2clock/api-proto/webhooks/git"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/philips/go-bindata-assetfs"
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
	err = resetpswdpb.RegisterResetPasswordServiceHandlerFromEndpoint(ctx, mux, *authpoint, dialOpts)
	if err != nil {
		return err
	}
	// subMux := http.NewServeMux()
	// subMux.HandleFunc("/sub_path", TestHandler)

	// grpcMux := http.NewServeMux()
	// grpcMux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
	// 	io.Copy(w, strings.NewReader(pb.Swagger))
	// })

	log.Println("Starting Endpoint Exposed Server: localhost:5051")
	http.ListenAndServe(address, mux)
	return nil
}

// test handler
func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to o2Clock.")
}

func main() {
	Init()
}

func Init() {
	if err := ExposePoint(":5051"); err != nil {
		log.Fatal("Error in serve", err)
	}
}

// serve swagger
func serveSwagger(mux *http.ServeMux) {
	mime.AddExtensionType(".svg", "image/svg+xml")

	// Expose files in third_party/swagger-ui/ on <host>/swagger-ui
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
