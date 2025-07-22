package presenter

import (
	"ticket-system/api/dto"
)

func ErrorResponse(err error) *dto.ErrorResponseDTO {
	return &dto.ErrorResponseDTO{Error: err.Error()}
}
