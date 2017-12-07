package ajax

type ajaxResult struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Errmsg  string      `json:"errmsg"`
}

func NewAjaxResult() *ajaxResult {
	return &ajaxResult{
		Success: true,
	}
}

func (ar *ajaxResult) SetError(msg string) {
	ar.Success = false
	ar.Errmsg = msg
}
