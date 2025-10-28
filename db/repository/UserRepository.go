package repository

import (
	"blog/db/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// 创建用户
func (r *UserRepository) CreateUser(user *model.User) error {
	return r.DB.Create(user).Error
}

// 通过用户名和密码校验用户是否存在，密码直接用明文
func (r *UserRepository) GetUserByLoginParam(username string, password string) *model.User {
	user := &model.User{}
	result := r.DB.Where("user_name = ? and password = ?", username, password).First(user)
	if result.Error != nil {
		panic("没查询到用户")
	}
	return user
}
