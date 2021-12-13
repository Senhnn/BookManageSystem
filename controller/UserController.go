package controller

import (
	"bookmanagesystem/dao"
	"bookmanagesystem/module"
	"bookmanagesystem/myerrors"
	"bookmanagesystem/protocal"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"net/http"
)

var validate = validator.New()

func UserAdd(c *gin.Context) {
	var addUser protocal.AddUserRequest
	if err := c.ShouldBindJSON(&addUser); err != nil {
		//resp := &protouser.AddUserResponse{}
		//resp.Code = 1
		//resp.ErrMsg = "bind json err"
		//c.ProtoBuf(http.StatusBadRequest, resp)
		c.JSON(http.StatusBadRequest, &protocal.AddUserResponse{
			Code:   myerrors.REQ_INFO_INVALID,
			ErrMsg: err.Error(),
		})
		c.Abort()
		return
	}
	if err := validate.Struct(addUser); err != nil {
		c.JSON(http.StatusBadRequest, &protocal.AddUserResponse{
			Code:   myerrors.REQ_INFO_INVALID,
			ErrMsg: err.Error(),
		})
		c.Abort()
		return
	}
	ok, err := dao.AddUser(&module.User{
		Model:      gorm.Model{},
		PowerLevel: addUser.PowerLevel,
		RealName:   addUser.RealName,
		Age:        addUser.Age,
		Account:    addUser.Account,
		UserName:   addUser.UserName,
	})
	if !ok {
		fmt.Println(err)
	}
	c.JSON(http.StatusBadRequest, &protocal.AddUserResponse{
		Code: myerrors.SUCCESS,
	})
	return
}

func UserRegister(c *gin.Context) {
	var regUserReq protocal.RegUserRequest
	// 把json消息中的body绑定为结构体
	if err := c.ShouldBindJSON(&regUserReq); err != nil {
		c.JSON(http.StatusOK, &protocal.RegUserResponse{
			Code:   myerrors.REQ_INFO_INVALID,
			ErrMsg: err.Error(),
		})
		c.Abort()
		return
	}

	// 使用validate校验
	if err := validate.Struct(regUserReq); err != nil {
		c.JSON(http.StatusBadRequest, &protocal.AddUserResponse{
			Code:   myerrors.REQ_INFO_INVALID,
			ErrMsg: err.Error(),
		})
		c.Abort()
		return
	}

	// todo 此时无法在数据库中找到这个传入的account

}

func UserLogin(c *gin.Context) {

}
