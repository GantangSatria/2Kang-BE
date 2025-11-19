package request

type UpdateUserNameRequest struct {
    Name string `json:"name" validate:"required,min=3"`
}
