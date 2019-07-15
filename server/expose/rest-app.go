package main

import (
	"context"
	"crypto/x509"
	"flag"
	"io"
	"log"
	"net/http"
	"o2clock/api-desc/onboarding/forgotpassword/resetpswd"
	ctest "o2clock/collection/forgotpassword/resetpswd"
	"path"

	"o2clock/api-proto/home"
	"o2clock/api-proto/home/chat"
	"o2clock/api-proto/home/logout"
	"o2clock/api-proto/onboarding/accesstoken"
	"o2clock/api-proto/onboarding/forgotpassword"
	"strings"

	"o2clock/api-proto/onboarding/login"

	"o2clock/api-proto/onboarding/register"
	"o2clock/api-proto/webhooks/git"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type myService struct{}

var (
	authpoint    = flag.String("auth_end_points", "localhost:50051", "expose end point of oAuth")
	demoCertPool *x509.CertPool
)

func newServer() *myService {
	return new(myService)
}

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

	grpcMux := http.NewServeMux()
	grpcMux.HandleFunc("/v1/user/setpassword/", resetpswd.ResetPasswordHandler)
	grpcMux.HandleFunc("/test", TestHandler)

	grpcMux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, strings.NewReader(forgotpasswordpb.Swagger))
	})

	grpcMux.Handle("/", mux)
	// serveSwagger(grpcMux)
	grpcMux.HandleFunc("/swagger/", serveSwagger)

	log.Println("Starting Endpoint Exposed Server: localhost:5051")
	http.ListenAndServe(address, grpcMux)
	return nil
}

// test handler
func TestHandler(w http.ResponseWriter, r *http.Request) {
	ctest.OutputHTML(w, "./templates/static/test.html", nil)
	return
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
func serveSwagger(w http.ResponseWriter, r *http.Request) {
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join("swagger/third_party/swagger-new/swagger-ui/", p)
	http.ServeFile(w, r, p)
}
