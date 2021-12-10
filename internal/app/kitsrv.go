package app

import (
	"context"

	"github.com/gimmickless/keat-kit-service/internal/domain"
	"github.com/gimmickless/keat-kit-service/pkg/enum"
	"go.uber.org/zap"
)

type IKitService interface {
	Create(ctx context.Context, catg domain.Kit) (string, error)
	Update(ctx context.Context, id string, catg domain.Kit) error
	UpdatePrice(ctx context.Context, id string, price domain.Price) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (domain.Kit, error)
	GetPage(
		ctx context.Context, page int, pageSize int, sortField string, sortDirection enum.SortDirection,
	) ([]domain.Kit, error)
	GetAll(ctx context.Context) ([]domain.Kit, error)
}

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

func (s *KitService) UpdatePrice(ctx context.Context, id string, prices []domain.Price) error {
	return s.kitRepo.UpdatePrice(ctx, id, prices)
}

func (s *KitService) Delete(ctx context.Context, id string) error {
	return s.kitRepo.Delete(ctx, id)
}

func (s *KitService) Get(ctx context.Context, id string) (domain.Kit, error) {
	return s.kitRepo.Get(ctx, id)
}

func (s *KitService) GetPage(
	ctx context.Context, page int, pageSize int, sortField string, sortDirection enum.SortDirection,
) ([]domain.Kit, error) {
	return s.kitRepo.GetPaginated(ctx, pageSize, page-1, sortField, sortDirection)
}

func (s *KitService) GetAll(ctx context.Context) ([]domain.Kit, error) {
	return s.kitRepo.GetAll(ctx)
}
