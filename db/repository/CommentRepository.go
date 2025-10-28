package repository

import (
	"blog/db/model"

	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

// 创建评论
func (cr *CommentRepository) CreateComment(comment *model.Comment) error {
	return cr.DB.Create(comment).Error
}

// 查询某篇文章的全部评论
func (cr *CommentRepository) GetCommentListByPostId(postId uint) ([]model.Comment, error) {
	comments := []model.Comment{}
	result := cr.DB.Where("post_id = ?", postId).Find(&comments)
	return comments, result.Error
}
