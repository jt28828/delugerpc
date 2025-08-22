package jsonrpc

import "fmt"

// V1Request is the JSON-RPC v1 request structure.
type V1Request[T any] struct {
	Method string `json:"method"`
	Params []T    `json:"params"`
	Id     int    `json:"id"`
}

func NewV1Request[T any](method string, params ...T) V1Request[T] {
	if params == nil {
		params = make([]T, 0)
	}

	return V1Request[T]{
		Method: method,
		Params: params,
		Id:     1,
	}
}

// V1Error is the JSON-RPC v1 error structure.
type V1Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *V1Error) Error() string {
	return fmt.Sprintf("code: %d, message: %s", err.Code, err.Message)
}

func (err *V1Error) IsNotAuthenticated() bool {
	return err.Message == "Not authenticated"
}

// V1Response is the JSON-RPC v1 response structure.
type V1Response[T any] struct {
	Result *T       `json:"result,omitempty"`
	Error  *V1Error `json:"error,omitempty"`
	Id     int      `json:"id,omitempty"`
}
