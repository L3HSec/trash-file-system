package common

type Error struct {
	info string
}

func (p *Error) Error() string {
	return p.info
}

func (p *Error) Base(err error) *Error {
	p.info = err.Error() + ">" + p.info
	return p
}

func NewError(info string) *Error {
	return &Error{
		info: info,
	}
}
