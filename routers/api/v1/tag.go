package v1

import (
	"example.com/example/internal/service"
	"example.com/example/models"
	"example.com/example/pkg/app"
	"example.com/example/pkg/e"
	"example.com/example/pkg/response"
	"example.com/example/pkg/setting"
	"example.com/example/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetTags(c *gin.Context) {
	var request service.TagListRequest
	valid, errs := app.BindAndValid(c, &request)
	if !valid {
		response.Result(e.INVALID_PARAMS, nil, errs.Error(), c)
		return
	}
	data := make(map[string]interface{})
	data["lists"] = models.GetTags(util.GetPage(c), setting.AppSetting.PageSize, request)
	data["total"] = models.GetTotal(request)
	response.OkWithData(data, c)
}
func AddTag(c *gin.Context) {
	var request service.CreateTagRequest
	valid, errs := app.BindAndValid(c, &request)
	if !valid {
		response.Result(e.INVALID_PARAMS, nil, errs.Error(), c)
		return
	}
	if models.ExistTagByName(request.Name) {
		response.Result(e.ERROR_EXIST_TAG, nil, "标签已经存在", c)
	}
	models.AddTag(request.Name, request.State, request.CreatedBy)

	response.Ok(c)
}

func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	var request service.UpdateTagRequest
	valid, errs := app.BindAndValid(c, &request)
	if !valid {
		response.Result(e.INVALID_PARAMS, nil, errs.Error(), c)
		return
	}

	if !models.ExistTagById(id) {
		response.Result(e.ERROR_NOT_EXIST_TAG, nil, "标签不存在", c)
		return
	}
	models.EditTag(id, request)
	response.Ok(c)
}
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	code := e.INVALID_PARAMS
	code = e.SUCCESS
	if models.ExistTagById(id) {
		models.DeleteTag(id)
	} else {
		code = e.ERROR_NOT_EXIST_TAG
	}

	response.Result(code, make(map[string]interface{}), e.GetMsg(code), c)
}
