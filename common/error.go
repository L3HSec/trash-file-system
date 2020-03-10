package common

//Error is a more conveninent error struct
type Error struct {
	info string
}

//Error implements the error interface
func (p *Error) Error() string {
	return p.info
}

//Base generates a error base on another error
func (p *Error) Base(err error) *Error {
	p.info = err.Error() + ">" + p.info
	return p
}

//NewError generates a new error
func NewError(info string) *Error {
	return &Error{
		info: info,
	}
}

//Must the err must be nil or panic
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
