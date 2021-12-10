package controller

import (
	"bookmanagesystem/dao"
	"bookmanagesystem/module"
	"bookmanagesystem/protocal"
	"bookmanagesystem/respcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"net/http"
)

var validate = validator.New()

func AddUser(c *gin.Context) {
	var user protocal.AddUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		//resp := &protouser.AddUserResponse{}
		//resp.Code = 1
		//resp.ErrMsg = "bind json err"
		//c.ProtoBuf(http.StatusBadRequest, resp)
		c.JSON(http.StatusBadRequest, &protocal.AddUserResponse{
			Code:   respcode.REQ_INFO_INVALID,
			ErrMsg: err.Error(),
		})
		c.Abort()
		return
	}
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, &protocal.AddUserResponse{
			Code:   respcode.REQ_INFO_INVALID,
			ErrMsg: err.Error(),
		})
		c.Abort()
		return
	}
	ok, err := dao.AddUser(&module.User{
		Model:      gorm.Model{},
		PowerLevel: user.PowerLevel,
		RealName:   user.RealName,
		Age:        user.Age,
		Account:    user.Account,
		UserName:   user.UserName,
	})
	if !ok {
		fmt.Println(err)
	}
	c.JSON(http.StatusBadRequest, &protocal.AddUserResponse{
		Code: respcode.SUCCESS,
	})
	return
}
