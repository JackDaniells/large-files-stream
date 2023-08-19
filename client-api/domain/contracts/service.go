package contracts

import (
	"context"
	"github.com/JackDaniells/port-service/client-api/domain/handlers/request"
	"github.com/JackDaniells/port-service/client-api/domain/handlers/response"
)

type PortService interface {
	FindByID(ctx context.Context, id string) (*response.FindPortResponse, error)
	StreamCreate(ctx context.Context, ports request.UploadPortByFileRequest) (*response.UploadPortByFileResponse, error)
}
