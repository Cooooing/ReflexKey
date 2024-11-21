package model

import "time"

// Result represents a common-used result struct.
type Result struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // message
	Data any    `json:"data"` // data object
	Time string `json:"time"` // time
}

func Fail(msg string) *Result {
	return &Result{
		Code: 500,
		Msg:  msg,
		Time: time.Now().Format("2006-01-02 15:04:05"),
	}
}

func Success(data any) *Result {
	return &Result{
		Code: 200,
		Data: data,
		Time: time.Now().Format("2006-01-02 15:04:05"),
	}
}

func BadRequest() *Result {
	return &Result{
		Code: 400,
		Data: "Bad Request",
		Time: time.Now().Format("2006-01-02 15:04:05"),
	}
}
