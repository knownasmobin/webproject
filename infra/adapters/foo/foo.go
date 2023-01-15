package foo

import (
	"context"

	pb "git.ecobin.ir/ecomicro/protobuf/foo/grpc"
	userAdapters "git.ecobin.ir/ecomicro/template/app/user/adapter"
	"git.ecobin.ir/ecomicro/template/domain"
	"git.ecobin.ir/ecomicro/tooty"
	"google.golang.org/grpc"
)

type fooAdapter struct {
	fooClient pb.FooClient
}

var _ userAdapters.FooAdapter = &fooAdapter{}

func NewFooAdapter(fooConnection *grpc.ClientConn) *fooAdapter {
	fooClient := pb.NewFooClient(fooConnection)
	return &fooAdapter{fooClient}
}
func (e *fooAdapter) Bar(ctx context.Context, user domain.User) error {
	span := tooty.OpenAnAPMSpan(ctx, "[A] bar", "adapter")
	defer tooty.CloseTheAPMSpan(span)

	_, err := e.fooClient.Bar(ctx, &pb.BarRequest{
		UserId: user.Id,
	})

	return err
}
