package services

import (
	"context"
	"github.com/JackDaniells/port-service/client-api/domain/contracts"
	"github.com/JackDaniells/port-service/client-api/domain/handlers/request"
	"github.com/JackDaniells/port-service/client-api/domain/handlers/response"
)

type portService struct {
	grpcClient contracts.PortServiceClient
}

func NewPortService(grpcClient contracts.PortServiceClient) *portService {
	return &portService{grpcClient: grpcClient}
}

func (p *portService) FindByID(ctx context.Context, id string) (*response.FindPortResponse, error) {
	protoPort, err := p.grpcClient.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return response.NewFindPortResponse(protoPort), nil
}

func (p *portService) StreamCreate(ctx context.Context, ports request.UploadPortByFileRequest) (*response.UploadPortByFileResponse, error) {

	portsDomain := request.ParseUploadPortRequestToProtoPortArray(ports)

	resp, err := p.grpcClient.StreamCreate(ctx, portsDomain)
	if err != nil {
		return nil, err
	}

	return response.NewUploadPortByFileResponse(resp), nil
}
