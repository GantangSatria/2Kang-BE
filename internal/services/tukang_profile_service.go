package services

import (
	"2Kang/internal/domain/repository"
	"2Kang/pkg/dto/request"
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

func (s *TukangProfileService) UpdateProfile(tukangID uint, req request.UpdateTukangProfileRequest) error {

	t, err := s.TukangRepo.FindByID(tukangID)
	if err != nil {
		return err
	}

	// Update hanya jika dikirim
	if req.Name != nil {
		t.Name = *req.Name
	}
	if req.Bio != nil {
		t.Bio = *req.Bio
	}
	if req.Services != nil {
		t.Services = *req.Services
	}
	if req.Category != nil {
		t.Category = *req.Category
	}

	return s.TukangRepo.Update(t)
}