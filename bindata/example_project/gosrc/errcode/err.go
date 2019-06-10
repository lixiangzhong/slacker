package errcode

import "fmt"

var (
	codeerrors = make(map[int]string)
)

type CodeError int

func (c CodeError) Error() string {
	s, ok := codeerrors[c.Code()]
	if !ok {
		return "未知错误"
	}
	return s
}

func (c CodeError) Message() string {
	return c.Error()
}

func (c CodeError) Code() int {
	return int(c)
}

func (c CodeError) String() string {
	return fmt.Sprintf("<code:%v,message:%v>", c.Code(), c.Message())
}

func add(code int, message string) CodeError {
	if _, ok := codeerrors[code]; ok {
		panic(fmt.Sprintf("errcode: %d already exist", code))
	}
	codeerrors[code] = message
	return CodeError(code)
}
