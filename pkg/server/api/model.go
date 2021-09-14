package handler

import (
	"strconv"

	"github.com/AndrewOYLK/ou-cmdb/model"
	"github.com/AndrewOYLK/ou-cmdb/utils"
	"github.com/gin-gonic/gin"
)

// @tags 模型
// @Summary 增加模型
// @Description 增加模型
// @Accept json
// @Produce json
// @Param data body model.Model true "数据体"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/model [post]
func CreateModel(ctx *gin.Context) {
	m := model.Model{}

	err := ctx.BindJSON(&m)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	err = model.CreateModel(m)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", nil)
}

// @tags 模型
// @Summary 删除模型
// @Description 删除模型
// @Accept json
// @Produce json
// @Param id path int true "模型ID"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/model/{id} [delete]
func DeleteModel(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 64)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	err = model.DeleteModel(id)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", nil)
}

// @tags 模型
// @Summary 更改模型
// @Description 更改模型
// @Accept json
// @Produce json
// @Param id path int true "模型ID"
// @Param data body model.Model true "数据体"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/model/{id} [put]
func UpdateModel(ctx *gin.Context) {
	var m = model.Model{}

	id, err := strconv.ParseInt(ctx.Param("id"), 0, 64)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	if err := ctx.BindJSON(&m); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	m.ID = id
	err = model.UpdateModel(m)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", nil)
}

// @tags 模型
// @Summary 获取模型
// @Description 获取模型
// @Accept json
// @Produce json
// @Param id path int true "模型ID"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/model/{id} [get]
func GetModel(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 64)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	model, err := model.GetModel(id)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", model)
}

// @tags 模型
// @Summary 列出模型
// @Description 列出模型
// @Accept json
// @Produce json
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/models [get]
func ListModels(ctx *gin.Context) {
	models, err := model.ListModel()
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", models)
}
