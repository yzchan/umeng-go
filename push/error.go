package push

type UmengError struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func (e *UmengError) Error() string {
	return e.ErrorMsg
}

func NewUmengError(code string, msg string) error {
	return &UmengError{
		ErrorCode: code,
		ErrorMsg:  msg,
	}
}

type UmengErrorResp struct {
	Ret  string      `json:"ret"`
	Data *UmengError `json:"data"`
}
