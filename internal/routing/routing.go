package routing

import (
	"github.com/SolidShake/GoMonitoring/internal/db/models"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")

	v1.GET("allSites/", models.GetAllSites)
	v1.POST("addSite/", models.CreateSite)
	v1.GET("allStatuses/", models.GetAllStatuses)
}
