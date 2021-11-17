package app

import (
	"context"

	"github.com/gimmickless/keat-kit-service/internal/domain"
	"go.uber.org/zap"
)

type KitService struct {
	logger  *zap.SugaredLogger
	kitRepo IKitRepo
}

func NewKitService(
	logger *zap.SugaredLogger,
	kitRepo IKitRepo,
) *KitService {
	return &KitService{logger, kitRepo}
}

func (s *KitService) Create(ctx context.Context, catg domain.Kit) (string, error) {
	return s.kitRepo.Insert(ctx, catg)
}

func (s *KitService) Update(ctx context.Context, id string, catg domain.Kit) error {
	return s.kitRepo.Update(ctx, id, catg)
}

func (s *KitService) UpdatePrice(ctx context.Context, id string, price domain.Price) error {
	return s.kitRepo.UpdatePrice(ctx, id, price)
}

func (s *KitService) Delete(ctx context.Context, id string) error {
	return s.kitRepo.Delete(ctx, id)
}

func (s *KitService) Get(ctx context.Context, id string) (domain.Kit, error) {
	return s.kitRepo.Get(ctx, id)
}

func (s *KitService) GetAll(ctx context.Context) ([]domain.Kit, error) {
	return s.kitRepo.GetAll(ctx)
}
