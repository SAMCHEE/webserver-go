package util

// interface{} == any
type Response struct {
	Error interface{} `json:"error`
	Data  interface{} `json:"data`
}

func ResponseError(err error) *Response     { return &Response{Error: err.Error(), Data: nil} }
func ResponseOK(data interface{}) *Response { return &Response{Error: nil, Data: data} }
