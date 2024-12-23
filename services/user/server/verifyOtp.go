package server

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/1001bit/pathgoer/services/user/shared/accesstoken"
	"github.com/1001bit/pathgoer/services/user/shared/userpb"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) VerifyOtp(ctx context.Context, req *userpb.VerifyOtpRequest) (*userpb.TokensResponse, error) {
	otpExists, err := s.otpStorage.VerifyOTP(ctx, req.Email, req.Otp)
	if err == redis.Nil || !otpExists {
		return nil, status.Error(codes.NotFound, "not found")
	} else if err != nil {
		slog.With("err", err).Error("Failed to verify OTP")
		return nil, status.Error(codes.Internal, "an error occurred")
	}

	// Ask database for id by email
	id, err := s.userStore.GetIdByEmail(ctx, req.Email)
	if err == sql.ErrNoRows {
		// Create new user if doesn't exist
		id, err = s.userStore.CreateUserGetId(ctx, req.Email)
		if err != nil {
			slog.With("err", err).Error("Failed to create user")
			return nil, status.Error(codes.Internal, "an error occurred")
		}
	} else if err != nil {
		// if something went wrong
		slog.With("err", err).Error("Failed to get name and id by email")
		return nil, status.Error(codes.Internal, "an error occurred")
	}

	access, err := accesstoken.Generate(id)
	if err != nil {
		slog.With("err", err).Error("Failed to generate access token")
		return nil, status.Error(codes.Internal, "an error occurred")
	}

	refresh, err := s.refreshStorage.GenerateUUID(ctx, id)
	if err != nil {
		slog.With("err", err).Error("Failed to generate refresh token")
		return nil, status.Error(codes.Internal, "an error occurred")
	}

	return &userpb.TokensResponse{
		AccessJWT:   access,
		RefreshUUID: refresh,
	}, nil
}
