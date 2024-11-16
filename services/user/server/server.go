package server

import (
	"context"
	"log/slog"
	"net"

	"github.com/1001bit/pathgoer/services/user/shared/userpb"
	"github.com/1001bit/pathgoer/services/user/usermodel"
	"google.golang.org/grpc"
)

type UserStorage interface {
	GetNameAndEmailByLogin(ctx context.Context, login string) (string, string, error)
	GetProfileByName(ctx context.Context, name string) (*usermodel.Profile, error)
	GetNameAndIdByEmail(ctx context.Context, email string) (string, string, error)
	CreateUserGetNameAndId(ctx context.Context, email string) (string, string, error)
	GetNameById(ctx context.Context, id string) (string, error)
	ChangeUsername(ctx context.Context, id, newName string) error
}

type OtpStorage interface {
	GenerateOTP(ctx context.Context, email string) (string, error)
	VerifyOTP(ctx context.Context, email, otp string) bool
}

type RefreshStorage interface {
	GenerateUUID(ctx context.Context, userID string) (string, error)
	DeleteUUID(ctx context.Context, uuid string) error
	GetUserIDAndRefresh(ctx context.Context, uuid string) (string, string, error)
}

type QueuePublisher interface {
	Publish(queueName string, body string) error
}

type Server struct {
	userpb.UnimplementedUserServiceServer

	userStore      UserStorage
	otpStorage     OtpStorage
	refreshStorage RefreshStorage
	publisher      QueuePublisher
}

func New(userStore UserStorage, otpStorage OtpStorage, refreshStorage RefreshStorage, publisher QueuePublisher) *Server {
	return &Server{
		userStore:      userStore,
		otpStorage:     otpStorage,
		refreshStorage: refreshStorage,
		publisher:      publisher,
	}
}

func (s *Server) Start(addr string) error {
	slog.With("addr", addr).Info("Starting gRPC server")

	// start tcp listener
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// create grpc server
	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, s)

	// start grpc server
	return grpcServer.Serve(lis)
}
