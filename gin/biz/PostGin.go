package biz

import (
	"blog/db/config"
	"blog/db/model"
	"blog/db/repository"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 创建文章，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容
// @Tags auth
// @Accept  json
// @Produce json
// @Param   post body string true "文章信息"
// @Router /post/create [post]
func CreatePost(c *gin.Context) {
	post := model.Post{}
	c.ShouldBindBodyWithJSON(&post)
	pr := getPostRepository()

	// todo 需要鉴权
	log.Println("创建文章的参数：", post)
	if err := pr.CreatePost(&post); err != nil {
		responseFail(c, "创建文章失败", err)
		return
	}
	responseSuccess(c, "创建文章成功")
}

// @Summary 查询文章功能，支持获取所有文章列表和单个文章的详细信息。
// @Tags auth
// @Produce json
// @Router /post/query [get]
func GetPostList(c *gin.Context) {
	post := model.Post{}
	c.ShouldBind(&post)
	pr := getPostRepository()
	posts, err := pr.SelectPostList()
	if err != nil {
		responseFail(c, "查询文章失败", err)
		return
	}
	responseSuccessWithData(c, "查询文章成功", posts)
}

// @Summary 更新文章功能，只有文章的作者才能更新自己的文章。
// @Tags auth
// @Accept  json
// @Produce json
// @Param   content body string true "内容json字符串"
// @Router /post/update [post]
func UpdatePost(c *gin.Context) {
	var requestData struct {
		PostID  uint   `json:"post_id"`
		UserID  uint   `json:"user_id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	c.ShouldBindBodyWithJSON(&requestData)
	pr := getPostRepository()
	post, err := pr.SelectPostDetail(requestData.PostID)
	if err != nil {
		responseFail(c, "查询文章失败", nil)
		return
	}
	if post.UserID != requestData.UserID {
		responseFail(c, "您没有权限更新此文章", nil)
		return
	}
	post.Title = requestData.Title
	post.Content = requestData.Content

	updateErr := pr.UpdatePost(&post)
	if updateErr != nil {
		responseFail(c, "更新文章失败", updateErr)
		return
	}
	responseSuccess(c, "更新文章成功")
}

// @Summary 删除文章功能，只有文章的作者才能删除自己的文章。
// @Tags auth
// @Accept  json
// @Produce json
// @Param   post_id query uint true "文章ID"
// @Router /post/delete [delete]
func DeletePost(c *gin.Context) {
	post_id, err := strconv.ParseUint(c.Query("post_id"), 10, 32)
	if err != nil {
		responseFail(c, "获取post_id失败", err)
		return
	}
	log.Println("删除文章的ID：", post_id)
	pr := getPostRepository()
	err = pr.DeletePost(uint(post_id))
	if err != nil {
		responseFail(c, "删除文章失败", nil)
		return
	}
	responseSuccess(c, "删除文章成功")
}

func getPostRepository() repository.PostRepository {
	db := config.InitDB()
	return repository.PostRepository{DB: db}
}
