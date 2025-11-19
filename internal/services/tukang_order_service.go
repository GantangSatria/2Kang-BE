package services

import (
	"2Kang/internal/domain/entity"
	"2Kang/internal/domain/repository"
	"errors"
)

type TukangOrderService struct {
	TransactionRepo repository.TransactionRepository
}

func NewTukangOrderService(
	transactionRepo repository.TransactionRepository,
) *TukangOrderService {
	return &TukangOrderService{
		TransactionRepo: transactionRepo,
	}
}

func (s *TukangOrderService) AcceptOrder(orderID uint, tukangID uint) error {
	tx, err := s.TransactionRepo.GetByID(orderID)
	if err != nil {
		return err
	}

	if tx.TukangID != tukangID {
		return errors.New("not your order")
	}

	if tx.Status != entity.StatusPending {
		return errors.New("order cannot be accepted")
	}

	return s.TransactionRepo.UpdateStatus(orderID, entity.StatusConfirmed)
}

func (s *TukangOrderService) StartOrder(orderID uint, tukangID uint) error {
	tx, err := s.TransactionRepo.GetByID(orderID)
	if err != nil {
		return err
	}

	if tx.TukangID != tukangID {
		return errors.New("not your order")
	}

	if tx.Status != entity.StatusConfirmed {
		return errors.New("order cannot be started")
	}

	return s.TransactionRepo.UpdateStatus(orderID, entity.StatusOngoing)
}

func (s *TukangOrderService) FinishOrder(orderID uint, tukangID uint) error {
	tx, err := s.TransactionRepo.GetByID(orderID)
	if err != nil {
		return err
	}

	if tx.TukangID != tukangID {
		return errors.New("not your order")
	}

	if tx.Status != entity.StatusOngoing {
		return errors.New("order cannot be finished")
	}

	return s.TransactionRepo.UpdateStatus(orderID, entity.StatusDone)
}
