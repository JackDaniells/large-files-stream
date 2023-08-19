package contracts

import (
	"context"
	proto "github.com/JackDaniells/port-service/proto"
)

type PortService interface {
	Find(ctx context.Context, id string) (*proto.Port, error)
	Store(ctx context.Context, port *proto.Port) error
}
