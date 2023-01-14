package grpc

import (
	"context"

	pb "git.ecobin.ir/ecomicro/protobuf/template/grpc"
	"git.ecobin.ir/ecomicro/template/domain"
	"git.ecobin.ir/ecomicro/tooty"
	"google.golang.org/grpc"
)

type userGrpcHandler struct {
	Usecase domain.UserUsecase
	pb.UnimplementedTemplateServer
}

func NewUserHandler(server *grpc.Server, usecase domain.UserUsecase) {
	pb.RegisterTemplateServer(server, &userGrpcHandler{Usecase: usecase})
}

// get user data ; other services use this to get user data
func (ugh *userGrpcHandler) CreateUser(
	ctx context.Context,
	protoUser *pb.CreateUserRequest,
) (*pb.CreateUserResponse, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[G] create user", "gRPC handler")
	defer tooty.CloseTheAPMSpan(span)

	if protoUser == nil {
		return nil, domain.ErrUnprocessableEntity
	}

	user, err := ugh.Usecase.Create(ctx, createUserToDomain(protoUser))
	if err != nil {
		return nil, GrpcResponseError(err, errMap)
	}
	return toCreateUserResponse(user), nil
}
func (ugh *userGrpcHandler) UpdateUser(
	ctx context.Context,
	protoUser *pb.UpdateUserRequest,
) (*pb.UpdateUserResponse, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[G] update user", "gRPC handler")
	defer tooty.CloseTheAPMSpan(span)

	if protoUser == nil {
		return nil, domain.ErrUnprocessableEntity
	}

	user, err := ugh.Usecase.Update(ctx, updateUserToDomain(protoUser))
	if err != nil {
		return nil, GrpcResponseError(err, errMap)
	}
	return toUpdateUserResponse(user), nil
}
func (ugh *userGrpcHandler) GetUser(
	ctx context.Context,
	protoUser *pb.GetUserRequest,
) (*pb.GetUserResponse, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[G] get user", "gRPC handler")
	defer tooty.CloseTheAPMSpan(span)

	if protoUser == nil {
		return nil, domain.ErrUnprocessableEntity
	}

	user, err := ugh.Usecase.GetUserById(ctx, protoUser.UserId)
	if err != nil {
		return nil, GrpcResponseError(err, errMap)
	}
	return toGetUserResponse(user), nil
}
