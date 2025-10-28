package biz

// 初始化Swagger配置
// @title Gin Web API
// @version 1.0
// @description RESTful API 文档
// @host localhost:8080
// @BasePath /user

import (
	"blog/db/config"
	"blog/db/model"
	"blog/db/repository"
	"blog/gin/middleware"

	"github.com/gin-gonic/gin"
)

// @Summary 登录接口
// @Tags auth
// @Accept  json
// @Produce json
// @Param   username query string true "用户名"
// @Param   password query string true "密码"
// @Router /user/login [get]
func Login(c *gin.Context) {
	db := config.InitDB()
	ur := repository.UserRepository{DB: db}

	username := c.Query("username")
	password := c.Query("password")

	loginUser := ur.GetUserByLoginParam(username, password)
	// 鉴权
	token, err := middleware.GenerateToken(loginUser.ID)
	if err != nil {
		responseFail(c, "生成token失败", err)
		return
	}
	c.Header("Authorization", "Bearer "+token)
	responseSuccess(c, "用户登录成功")
}

// @Summary 注册接口
// @Tags auth
// @Accept  json
// @Produce json
// @Param   username query string true "用户名"
// @Param   password query string true "密码"
// @Router /user/registry [get]
func Registory(c *gin.Context) {
	db := config.InitDB()
	ur := repository.UserRepository{DB: db}

	user := &model.User{}
	c.ShouldBind(user)

	error := ur.CreateUser(user)
	if error != nil {
		responseFail(c, "创建用户失败", error)
		return
	}

	// todo 鉴权逻辑后续实现
	responseSuccess(c, "创建用户成功")
}
