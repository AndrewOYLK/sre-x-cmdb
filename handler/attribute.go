package handler

import (
	"strconv"

	"github.com/AndrewOYLK/ou-cmdb/model"
	"github.com/AndrewOYLK/ou-cmdb/utils"
	"github.com/gin-gonic/gin"
)

// @tags 属性
// @Summary 增加属性
// @Description 增加属性
// @Accept json
// @Produce json
// @Param data body model.Attribute true "数据体"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/attribute [post]
func CreateAttribute(ctx *gin.Context) {
	attr := model.Attribute{}

	if err := ctx.BindJSON(&attr); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	if err := model.CreateAttribute(attr); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", nil)
}

// @tags 属性
// @Summary 删除属性
// @Description 删除属性
// @Accept json
// @Produce json
// @Param id path int true "属性ID"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/attribute/{id} [delete]
func DeleteAttribute(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 64)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	if err := model.DeleteAttribute(id); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", nil)
}

// @tags 属性
// @Summary 更改属性
// @Description 更改属性
// @Accept json
// @Produce json
// @Param id path int true "属性ID"
// @Param data body model.Attribute true "数据体"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/attribute/{id} [put]
func UpdateAttribute(ctx *gin.Context) {
	attr := model.Attribute{}

	if err := ctx.BindJSON(&attr); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	if err := model.UpdateAttribute(attr); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", nil)
}

// @tags 属性
// @Summary 获取属性
// @Description 获取属性
// @Accept json
// @Produce json
// @Param id path int true "属性ID"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/attribute/{id} [get]
func GetAttribute(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 64)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	attribute, err := model.GetAttribute(id)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", attribute)
}

// @tags 属性
// @Summary 列出属性
// @Description 根据模型ID列出属性
// @Accept json
// @Produce json
// @param modelID query int true "模型ID"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/attributes [get]
func ListAttribute(ctx *gin.Context) {
	modelID, err := strconv.ParseInt(ctx.Query("modelID"), 0, 64)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	attributes, err := model.ListAttribute(modelID)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", attributes)
}

// @tags 唯一属性组
// @Summary 增加唯一属性组
// @Description 增加唯一属性组
// @Accept json
// @Produce json
// @Param data body model.UniqueAttrs true "数据体"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/uniqueattrs [post]
func CreateUniqueAttrs(ctx *gin.Context) {
	uniqueAttrs := model.UniqueAttrs{}

	if err := ctx.BindJSON(&uniqueAttrs); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	if err := model.CreateUniqueAttrs(uniqueAttrs); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", nil)
}

// @tags 唯一属性组
// @Summary 删除唯一属性组
// @Description 删除唯一属性组
// @Accept json
// @Produce json
// @Param id path int true "唯一属性组ID"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/uniqueattrs/{id} [delete]
func DeleteUniqueAttrs(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 64)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	if err := model.DeleteUniqueAttrs(id); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", nil)
}

// @tags 唯一属性组
// @Summary 更改唯一属性组
// @Description 更改唯一属性组
// @Accept json
// @Produce json
// @Param id path int true "唯一属性组ID"
// @Param data body model.UniqueAttrs true "数据体"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/uniqueattrs/{id} [put]
func UpdateUniqueAttrs(ctx *gin.Context) {
	uniqueAttrs := model.UniqueAttrs{}

	id, err := strconv.ParseInt(ctx.Param("id"), 0, 64)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	if err := ctx.BindJSON(&uniqueAttrs); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	uniqueAttrs.ID = id
	if err := model.UpdateUniqueAttrs(uniqueAttrs); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", nil)
}

// @tags 唯一属性组
// @Summary 列出唯一属性组
// @Description 根据模型ID列出唯一属性组
// @Accept json
// @Produce json
// @param modelID query int true "模型ID"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/uniqueattrs [get]
func ListUniqueAttrs(ctx *gin.Context) {
	modelID, err := strconv.ParseInt(ctx.Query("modelID"), 0, 64)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	uniqueAttrsList, err := model.ListUniqueAttrs(modelID)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", uniqueAttrsList)
}
