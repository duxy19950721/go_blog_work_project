package repository

import (
	"blog/db/model"

	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

// 创建文章
func (r *PostRepository) CreatePost(post *model.Post) error {
	return r.DB.Create(post).Error
}

// 查询文章
func (r *PostRepository) SelectPostList() ([]model.Post, error) {
	posts := []model.Post{}
	result := r.DB.Preload("Comments").Preload("User").Find(&posts)
	return posts, result.Error
}

// 查询单个详情
func (r *PostRepository) SelectPostDetail(postId uint) (model.Post, error) {
	post := model.Post{}
	result := r.DB.Preload("Comments").Preload("User").First(&post, postId)
	return post, result.Error
}

// 更新文章
func (r *PostRepository) UpdatePost(post *model.Post) error {
	return r.DB.Select("title", "content").Updates(post).Error
}

// 删除文章
func (r *PostRepository) DeletePost(postId uint) error {
	return r.DB.Delete(&model.Post{}, postId).Error
}
