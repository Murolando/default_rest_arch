package handler

import (
	"github.com/Murolando/default_rest_arch/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	// router.Static("/storage","./storage")
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	return router
}
