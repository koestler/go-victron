package vedirect

import "errors"

var (
	ErrUnknownId        = errors.New("device responded with Unknown Id")
	ErrorNotSupported   = errors.New("device responded with Not Supported")
	ErrorParameterError = errors.New("device responded with Parameter Error")
)

func responseError(flag VeResponseFlag) error {
	switch flag {
	case VeResponseFlagUnknownId:
		return ErrUnknownId
	case VeResponseFlagNotSupported:
		return ErrorNotSupported
	case VeResponseFlagParameterError:
		return ErrorParameterError
	default:
		return nil
	}
}
