package protocal

// 添加用户的相应
type AddUserResponse struct {
	Code   uint32 `json:"code"`
	ErrMsg string `json:"err_msg,omitempty"`
}

// 用户注册相应
type RegUserResponse struct {
	Code   uint64 `json:"code"`
	ErrMsg string `json:"err_msg,omitempty"`
}
