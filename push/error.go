package push

import "fmt"

type UmengError struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func (e *UmengError) Error() string {
	return fmt.Sprintf("Umeng ErrCode=%d ErrMsg=%s", e.ErrorCode, e.ErrorMsg)
}

func NewUmengError(code int, msg string) error {
	return &UmengError{
		ErrorCode: code,
		ErrorMsg:  msg,
	}
}

type UmengErrorResp struct {
	Ret  string      `json:"ret"`
	Data *UmengError `json:"data"`
}
