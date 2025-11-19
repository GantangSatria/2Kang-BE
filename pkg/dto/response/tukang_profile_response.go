package response

import "2Kang/internal/domain/entity"

type TukangProfileResponse struct {
	Profile entity.Tukang     `json:"profile"`
	Reviews []entity.Review   `json:"reviews"`
}
