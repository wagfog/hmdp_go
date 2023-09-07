package result

type Result struct {
	Success  bool
	ErrorMsg string
	Data     interface{}
	Total    int64
}

func NewResult(success bool, errorMsg string, data interface{}, total int64) *Result {
	return &Result{
		Success:  success,
		ErrorMsg: errorMsg,
		Data:     data,
		Total:    total,
	}
}

func (result Result) ok() *Result {
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
