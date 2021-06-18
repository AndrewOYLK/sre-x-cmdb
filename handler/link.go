package handler

import (
	"strconv"

	"github.com/AndrewOYLK/ou-cmdb/model"
	"github.com/AndrewOYLK/ou-cmdb/utils"
	"github.com/gin-gonic/gin"
)

// LinkType

// @tags 关联类型
// @Summary 新增关联类型
// @Description 新增关联类型
// @Accept json
// @Produce json
// @Param data body model.LinkType true "数据体"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/linktype [post]
func CreateLinkType(ctx *gin.Context) {
	linkType := model.LinkType{}

	if err := ctx.BindJSON(&linkType); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}
	utils.NewAPISuccess(ctx, 200, "ok", nil)
}

// @tags 关联类型
// @Summary 删除关联类型
// @Description 删除关联类型
// @Accept json
// @Produce json
// @Param id path int true "关联类型ID"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/linktype/{id} [delete]
func DeleteLinkType(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 64)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	if err := model.DeleteLinkType(id); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", nil)
}

// @tags 关联类型
// @Summary 更新关联类型
// @Description 更新关联类型
// @Accept json
// @Produce json
// @Param id path int true "关联类型ID"
// @Param data body model.LinkType true "数据体"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/linktype/{id} [put]
func UpdateLinkType(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 64)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	linkType := model.LinkType{}
	if err := ctx.BindJSON(&linkType); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	linkType.ID = id
	if err := model.UpdateLinkType(linkType); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", nil)
}

// @tags 关联类型
// @Summary 列出关联类型
// @Description 列出关联类型
// @Accept json
// @Produce json
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/linktypes [get]
func ListLinkTypes(ctx *gin.Context) {
	result, err := model.ListLinkTypes()
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", result)
}

// LinkModel

// @tags 模型关联关系
// @Summary 新增模型关联关系
// @Description 新增模型关联关系
// @Accept json
// @Produce json
// @Param data body model.LinkModel true "数据体"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/linkmodel [post]
func CreateLinkModel(ctx *gin.Context) {
	linkModel := model.LinkModel{}

	if err := ctx.BindJSON(&linkModel); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	if err := model.CreateLinkModel(linkModel); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", nil)
}

// @tags 模型关联关系
// @Summary 删除模型关联关系
// @Description 删除模型关联关系
// @Accept json
// @Produce json
// @Param id path int true "模型关联关系ID"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/linkmodel/{id} [delete]
func DeleteLinkModel(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 64)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	if err := model.DeleteLinkModel(id); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", nil)
}

// @tags 模型关联关系
// @Summary 列出模型关联关系
// @Description 列出模型关联关系
// @Accept json
// @Produce json
// @param modelID query int true "模型ID"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/linkmodels [get]
func ListLinkModels(ctx *gin.Context) {
	modelID, err := strconv.ParseInt(ctx.Query("modelID"), 0, 64)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	result, err := model.ListLinkModels(modelID)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", result)
}

// LinkEntity

// @tags 实例关联关系
// @Summary 新增实例关联关系
// @Description 新增实例关联关系
// @Accept json
// @Produce json
// @Param data body model.LinkEntity true "数据体"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/linkentity [post]
func CreateLinkEntity(ctx *gin.Context) {
	linkEntity := model.LinkEntity{}

	if err := ctx.BindJSON(&linkEntity); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	if err := model.ValidAndSaveLinkEntity(linkEntity); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", nil)
}

// @tags 实例关联关系
// @Summary 删除实例关联关系
// @Description 删除实例关联关系
// @Accept json
// @Produce json
// @Param id path int true "实例关联关系ID"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /api/v1/linkentity/{id} [delete]
func DeleteLinkEntity(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 64)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	if err := model.DeleteLinkEntity(id); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", nil)
}
