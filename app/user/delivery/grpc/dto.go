package grpc

import (
	"time"

	pb "git.ecobin.ir/ecomicro/protobuf/template/grpc"
	"git.ecobin.ir/ecomicro/template/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ValueOrDefault(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func createUserToDomain(createUserRequest *pb.CreateUserRequest) domain.User {
	return domain.User{
		Roles: createUserRequest.Roles,
		Allow: createUserRequest.Allow,
		Deny:  createUserRequest.Deny,
	}
}
func updateUserToDomain(updateUserRequest *pb.UpdateUserRequest) domain.User {
	return domain.User{
		Id:    updateUserRequest.UserId,
		Roles: updateUserRequest.Roles,
		Allow: updateUserRequest.Allow,
		Deny:  updateUserRequest.Deny,
	}
}
func toUpdateUserResponse(user *domain.User) *pb.UpdateUserResponse {
	return &pb.UpdateUserResponse{
		UserId:     user.Id,
		Roles:      user.Roles,
		Allow:      user.Allow,
		Deny:       user.Deny,
		CreateDate: ValueOrDefault(&user.CreatedDate),
		UpdateDate: ValueOrDefault(&user.UpdatedDate),
	}
}
func toCreateUserResponse(user *domain.User) *pb.CreateUserResponse {
	return &pb.CreateUserResponse{
		UserId:     user.Id,
		Roles:      user.Roles,
		Allow:      user.Allow,
		Deny:       user.Deny,
		CreateDate: ValueOrDefault(&user.CreatedDate),
		UpdateDate: ValueOrDefault(&user.UpdatedDate),
	}
}
func toGetUserResponse(user *domain.User) *pb.GetUserResponse {
	return &pb.GetUserResponse{
		UserId:     user.Id,
		Roles:      user.Roles,
		Allow:      user.Allow,
		Deny:       user.Deny,
		CreateDate: ValueOrDefault(&user.CreatedDate),
		UpdateDate: ValueOrDefault(&user.UpdatedDate),
	}
}
