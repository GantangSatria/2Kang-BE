package services

import "2Kang/internal/domain/repository"

type TukangService struct {
	TukangRepo repository.TukangRepository
}

func NewTukangService(r repository.TukangRepository) *TukangService {
	return &TukangService{TukangRepo: r}
}

func (s *TukangService) GetTukangList(kategori string) (interface{}, error) {
	return s.TukangRepo.GetTukangList(kategori)
}

func (s *TukangService) GetTukangDetail(userID string) (interface{}, error) {
	return s.TukangRepo.GetTukangDetail(userID)
}
