package dto

type ScrollResult struct {
	List    []interface{} `json:"list"`
	MinTime int64         `json:"minTime"`
	Offset  int           `json:"offset"`
}
