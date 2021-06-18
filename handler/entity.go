package handler

import (
	"strconv"

	"github.com/AndrewOYLK/ou-cmdb/model"
	"github.com/AndrewOYLK/ou-cmdb/utils"
	"github.com/gin-gonic/gin"
)

/*
	创建一条记录的流程：（一次请求只提交一个记录，一次请求可提交多个记录）
	1. 属性值的合法性（Attribute.Matedata）
	2. 该条记录根据unique attributes判断是否违背属性组的唯一性（两种情况：数据库原有记录验证 + 批量提交记录内验证）
	3. 将值插入entity、value表（考虑并发写如何处理）

	删除一条记录的流程：
	1. 检查是否有关联实例，有则删除关联实例
	2. 删除entity、value表的记录

	列出记录的流程：
	1. 根据modelid搜索entity表
	2. 根据entity表搜索value表
	3. 把value表组装成一条条记录
*/

// @tags 记录
// @Summary 增加记录
// @Description 增加记录
// @Accept json
// @Produce json
// @param modelID query int true "模型ID"
// @Param data body []model.Value true "数据体"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /entity [post]
func CreateEntity(ctx *gin.Context) {
	modelID, err := strconv.ParseInt(ctx.Query("modelID"), 0, 64)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	values := []model.Value{}
	if err := ctx.BindJSON(&values); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	if err := model.ValidAndSaveEntity(modelID, values); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	utils.NewAPISuccess(ctx, 200, "ok", nil)
}

// @tags 记录
// @Summary 删除记录
// @Description 删除记录
// @Accept json
// @Produce json
// @Param id path int true "记录实例ID"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /entity/{id} [delete]
func DeleteEntity(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 64)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	if err := model.CheckLinkAndDeleteEntity(id); err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}
	utils.NewAPISuccess(ctx, 200, "ok", nil)
}

// @tags 记录
// @Summary 根据ModelID列出记录
// @Description 根据ModelID列出记录
// @Accept json
// @Produce json
// @param modelID query int true "模型ID"
// @Success 200 {object} utils.APISuccess
// @Failure 400 {object} utils.APIError
// @Router /entities [get]
func ListEntity(ctx *gin.Context) {
	/*
		[{
			entityID: 1,
			values: [[],[],[],[]]
		},{
			entityID: 2,
			values: [[],[],[],[]]
		},{
			entityID: 3,
			values: [[],[],[],[]]
		}]
	*/
	modelID, err := strconv.ParseInt(ctx.Query("modelID"), 0, 64)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}

	result, err := model.ListEntityWithValues(modelID)
	if err != nil {
		utils.NewAPIError(ctx, 400, err)
		return
	}
	utils.NewAPISuccess(ctx, 200, "ok", result)
}
