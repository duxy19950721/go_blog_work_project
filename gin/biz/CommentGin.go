package biz

import (
	"blog/db/config"
	"blog/db/model"
	"blog/db/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 创建评论
// @Tags auth
// @Accept  json
// @Produce json
// @Param   comment body string true "评论内容json字符串"
// @Router /comment/create [post]
func CreateComment(c *gin.Context) {
	comment := model.Comment{}
	err := c.ShouldBind(&comment)
	if err != nil {
		panic("参数绑定失败: " + err.Error())
	}
	commentRepository := getCommentRepository()
	err = commentRepository.CreateComment(&comment)
	if err != nil {
		panic("创建评论失败: " + err.Error())
	}
	responseSuccess(c, "创建评论成功")
}

// @Summary 根据文章ID查询评论
// @Tags auth
// @Accept  json
// @Produce json
// @Param   post_id query uint true "文章ID"
// @Router /comment/getListByPostId [get]
func GetCommentListByPostId(c *gin.Context) {
	postId, err := strconv.ParseUint(c.Query("post_id"), 10, 32)
	if postId == 0 {
		panic("post_id不能为空")
	}
	if err != nil {
		responseFail(c, "获取post_id失败", err)
		return
	}
	commentRepository := getCommentRepository()
	comments, err := commentRepository.GetCommentListByPostId(uint(postId))
	if err != nil {
		panic("查询评论失败" + err.Error())
	}
	responseSuccessWithData(c, "查询评论成功", comments)
}

func getCommentRepository() repository.CommentRepository {
	db := config.InitDB()
	return repository.CommentRepository{DB: db}
}
