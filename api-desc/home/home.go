package home

import (
	"io"
	"log"
	"net/http"
	homepg "o2clock/api-proto/home"
	"o2clock/collection/home"
	"o2clock/constants/appconstant"

	"google.golang.org/grpc"
)

type Server struct {
}

func RegisterVerifyUserService(srv *grpc.Server) {
	homepg.RegisterVerifyServiceServer(srv, &Server{})
}

func (*Server) UserVerifyService(stream homepg.VerifyService_UserVerifyServiceServer) error {
	var result []byte
	for {
		data, err := stream.Recv()
		log.Println("Data Chunk:", data.GetFileChunk())
		if err == io.EOF {
			if err := home.VerifyUser(result, ""); err != nil {
				return stream.SendAndClose(&homepg.VerifyUserResponse{
					Message: appconstant.MSG_FAILURE,
					Code:    http.StatusNotImplemented,
				})
			} else {
				return stream.SendAndClose(&homepg.VerifyUserResponse{
					Message: appconstant.MSG_SUCCESS,
					Code:    http.StatusOK,
				})
			}
		}
		if err != nil {
			return err
		}
		result = append(result, data.GetFileChunk()...)
	}
	return nil
}
