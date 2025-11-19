package services

import (
	"2Kang/internal/domain/repository"
	"2Kang/pkg/dto/response"
	"errors"
)

type TukangProfileService struct {
	TukangRepo repository.TukangRepository
	ReviewRepo repository.ReviewRepository
}

func NewTukangProfileService(tukangRepo repository.TukangRepository, reviewRepo repository.ReviewRepository) *TukangProfileService {
	return &TukangProfileService{
		TukangRepo: tukangRepo,
		ReviewRepo: reviewRepo,
	}
}

func (s *TukangProfileService) GetProfile(tukangID uint) (*response.TukangProfileResponse, error) {
	t, err := s.TukangRepo.FindByID(tukangID)
	if err != nil {
		return nil, errors.New("tukang not found")
	}

	reviews, _ := s.ReviewRepo.GetReviewsByTukangID(tukangID)

	return &response.TukangProfileResponse{
		Profile: *t,
		Reviews: reviews,
	}, nil
}
