package errcode

import "fmt"

var (
	codeerrors = make(map[int]string)
)

type CodeError interface {
	Code() int
	Message() string
	Error() string
}

type Code int

func (c Code) Error() string {
	s, ok := codeerrors[c.Code()]
	if !ok {
		return "未知错误"
	}
	return s
}

func (c Code) Message() string {
	return c.Error()
}

func (c Code) Code() int {
	return int(c)
}

func (c Code) String() string {
	return fmt.Sprintf("<code:%v,message:%v>", c.Code(), c.Message())
}

func add(code int, message string) CodeError {
	if _, ok := codeerrors[code]; ok {
		panic(fmt.Sprintf("errcode: %d already exist", code))
	}
	codeerrors[code] = message
	return Code(code)
}

func New(ecode CodeError, err interface{}) CodeError {
	return DiyCode{
		CodeError: ecode,
		e:         err,
	}
}

type DiyCode struct {
	CodeError
	e interface{}
}

func (d DiyCode) Message() string {
	return fmt.Sprintf("%v,%v", d.CodeError.Message(), d.e)
}

func (d DiyCode) Error() string {
	return d.Message()
}
