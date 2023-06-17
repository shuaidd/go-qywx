package api

type BaseResponse struct {
	ErrCode    string `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	NextCursor string `json:"next_cursor"`
	HasMore    int    `json:"has_more"`
}
