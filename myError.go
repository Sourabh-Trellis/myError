package myerror

import "fmt"

type CustomError struct {
	Status  bool
	Message string
	Code    int
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Status:%v\nMessage:%v\nCode:%v", e.status, e.message, e.code)
}
