package middleware

import (
	"bookmanagesystem/config"
	"bookmanagesystem/myerrors"
	"bookmanagesystem/protocal"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// 标准的claim，v1版本
type StandardClaimsV1 struct {
	Username string `json:"username"`
	Account  string `json:"account"`
	jwt.StandardClaims
}

// 生成token
func GenerateToken(username, account string) (string, uint64) {
	expireTime := time.Now().Add(5 * time.Hour)
	claim := StandardClaimsV1{
		username,
		account,
		jwt.StandardClaims{
			// 受众
			Audience: "user",
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 标号
			Id: "",
			// 签发时间
			IssuedAt: time.Now().Unix(),
			// 签发人
			Issuer: config.JwtIss,
			// 生效时间
			NotBefore: time.Now().Unix(),
			// 主题
			Subject: "",
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(config.JwtSigKey)
	if err != nil {
		return "", myerrors.JWT_GEN_ERR
	}
	return token, myerrors.SUCCESS
}

// 验证token
func CheckToken(token string) (*StandardClaimsV1, uint64) {
	var claim StandardClaimsV1
	// 解析token字符串
	jwtToken, err := jwt.ParseWithClaims(token, &claim, func(token *jwt.Token) (interface{}, error) {
		return config.JwtSigKey, nil
	})
	if err != nil {
		// jwt.ValidationError 是一个无效token的错误结构
		if value, ok := err.(*jwt.ValidationError); ok {
			if value.Errors&jwt.ValidationErrorMalformed /*token解析畸形*/ != 0 {
				return nil, myerrors.JWT_TOKEN_MALFORMED
			} else if value.Errors&jwt.ValidationErrorExpired /*token过期*/ != 0 {
				return nil, myerrors.JWT_TOKEN_EXPIRED
			} else if value.Errors&jwt.ValidationErrorNotValidYet /*token未生效*/ != 0 {
				return nil, myerrors.JWT_TOKEN_NOT_VALID_YET
			} else if value.Errors&jwt.ValidationErrorIssuedAt /*token签发时间错误*/ != 0 {
				return nil, myerrors.JWT_TOKEN_ISSUEAT_ERROR
			} else if value.Errors&jwt.ValidationErrorSignatureInvalid /*token签名验证错误*/ != 0 {
				return nil, myerrors.JWT_TOKEN_SIGNATURE_ERROR
			} else if value.Errors&jwt.ValidationErrorAudience /*token AUD错误*/ != 0 {
				return nil, myerrors.JWT_TOKEN_AUDIENCE_ERROR
			} else if value.Errors&jwt.ValidationErrorIssuer /*token 签发人错误*/ != 0 {
				return nil, myerrors.JWT_TOKEN_ISSUER_ERROR
			} else {
				return nil, myerrors.JWT_TOKEN_OTHER_ERROR
			}
		}
	}
	// 把claim接口复原
	if key, ok := jwtToken.Claims.(*StandardClaimsV1); ok && jwtToken.Valid {
		return key, myerrors.SUCCESS
	}
	return nil, myerrors.JWT_TOKEN_OTHER_ERROR
}

// jwt中间件
func JwtTokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中
		Authorization := c.Request.Header.Get("Authorization")
		if Authorization == "" {
			c.JSON(http.StatusOK, &protocal.MiddleWareAuthErrResponse{
				Type:   "JWT",
				Code:   myerrors.JWT_TOKEN_NOT_EXIST,
				ErrMsg: "",
			})
			c.Abort()
			return
		}

		// 分割Authorization的内容
		// Authorization，用Bearer schema: Authorization: Bearer <token>
		splitAuthorization := strings.SplitN(Authorization, " ", 2)
		if len(splitAuthorization) != 2 || splitAuthorization[0] != "Bearer" {
			c.JSON(http.StatusOK, &protocal.MiddleWareAuthErrResponse{
				Type:   "JWT",
				Code:   myerrors.JWT_TOKEN_AUTHORIZATION_ERROR,
				ErrMsg: "",
			})
			c.Abort()
			return
		}

		// splitAuthorization[1]是token的字符串
		tokenClaim, errCode := CheckToken(splitAuthorization[1])
		if errCode != 0 {
			c.JSON(http.StatusOK, &protocal.MiddleWareAuthErrResponse{
				Type:   "JWT",
				Code:   errCode,
				ErrMsg: "",
			})
			c.Abort()
			return
		}
		c.Set("username", tokenClaim.Username)
		c.Set("account", tokenClaim.Account)
		c.Next()
	}
}
