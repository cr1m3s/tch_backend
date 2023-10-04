package models

const (
	success = "success"
	failed  = "failed"
)

type response struct {
	Data   interface{}
	Status string
}

func NewResponseFailed(data interface{}) (r response) {
	r = response{Data: data, Status: failed}
	return r
}

func NewResponseSuccess(data interface{}) (r response) {
	r = response{Data: data, Status: success}
	return r
}
