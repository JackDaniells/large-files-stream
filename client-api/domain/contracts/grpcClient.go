package contracts

import (
	"context"
	proto "github.com/JackDaniells/port-service/proto"
)

type PortServiceClient interface {
	FindByID(ctx context.Context, id string) (*proto.Port, error)
	StreamCreate(ctx context.Context, ports []*proto.Port) (*proto.CreateResponse, error)
}
