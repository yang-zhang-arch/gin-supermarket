package controller

import (
	"WebFull/common"
	"WebFull/dto"
	"WebFull/model"
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
		cont.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    "422",
			"message": "手机号长度必须为11位",
		})
		return
	}

	if len(password) < 6 {
		cont.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    "423",
			"message": "密码不能小于6位",
		})
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	// 后台验证
	if IsTelExist(db, tel) {
		cont.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    "424",
			"message": "用户已存在",
		})
		return
	}

	// 对密码做加密处理
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		cont.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    "500",
			"message": "密码加密错误",
		})
		return
	}
	// 添加注册
	newUser := model.User{
		Name:     name,
		Tel:      tel,
		Password: hashPassword,
	}
	db.Create(&newUser)
	cont.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "注册成功",
	})
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
		cont.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    "422",
			"message": "手机号长度必须为11位",
		})
		return
	}

	if len(password) < 6 {
		cont.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    "423",
			"message": "密码不能小于6位",
		})
		return
	}

	// 后台用户密码验证
	var user model.User
	db.Where("telephone = ?", tel).First(&user)

	if user.ID == 0 {
		cont.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    "422",
			"message": "手机号长度必须为11位",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		cont.JSON(http.StatusBadRequest, gin.H{
			"code":    "401",
			"message": "登录失败",
		})
	}

	// 登陆成功，生成并返回token
	token, err := common.ReleaseToken(user)
	if err != nil {
		cont.JSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "系统异常",
		})
	}
	cont.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"code":    "200",
		"data":    gin.H{"token": token},
	})

}

func Info(cont *gin.Context) {
	user, _ := cont.Get("user")
	cont.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": gin.H{"user": dto.ToUserDTO(user.(model.User))}, // 类型转换
	})

}

func IsTelExist(db *gorm.DB, tel string) bool {
	var user model.User
	db.Where("telephone = ?", tel).First(&user)

	if user.ID == 0 {
		return false
	}
	return true

}
