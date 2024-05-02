package dto

type SuccessResponse struct {
	statusCode int
	message    string
	data       interface{}
}

func NewSuccessResponse(statusCode int, message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		statusCode: statusCode,
		message:    message,
		data:       data,
	}
}

// GetData implements OKRes
func (s *SuccessResponse) GetModel() *SuccessResponse {
	return &SuccessResponse{
		statusCode: s.statusCode,
		message:    s.message,
		data:       s.data,
	}
}
