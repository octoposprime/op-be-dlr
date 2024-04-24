package presentation

import (
	"context"

	dto "github.com/octoposprime/op-be-dlr/pkg/presentation/dto"
	pb_authentication "github.com/octoposprime/op-be-shared/pkg/proto/pb/authentication"
)

// Login generates an authentication token if the given user values are valid.
func (a *Grpc) Login(ctx context.Context, loginRequest *pb_authentication.LoginRequest) (*pb_authentication.Token, error) {
	data, err := a.commandHandler.Login(ctx, *dto.NewLoginRequest(loginRequest).ToObject())
	return dto.NewTokenFromObject(&data).ToPb(), err
}

// Refresh regenerate an authentication token.
func (a *Grpc) Refresh(ctx context.Context, token *pb_authentication.Token) (*pb_authentication.Token, error) {
	data, err := a.commandHandler.Refresh(ctx, *dto.NewToken(token).ToObject())
	return dto.NewTokenFromObject(&data).ToPb(), err
}

// Logout clears some footprints for the user.
func (a *Grpc) Logout(ctx context.Context, token *pb_authentication.Token) (*pb_authentication.LogoutResponse, error) {
	err := a.commandHandler.Logout(ctx, *dto.NewToken(token).ToObject())
	return &pb_authentication.LogoutResponse{}, err
}
