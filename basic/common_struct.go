package basic

type CommonError struct {
    ErrMsg  string `json:"errmsg"`
    ErrCode int64 `json:"errcode"`
}

type ResponseAccessToken struct {
    CommonError
    AccessToken string `json:"access_token"`
    ExpiresIn   int64  `json:"expires_in"`
}
