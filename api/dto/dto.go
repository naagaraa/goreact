package dto

import (
    "encoding/json"
    "time"
)

// ResponseDTO represents the structure of the API response
type ResponseDTO struct {
    Status        bool        `json:"status"`
    Time          time.Time   `json:"time"`
    ResponseCode  int         `json:"response_code"`
    Message       string      `json:"message"`
    Data          interface{} `json:"data,omitempty"`
    Error         string      `json:"error,omitempty"`
}

// NewResponseDTO creates a new instance of ResponseDTO
func NewResponseDTO(status bool, code int, message string, data interface{}) *ResponseDTO {
    return &ResponseDTO{
        Status:        status,
        Time:          time.Now(),
        ResponseCode:  code,
        Message:       message,
        Data:          data,
        Error:         "",
    }
}

// ToMap converts ResponseDTO to a map
func (r *ResponseDTO) ToMap() map[string]interface{} {
    result := map[string]interface{}{
        "status":        r.Status,
        "time":          r.Time.Format(time.RFC3339),
        "response_code": r.ResponseCode,
        "message":       r.Message,
        "data":          r.Data,
    }
    if !r.Status {
        result["error"] = r.Message
    }
    return result
}

// ToJSON converts ResponseDTO to JSON
func (r *ResponseDTO) ToJSON() (string, error) {
    b, err := json.Marshal(r.ToMap())
    if err != nil {
        return "", err
    }
    return string(b), nil
}
