package models

type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Key     string `json:"key,omitempty" example:"invalid_request"`
	Message string `json:"message,omitempty" example:"Invalid request."`
	Details string `json:"details,omitempty" example:"Invalid request."`
}
