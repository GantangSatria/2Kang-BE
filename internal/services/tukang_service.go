package services

import (
	"fmt"
	"2Kang/internal/domain/entity"
	"2Kang/internal/domain/repository"
)

type TukangService struct {
	TukangRepo repository.TukangRepository
}

func NewTukangService(r repository.TukangRepository) *TukangService {
	return &TukangService{TukangRepo: r}
}

func (s *TukangService) GetTukangList(kategori string) ([]entity.Tukang, error) {
	list, err := s.TukangRepo.GetTukangList(kategori)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tukang list: %w", err)
	}
	return list, nil
}

func (s *TukangService) GetTukangDetail(id uint) (*entity.Tukang, error) {
	if id == 0 {
		return nil, fmt.Errorf("invalid tukang id")
	}

	data, err := s.TukangRepo.GetTukangDetail(id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tukang detail: %w", err)
	}
	return data, nil
}
