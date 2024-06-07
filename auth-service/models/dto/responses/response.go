package responses

/**
 * created by	: albym
 * created at	: 5 Jun 2024
 *
 * return response type
 */
type Response struct {
	Message string
	Body    interface{}
}

func NewResponse(message string, body interface{}) *Response {
	return &Response{
		Message: message,
		Body:    body,
	}
}
