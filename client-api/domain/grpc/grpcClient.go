package grpc

import (
	"context"
	"fmt"
	proto "github.com/JackDaniells/port-service/proto"
	"google.golang.org/grpc"
)

type PortServiceClient struct {
	client proto.PortServiceClient
}

func NewPortServiceClient(conn grpc.ClientConnInterface) PortServiceClient {
	protoClient := proto.NewPortServiceClient(conn)

	return PortServiceClient{
		client: protoClient,
	}
}

func (p PortServiceClient) FindByID(ctx context.Context, id string) (*proto.Port, error) {
	request := &proto.FindByIDRequest{
		Id: id,
	}

	response, err := p.client.FindByID(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("find by id: %w", err)
	}

	return response, nil
}

func (p PortServiceClient) StreamCreate(ctx context.Context, ports []*proto.Port) (*proto.CreateResponse, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	stream, err := p.client.Create(ctx)
	if err != nil {
		return nil, fmt.Errorf("create stream: %w", err)
	}

	for _, port := range ports {
		if err := stream.Send(port); err != nil {
			return nil, fmt.Errorf("send stream: %w", err)
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		return nil, fmt.Errorf("close and receive: %w", err)
	}

	return response, nil
}
