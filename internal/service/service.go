// Package service 实现业务逻辑层。
package service

import (
	"readinglog/internal/config"
	"readinglog/internal/scopepool"
	"readinglog/internal/store"
	"readinglog/pkg/logger"
)

// Service 业务服务，聚合各实体业务方法。
type Service struct {
	store      store.Store
	log        *logger.Logger
	cfg        *config.Config
	identities *scopepool.IdentityPool
}

// New 创建业务服务。
func New(st store.Store, log *logger.Logger, cfg *config.Config) *Service {
	return &Service{store: st, log: log, cfg: cfg, identities: scopepool.New()}
}
