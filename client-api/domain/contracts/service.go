package contracts

import (
	"context"
	"github.com/JackDaniells/port-service/client-api/domain/handlers/response"
	port "github.com/JackDaniells/port-service/proto"
	"io"
)

type PortService interface {
	FindByID(ctx context.Context, id string) (*response.FindPortResponse, error)
	UploadPortFile(ctx context.Context, fileStream io.ReadCloser) (*port.CreateResponse, error)
}
