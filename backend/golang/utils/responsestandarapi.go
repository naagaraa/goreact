package utils

import "goreact/backend/golang/dto"

// ResponseStandart generates a standardized response
func ResponseStandart(status bool, code int, message string, data interface{}) (*dto.ResponseDTO, error) {
    responseDTO := dto.NewResponseDTO(status, code, message, data)
    return responseDTO, nil
}
