package services

import (
	"context"
	"errors"
	"github.com/JackDaniells/port-service/client-api/domain/contracts"
	"github.com/JackDaniells/port-service/client-api/domain/contracts/mocks"
	"github.com/JackDaniells/port-service/client-api/domain/handlers/request"
	"github.com/JackDaniells/port-service/client-api/domain/handlers/response"
	proto "github.com/JackDaniells/port-service/proto"
	"reflect"
	"testing"
)

func Test_portService_FindByID(t *testing.T) {
	type fields struct {
		grpcClient contracts.PortServiceClient
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.FindPortResponse
		wantErr bool
	}{
		{
			name: "Should return port entity",
			fields: fields{
				grpcClient: func() contracts.PortServiceClient {
					client := &mocks.PortServiceClient{}
					client.On("FindByID", context.Background(), "123").
						Return(&proto.Port{Id: "123"}, nil)

					return client
				}(),
			},
			args: args{
				ctx: context.Background(),
				id:  "123",
			},
			want:    &response.FindPortResponse{ID: "123"},
			wantErr: false,
		},
		{
			name: "Should return error when grpc client return error",
			fields: fields{
				grpcClient: func() contracts.PortServiceClient {
					client := &mocks.PortServiceClient{}
					client.On("FindByID", context.Background(), "123").
						Return(nil, errors.New("error"))

					return client
				}(),
			},
			args: args{
				ctx: context.Background(),
				id:  "123",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &portService{
				grpcClient: tt.fields.grpcClient,
			}
			got, err := p.FindByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_portService_StreamCreate(t *testing.T) {
	type fields struct {
		grpcClient contracts.PortServiceClient
	}
	type args struct {
		ctx   context.Context
		ports request.UploadPortByFileRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.UploadPortByFileResponse
		wantErr bool
	}{
		{
			name: "Should create ports with success",
			fields: fields{
				grpcClient: func() contracts.PortServiceClient {
					client := &mocks.PortServiceClient{}
					client.On("StreamCreate", context.Background(), []*proto.Port{
						{Id: "PORT", Name: "Port"},
					}).Return(&proto.CreateResponse{Total: 1}, nil)

					return client
				}(),
			},
			args: args{
				ctx: context.Background(),
				ports: request.UploadPortByFileRequest{
					"PORT": request.CreatePortRequest{
						Name: "Port",
					},
				},
			},
			want: &response.UploadPortByFileResponse{
				Total: 1,
			},
			wantErr: false,
		},
		{
			name: "Should return error when grpc client return error",
			fields: fields{
				grpcClient: func() contracts.PortServiceClient {
					client := &mocks.PortServiceClient{}
					client.On("StreamCreate", context.Background(), []*proto.Port{
						{Id: "PORT", Name: "Port"},
					}).Return(nil, errors.New("error"))

					return client
				}(),
			},
			args: args{
				ctx: context.Background(),
				ports: request.UploadPortByFileRequest{
					"PORT": request.CreatePortRequest{
						Name: "Port",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &portService{
				grpcClient: tt.fields.grpcClient,
			}

			got, err := p.StreamCreate(tt.args.ctx, tt.args.ports)
			if (err != nil) != tt.wantErr {
				t.Errorf("StreamCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StreamCreate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
