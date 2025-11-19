package services

import (
	"2Kang/internal/domain/entity"
	"2Kang/internal/domain/repository"
	"2Kang/pkg/dto/response"
)

type TukangHomeService struct {
	TukangRepo      repository.TukangRepository
	TransactionRepo repository.TransactionRepository
}

func NewTukangHomeService(
	tukangRepo repository.TukangRepository,
	transactionRepo repository.TransactionRepository,
) *TukangHomeService {
	return &TukangHomeService{
		TukangRepo:      tukangRepo,
		TransactionRepo: transactionRepo,
	}
}

func (s *TukangHomeService) GetHome(tukangID uint) (*response.TukangHomeResponse, error) {

	// Get profile
	profile, err := s.TukangRepo.FindByID(tukangID)
	if err != nil {
		return nil, err
	}

	// Get transactions by status
	pending, _ := s.TransactionRepo.GetByTukangAndStatus(tukangID, entity.StatusPending)
	ongoing, _ := s.TransactionRepo.GetByTukangAndStatus(tukangID, entity.StatusOngoing)
	confirmed, _ := s.TransactionRepo.GetByTukangAndStatus(tukangID, entity.StatusConfirmed)
	done, _ := s.TransactionRepo.GetByTukangAndStatus(tukangID, entity.StatusDone)

	// Build response
	res := &response.TukangHomeResponse{
		Profile: *profile,
		Jobs: response.TukangJobsSection{
			Pending:   pending,
			Ongoing:   ongoing,
			Confirmed: confirmed,
			Done:      done,
		},
	}

	return res, nil
}
