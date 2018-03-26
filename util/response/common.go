package response

type CommonError struct {
    ErrorCode int64 `json:"errcode"`
    ErrorMessage string `json:"errmsg"`
}