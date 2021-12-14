package protocal

// 添加用户的响应
type AddUserResponse struct {
	Code   uint32 `json:"code"`
	ErrMsg string `json:"err_msg,omitempty"`
}

// 用户注册响应
type RegUserResponse struct {
	Code   uint64 `json:"code"`
	ErrMsg string `json:"err_msg,omitempty"`
}

// 中间件验证失败时的响应
type MiddleWareAuthErrResponse struct {
	Type   string `json:"type"`
	Code   uint64 `json:"code"`
	ErrMsg string `json:"err_msg,omitempty"`
}

// 用户登录响应
type UserLoginResponse struct {
	Code   uint64 `json:"code"`
	ErrMsg string `json:"err_msg"`
}
