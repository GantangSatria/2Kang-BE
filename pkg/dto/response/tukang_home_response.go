package response

import "2Kang/internal/domain/entity"

type TukangHomeResponse struct {
	Profile entity.Tukang     `json:"profile"`
	Jobs    TukangJobsSection `json:"jobs"`
}

type TukangJobsSection struct {
	Pending   []entity.Transaction `json:"pending"`
	Ongoing   []entity.Transaction `json:"ongoing"`
	Confirmed []entity.Transaction `json:"confirmed"`
	Done      []entity.Transaction `json:"done"`
}
