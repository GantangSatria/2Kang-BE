package request

type UpdateTukangProfileRequest struct {
	Name     *string `json:"name"`
	Bio      *string `json:"bio"`
	Services *string `json:"services"`
	Category *string `json:"category"`
}
