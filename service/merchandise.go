package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"productCatalog/ent"
)

type IMerchandiseService interface {
	CreateMerchandise() error
}

type MerchandiseService struct {
	client *ent.Client
	cache  *redis.Client
	ctx    context.Context
}

func NewMerchandiseService(client *ent.Client, cache *redis.Client) IMerchandiseService {
	return &MerchandiseService{
		client: client,
		cache:  cache,
		ctx:    context.Background(),
	}
}

func (s MerchandiseService) CreateMerchandise() error {
	return nil
}
