package main

import (
	"errors"
	"fmt"

	"github.com/ainiaa/go-exception"
)

func main() {
	e := exception.New(-1,"failure")
	e2 :=exception.New(-1,"sub")
	e3 := exception.GenerateWhenError(e,e2)
	fmt.Printf("e:%s\n", e)
	fmt.Printf("e3:%s\n", e3)
	fmt.Printf("error code:%d\n", e.GetCode())
	fmt.Printf("error hasException:%v\n", exception.HasException(e))

	err := errors.New("new error")
	e4 := e.SubError(err)
	fmt.Printf("e1:%s\n", e)
	fmt.Printf("e2:%s\n", e4)
}