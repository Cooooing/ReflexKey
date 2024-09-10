package model

import "time"

// Result represents a common-used result struct.
type Result struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // message
	Data any    `json:"data"` // data object
	Time string `json:"time"` // time
}

// NewResult creates a result with Code=0, Msg="", Data=nil.
func NewResult() *Result {
	return &Result{
		Code: 0,
		Msg:  "",
		Data: nil,
		Time: time.Now().Format("2006-01-02 15:04:05"),
	}
}

func Fail(msg string) *Result {
	return &Result{
		Code: 1,
		Msg:  msg,
		Data: nil,
		Time: time.Now().Format("2006-01-02 15:04:05"),
	}
}
func Success(data any) *Result {
	return &Result{
		Code: 0,
		Msg:  "",
		Data: data,
		Time: time.Now().Format("2006-01-02 15:04:05"),
	}
}
