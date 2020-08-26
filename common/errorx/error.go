package errorx

type (
	ErrorWithCode struct {
		code int
		message string
	}
)

const unknownCode = -9999

func (e *ErrorWithCode) Error() string {
	return e.message
}

func (e *ErrorWithCode) Code() int {
	return e.code
}

func NewErrorWithCode(code int, message string) *ErrorWithCode {
	return &ErrorWithCode{
		code: code,
		message: message,
	}
}

func ParseError(err error) *ErrorWithCode {
	if e, ok := err.(*ErrorWithCode); ok {
		return e
	} else {
		return NewErrorWithCode(unknownCode, err.Error())
	}
}

