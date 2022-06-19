package errorcode

type ServiceError struct {
	errmsg string
	code   int32
}

const (
	ErrCodeCommon      = 1000
	ErrCodeInvalidArgs = 1001
	ErrCodeRequestBusy = 1002
	ErrCodeServerBusy  = 1010

	ErrCodeNotAuthenticated      = 1011
	ErrCodeAuthenticationExpired = 1012
	ErrCodeWrongSig              = 1013

	ErrCodeTextFormat   = 1020
	ErrCodeRenameNum    = 1021
	ErrCodeNameRepeated = 1022

	ErrCodeChangePlanetInfo      = 1030
	ErrCodeYourAreNotOwner       = 1031
	ErrCodeChangePlanetTimeLimit = 1032
)

var (
	ErrServerBusy    = NewError(1010, "server busy")
	ErrOperateFailed = NewError(ErrCodeCommon)
)

func NewError(code int32, msg ...string) *ServiceError {
	if len(msg) > 0 {
		return &ServiceError{
			code:   code,
			errmsg: msg[0],
		}
	}

	return &ServiceError{
		code: code,
	}
}

func (s *ServiceError) Error() string {
	return s.errmsg
}

func (s *ServiceError) Code() int32 {
	return s.code
}
