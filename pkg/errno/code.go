package errno

import (
	"errors"
	"fmt"
)
//定义错误码
type ErrNo struct {
	Code int
	Message  string
}
//定义错误
type Err struct {
	Code int // 错误码
	Message string //展示给用户看的
	Errord error //保存内部信息错误
}

//定义HttpErr
type HttpErr struct {
	StatusCode int
	Errno ErrNo
}


func (e ErrNo)Error()string  {
	return fmt.Sprintf("err_code = %d,err_msg = %s",e.Code,e.Message)
}
func (e ErrNo) WithMessage(msg string)ErrNo  {
	return ErrNo{
		Message: msg,
	}
}

//新的ErrNo
func NewErrNo(code int,msg string)ErrNo  {
	return ErrNo{code,msg}
}
//新的HttpErrNo
func NewHttpErr(code int,httpCode int,msg string) HttpErr  {
	return HttpErr{
		StatusCode: httpCode,
		Errno: ErrNo{
			Code: code,
			Message: msg,
		},
	}
}


func (e *Err)Error()string{
	return fmt.Sprintf("err_code = %d,err_msg = %s,error = %s",e.Code,e.Message,e.Errord)
}
//新建错误
func New(errno *ErrNo,err error)*Err  {
	return &Err{
		Code: errno.Code,
		Message: errno.Message,
		Errord: err,
	}
}
//解码错误,获取Code和Message
func DecodeErr(err error)(int,string)  {
	if err == nil{
		return Success.Code, Success.Message
	}
	switch typed := err.(type) {
	case *Err:
		return typed.Code,typed.Message
	case *ErrNo:
		return typed.Code,typed.Message
	default:
	}
	return ErrUnknown.Code,ErrUnknown.Message
}

//ConvertErr convert error to Errno
func ConvertErr(err error)ErrNo  {
	Err := ErrNo{}
	if errors.As(err,&Err) {
		return Err
	}
	s := ErrUnknown
	s.Message = err.Error()
	return s
}

