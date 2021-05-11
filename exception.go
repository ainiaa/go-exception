package exception

import (
	"fmt"
)

var (
	SuccessCode int64 = 0
	FailureCode int64 = -1
	Success           = New(SuccessCode, "success")
	Failure           = New(FailureCode, "failure")
)

type exception struct {
	code         int64
	message      string
	subException *exception
}

type Exception interface {
	GetCode() int64
	GetSubException() Exception
	GetMessage() string
	Error() string
	SubError(err error) Exception
	NewSubError(code int64, msg string) Exception
}

func (e exception) GetCode() int64 {
	return e.code
}

func (e exception) GetMessage() string {
	return e.message
}
func (e exception) GetSubException() Exception {
	return e.subException
}

func (e exception) Error() string {
	if e.subException != nil {
		return fmt.Sprintf("code: %d, subcode: %d message: %s submessage: %s", e.code, e.subException.code, e.message, e.subException.message)
	}
	return fmt.Sprintf("code: %d message: %s ", e.code, e.message)
}

func (e exception) SubError(err error) Exception {
	switch err := err.(type) {
	case exception:
		e.subException = &err
	case *exception:
		e.subException = err
	case error:
		e.subException = &exception{
			code:         FailureCode,
			message:      err.Error(),
			subException: nil,
		}
	}

	return e
}

func (e exception) NewSubError(code int64, msg string) Exception {
	e.subException = &exception{
		code:         code,
		message:      msg,
		subException: nil,
	}
	return e
}

func New(code int64, message string) Exception {
	return exception{
		code:    code,
		message: message,
	}
}

func NewFromError(err error) Exception {
	switch err := err.(type) {
	case exception:
		return err
	case *exception:
		return err
	}
	return exception{
		code:         FailureCode,
		message:      err.Error(),
		subException: nil,
	}
}

func IsError(e interface{}) bool {
	switch e := e.(type) {
	case nil:
		return false
	case exception:
		return e.code != SuccessCode
	case error:
		return true
	}

	return false
}

/**
 * 如果 err为一个err的话，返回最终的exception，否则，返回nil
 * param: Exception mainException
 * param: error err
 * return Exception | nil
 */
func GenerateWhenError(mainException Exception, err error) Exception {
	if err != nil {
		return mainException.SubError(err)
	}
	return nil
}
