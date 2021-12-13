package myerrors

const (
	// 相应成功
	SUCCESS = 0
	// 请求消息不合理
	REQ_INFO_INVALID = 1

	// JWT相关错误
	// token生成失败
	JWT_GEN_ERR = 1000
	// token不存在
	JWT_TOKEN_NOT_EXIST = 1001
	// token解析失败
	JWT_TOKEN_MALFORMED = 1002
	// token已经过期
	JWT_TOKEN_EXPIRED = 1003
	// token未生效
	JWT_TOKEN_NOT_VALID_YET = 1004
	// token签发时间错误
	JWT_TOKEN_ISSUEAT_ERROR = 1005
	// 签名验证错误
	JWT_TOKEN_SIGNATURE_ERROR = 1006
	// 受众错误
	JWT_TOKEN_AUDIENCE_ERROR = 1007
	// 签发人错误
	JWT_TOKEN_ISSUER_ERROR = 1008
	// 其他错误
	JWT_TOKEN_OTHER_ERROR = 1009
	// Authorization错误
	JWT_TOKEN_AUTHORIZATION_ERROR = 10010
)
