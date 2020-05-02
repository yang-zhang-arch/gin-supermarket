package controller

import (
	"WebFull/common"
	"WebFull/dto"
	"WebFull/model"
	"WebFull/response"
	"WebFull/util"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Register(cont *gin.Context) {
	// 获取数据库实例
	db := common.GetDB()
	// 获取数据
	var requestUser = model.User{}
	cont.Bind(&requestUser)
	name := requestUser.Name
	tel := requestUser.Tel
	password := requestUser.Password

	// 格式验证
	if len(tel) != 11 {
		response.Fail(cont, nil, "手机号长度必须为11位")
		return
	}

	if len(password) < 6 {
		response.Fail(cont, nil, "密码不能小于6位")
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	// 后台验证
	if IsTelExist(db, tel) {
		response.Fail(cont, nil, "用户已存在")
		return
	}

	// 对密码做加密处理
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		response.Fail(cont, nil, "密码加密错误")
		return
	}
	// 添加注册
	newUser := model.User{
		Name:     name,
		Tel:      tel,
		Password: hashPassword,
	}
	db.Create(&newUser)
	response.Success(cont, nil, "注册成功")
}

func Login(cont *gin.Context) {
	// 获取数据库实例
	db := common.GetDB()
	var requestUser = model.User{}
	cont.Bind(&requestUser)
	tel := requestUser.Tel
	password := requestUser.Password

	// 格式验证
	if len(tel) != 11 {
		response.Fail(cont, nil, "手机号长度必须为11位")
		return
	}

	if len(password) < 6 {
		response.Fail(cont, nil, "密码不能小于6位")
		return
	}

	// 后台用户密码验证
	var user model.User
	db.Where("telephone = ?", tel).First(&user)

	if user.ID == 0 {
		response.Fail(cont, nil, "用户不存在")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(cont, nil, "登录失败")
		return
	}

	// 登陆成功，生成并返回token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(cont, http.StatusInternalServerError, 500, nil, "系统异常")
	}
	response.Success(cont, &gin.H{"token": token}, "登录成功")
}

func Info(cont *gin.Context) {
	user, _ := cont.Get("user")
	response.Success(cont, &gin.H{"user": dto.ToUserDTO(user.(model.User))}, "")
}

func IsTelExist(db *gorm.DB, tel string) bool {
	var user model.User
	db.Where("telephone = ?", tel).First(&user)

	if user.ID == 0 {
		return false
	}
	return true

}
