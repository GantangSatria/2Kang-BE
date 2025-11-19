package request

type UpdateTukangCategoryRequest struct {
    Category string `json:"category" validate:"required"`
}

type UpdateTukangBioRequest struct {
    Bio string `json:"bio"`
}

type UpdateTukangServicesRequest struct {
    Services string `json:"services"`
}
