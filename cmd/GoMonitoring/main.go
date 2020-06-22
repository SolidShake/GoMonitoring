package main

import (
	"fmt"

	"github.com/SolidShake/GoMonitoring/internal/config"
	"github.com/SolidShake/GoMonitoring/internal/db"
	"github.com/SolidShake/GoMonitoring/internal/db/migrations"
	"github.com/SolidShake/GoMonitoring/internal/routing"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	defer db.CloseDB()
	migrations.Migrate()

	router := gin.Default()
	routing.InitializeRoutes(router)

	port := fmt.Sprintf(":%s", config.GetConfig().Server.Port)
	router.Run(port)
}
