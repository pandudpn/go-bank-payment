package bankpayment

import "errors"

var (
	// ErrInternal used if there's an error internal, e.g: unmarshal or marshal
	ErrInternal = errors.New("internal server error")
	// ErrBadRequest contains data request not completed or filled
	ErrBadRequest = errors.New("bad request")
	// ErrUnauthorized is not authenticated
	ErrUnauthorized = errors.New("unauthorized access")
	// ErrForbidden is not authorized to access a resource
	ErrForbidden = errors.New("access forbidden")
)
