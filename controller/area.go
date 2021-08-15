package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-zelus/api-layout/entity/state"
	"github.com/go-zelus/api-layout/service"
)

type AreaController struct {
	Controller
}

// Provinces
// @Summary 获取省份列表
// @Description 获取省份列表
// @Tags 行政区域
// @Accept json
// @Produce json
// @Success 200 {array} entity.Area ""
// @Failure 400 {object} state.ErrorResponse ""
// @Router /areas/provinces [get]
func (c *AreaController) Provinces(ctx *gin.Context) {
	areaSvc := service.NewAreaService()
	res, err := areaSvc.Provinces(ctx)
	if err != nil {
		c.JSON(ctx, err)
		return
	}
	c.JSON(ctx, res)
}

// Cities
// @Summary 获取城市列表
// @Description 根据省份行政代码获取城市列表
// @Tags 行政区域
// @Accept json
// @Produce json
// @Param code path int true "省份行政代码"
// @Success 200 {array} entity.Area "城市列表"
// @Failure 400 {object} state.ErrorResponse "错误信息"
// @Router /areas/cities/{code} [get]
func (c *AreaController) Cities(ctx *gin.Context) {
	code := ctx.Param("code")
	if code == "" {
		c.JSON(ctx, state.InvalidParams)
		return
	}
	areaCode, err := strconv.ParseInt(code, 10, 64)
	if err != nil {
		c.JSON(ctx, state.InvalidParams)
		return
	}
	areaSvc := service.NewAreaService()
	res, err := areaSvc.Cities(ctx, areaCode)
	if err != nil {
		c.JSON(ctx, err)
		return
	}
	c.JSON(ctx, res)
}

// Districts
// @Summary 获取区县列表
// @Description 根据城市行政代码获取区县列表
// @Tags 行政区域
// @Accept json
// @Produce json
// @Param code path int true "城市行政代码"
// @Success 200 {array} entity.Area "区县列表"
// @Failure 400 {object} state.ErrorResponse "错误信息"
// @Router /areas/districts/{code} [get]
func (c *AreaController) Districts(ctx *gin.Context) {
	code := ctx.Param("code")
	if code == "" {
		c.JSON(ctx, state.InvalidParams)
		return
	}
	areaCode, err := strconv.ParseInt(code, 10, 64)
	if err != nil {
		c.JSON(ctx, state.InvalidParams)
		return
	}
	areaSvc := service.NewAreaService()
	res, err := areaSvc.Districts(ctx, areaCode)
	if err != nil {
		c.JSON(ctx, err)
		return
	}
	c.JSON(ctx, res)
}

// Streets
// @Summary 获取街道列表
// @Description 根据区县行政代码获取街道列表
// @Tags 行政区域
// @Accept json
// @Produce json
// @Param code path int true "区县行政代码"
// @Success 200 {array} entity.Area "街道列表"
// @Failure 400 {object} state.ErrorResponse "错误信息"
// @Router /areas/streets/{code} [get]
func (c *AreaController) Streets(ctx *gin.Context) {
	code := ctx.Param("code")
	if code == "" {
		c.JSON(ctx, state.InvalidParams)
		return
	}
	areaCode, err := strconv.ParseInt(code, 10, 64)
	if err != nil {
		c.JSON(ctx, state.InvalidParams)
		return
	}
	areaSvc := service.NewAreaService()
	res, err := areaSvc.Streets(ctx, areaCode)
	if err != nil {
		c.JSON(ctx, err)
		return
	}
	c.JSON(ctx, res)
}

// Committees
// @Summary 获取居委会列表
// @Description 根据街道行政代码获取绝味会列表
// @Tags 行政区域
// @Accept json
// @Produce json
// @Param code path int true "街道行政代码"
// @Success 200 {array} entity.Area "居委会列表"
// @Failure 400 {object} state.ErrorResponse "错误信息"
// @Router /areas/committees/{code} [get]
func (c *AreaController) Committees(ctx *gin.Context) {
	code := ctx.Param("code")
	if code == "" {
		c.JSON(ctx, state.InvalidParams)
		return
	}
	areaCode, err := strconv.ParseInt(code, 10, 64)
	if err != nil {
		c.JSON(ctx, state.InvalidParams)
		return
	}
	areaSvc := service.NewAreaService()
	res, err := areaSvc.Committees(ctx, areaCode)
	if err != nil {
		c.JSON(ctx, err)
		return
	}
	c.JSON(ctx, res)
}
