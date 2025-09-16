package grpcserver

import (
	"context"
	"log"

	"github.com/westcrime/auth/internal/converter"
	userservice "github.com/westcrime/auth/internal/user_service"
	desc "github.com/westcrime/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	desc.UnimplementedUserV1Server
	userService userservice.UserService
}

func NewServer(usp userservice.UserService) *Server {
	return &Server{userService: usp}
}

func (s *Server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("User id: %d", req.GetId())
	err, user := s.userService.Get(ctx, req.GetId())
	return &desc.GetResponse{
		User: &desc.User{
			Id: user.Id,
			Info: &desc.UserInfo{
				Name:  user.Name,
				Email: user.Email,
				Role:  desc.Role(user.Role),
			},
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
	}, err
}

func (s *Server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("User id: %d", req.GetId())

	err := s.userService.Delete(ctx, req.GetId())

	return &emptypb.Empty{}, err
}

func (s *Server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("User: %+v", req.Info)
	err, id := s.userService.Create(ctx, converter.ToCreateUserModelFromCreateRequestProto(req))

	return &desc.CreateResponse{
		Id: id}, err
}

func (s *Server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("User id: %d", req.GetId())
	err := s.userService.Update(ctx, converter.ToUpdateUserModelFromUpdateRequestProto(req))

	return &emptypb.Empty{}, err
}
