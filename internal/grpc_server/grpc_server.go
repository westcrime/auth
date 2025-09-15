package grpcserver

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/westcrime/auth/internal/converter"
	"github.com/westcrime/auth/internal/hash"
	userservice "github.com/westcrime/auth/internal/user_service"
	desc "github.com/westcrime/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	desc.UnimplementedUserV1Server
	Pool *pgxpool.Pool
}

func (s *Server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("User id: %d", req.GetId())
	err, user := userservice.Get(ctx, s.Pool, req.GetId())
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

	err := userservice.Delete(ctx, s.Pool, req.GetId())

	return &emptypb.Empty{}, err
}

func (s *Server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("User: %+v", req.Info)
	err, id := userservice.Create(ctx, s.Pool, converter.ToCreateUserModelFromCreateRequestProto(req), &hash.SHA256Hasher{})

	return &desc.CreateResponse{
		Id: id}, err
}

func (s *Server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("User id: %d", req.GetId())
	err := userservice.Update(ctx, s.Pool, converter.ToUpdateUserModelFromUpdateRequestProto(req))

	return &emptypb.Empty{}, err
}
