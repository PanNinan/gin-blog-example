package v1

import (
	"example.com/example/models"
	"example.com/example/pkg/e"
	"example.com/example/pkg/response"
	"example.com/example/pkg/setting"
	"example.com/example/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.INVALID_PARAMS
	var data interface{}
	if models.ExistArticleById(id) {
		data = models.GetArticle(id)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	data["lists"] = models.GetArticles(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetArticleTotal(maps)
	response.Result(e.SUCCESS, data, e.GetMsg(e.SUCCESS), c)
}

func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	code := e.INVALID_PARAMS
	if models.ExistTagById(tagId) {
		data := make(map[string]interface{})
		data["tag_id"] = tagId
		data["title"] = title
		data["desc"] = desc
		data["content"] = content
		data["created_by"] = createdBy
		data["state"] = state

		models.AddArticle(data)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_TAG
	}
	response.Result(code, make(map[string]interface{}), e.GetMsg(code), c)
}

func EditArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	code := e.INVALID_PARAMS
	if models.ExistArticleById(id) {
		if models.ExistTagById(tagId) {
			data := make(map[string]interface{})
			if tagId > 0 {
				data["tag_id"] = tagId
			}
			if title != "" {
				data["title"] = title
			}
			if desc != "" {
				data["desc"] = desc
			}
			if content != "" {
				data["content"] = content
			}

			data["modified_by"] = modifiedBy

			models.EditArticle(id, data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}
	response.Result(code, make(map[string]interface{}), e.GetMsg(code), c)
}

func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	code := e.INVALID_PARAMS
	if models.ExistArticleById(id) {
		models.DeleteArticle(id)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
