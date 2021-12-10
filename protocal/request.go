package protocal

// 添加用户的请求
type AddUserRequest struct {
	PowerLevel uint8  `json:"power_level" validate:"gt=0"`
	RealName   string `json:"real_name" validate:"gt=0"`
	Age        uint16 `json:"age"`
	Account    string `json:"account"`
	UserName   string `json:"user_name"`
}

// 用户注册请求
type RegUserRequest struct {
	Account  string `json:"account" validate:"gte=6,required,lte=12"`
	Password string `json:"password" validate:"gte=6,required,lte=12"`
	UserName string `json:"user_name" validate:"lte=20,gte=1"`
	Age      string `json:"age" validate:"numeric"`
}
