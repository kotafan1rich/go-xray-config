package handlers

import (
	"go-xray-config/internal/services"
	"go-xray-config/pkg/api"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type XRayHandler struct {
	xrayService services.XrayService
}

func NewXRayHandler(xrayService services.XrayService) *XRayHandler {
	return &XRayHandler{
		xrayService: xrayService,
	}
}

func (h *XRayHandler) AddNew(c *gin.Context) {
	result, err := h.xrayService.Add()
	if err != nil {
		api.InternalError(c, "add xray config failed: "+err.Error())
	}
	api.Success(c, result)
}

func (h *XRayHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.xrayService.Delete(uuid.MustParse(id))
	if err != nil {
		api.InternalError(c, "delete xray config failed: "+err.Error())
	}
	api.Success(c, "id deleted successfully")
}

func (h *XRayHandler) RegisterRoutes(router *gin.RouterGroup) {
	xrayGroup := router.Group("/xray")
	{
		xrayGroup.POST("/add", h.AddNew)
		xrayGroup.DELETE("/delete/:id", h.Delete)
	}
}
