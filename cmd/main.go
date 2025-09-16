package main

import (
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/westcrime/auth/internal/config"
	grpcserver "github.com/westcrime/auth/internal/grpc_server"
	hashCrypto "github.com/westcrime/auth/internal/hash/crypto"
	urp "github.com/westcrime/auth/internal/user_repository/postgres"
	usp "github.com/westcrime/auth/internal/user_service/postgres"
	desc "github.com/westcrime/auth/pkg/user_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to load grpc config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}
	pool, err := pgxpool.New(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	hashService := &hashCrypto.SHA256Hasher{}
	ur := urp.NewUserRepository(pool)
	us := usp.NewUserService(ur, hashService)

	desc.RegisterUserV1Server(s, grpcserver.NewServer(us))

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
