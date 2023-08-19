package repository

import (
	"context"
	"fmt"
	"github.com/JackDaniells/port-service/server/domain/entity"
	"github.com/JackDaniells/port-service/server/infra/redis"
)

type PortRepository struct {
	client redis.RedisClient
}

func NewPortRepository(client redis.RedisClient) *PortRepository {
	return &PortRepository{client: client}
}

func (r *PortRepository) Find(ctx context.Context, id string) (*entity.Port, error) {
	entity := &entity.Port{}
	err := r.client.Find(id, entity)
	if err != nil {
		return nil, fmt.Errorf("find port: %w", err)
	}

	return entity, nil

}

func (r *PortRepository) Store(ctx context.Context, domain *entity.Port) error {
	err := r.client.Store(domain.ID, domain, 0)
	if err != nil {
		return fmt.Errorf("store port: %w", err)
	}

	return nil
}
