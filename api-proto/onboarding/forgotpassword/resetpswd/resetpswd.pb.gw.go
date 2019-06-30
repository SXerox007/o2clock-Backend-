// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: api-proto/onboarding/forgotpassword/resetpswd/resetpswd.proto

/*
Package resetpswdpb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package resetpswdpb

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_ResetPasswordService_ResetPassowrdUserService_0(ctx context.Context, marshaler runtime.Marshaler, client ResetPasswordServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ResetPasswordRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	msg, err := client.ResetPassowrdUserService(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterResetPasswordServiceHandlerFromEndpoint is same as RegisterResetPasswordServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterResetPasswordServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterResetPasswordServiceHandler(ctx, mux, conn)
}

// RegisterResetPasswordServiceHandler registers the http handlers for service ResetPasswordService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterResetPasswordServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterResetPasswordServiceHandlerClient(ctx, mux, NewResetPasswordServiceClient(conn))
}

// RegisterResetPasswordServiceHandlerClient registers the http handlers for service ResetPasswordService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "ResetPasswordServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ResetPasswordServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ResetPasswordServiceClient" to call the correct interceptors.
func RegisterResetPasswordServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ResetPasswordServiceClient) error {

	mux.Handle("GET", pattern_ResetPasswordService_ResetPassowrdUserService_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ResetPasswordService_ResetPassowrdUserService_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ResetPasswordService_ResetPassowrdUserService_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_ResetPasswordService_ResetPassowrdUserService_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"v1", "user", "setpassword", "id"}, ""))
)

var (
	forward_ResetPasswordService_ResetPassowrdUserService_0 = runtime.ForwardResponseMessage
)