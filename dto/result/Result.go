package result

type Result struct {
	Success  bool        `json:"success"`
	ErrorMsg string      `json:"erromsg"`
	Data     interface{} `json:"data"`
	Total    int64       `json:"total"`
}

type ResultOk struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ResultFail struct {
	Success  bool   `json:"success"`
	ErrorMsg string `json:"errormsg"`
}

func NewResultOk(data interface{}) *ResultOk {
	return &ResultOk{
		Success: true,
		Data:    data,
	}
}

func NewResultFail(errorMsg string) *ResultFail {
	return &ResultFail{
		Success:  false,
		ErrorMsg: errorMsg,
	}
}

func NewResult(success bool, errorMsg string, data interface{}, total int64) *Result {
	return &Result{
		Success:  success,
		ErrorMsg: errorMsg,
		Data:     data,
		Total:    total,
	}
}

func Ok() *Result {
	return NewResult(true, "", nil, 0)
}

func OkWithData(data interface{}) *Result {
	return NewResult(true, "", data, 0)
}

func OkWithListAndTotal(data []interface{}, total int64) *Result {
	return NewResult(true, "", data, total)
}

func Fail(errorMsg string) *Result {
	return NewResult(false, errorMsg, nil, 0)
}
