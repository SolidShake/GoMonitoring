package models

import (
	"net/http"

	monitoringDB "github.com/SolidShake/GoMonitoring/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Status struct {
	gorm.Model
	StatusCode   uint
	ResponseTime uint
	Site         Site `gorm:foreignkey:SiteID`
	SiteID       uint
}

func CreateStatus(statusCode uint, responseTime uint, Site Site) {
	db := monitoringDB.GetDB()
	var newStatus = &Status{
		StatusCode:   statusCode,
		ResponseTime: responseTime,
		Site:         Site,
	}
	db.Create(newStatus)
}

func GetAllStatuses(c *gin.Context) {
	db := monitoringDB.GetDB()

	var allSites []Site
	db.Find(&allSites)

	type jsonSiteStatuses struct {
		Site     Site
		Statuses []Status
	}

	var allStatuses []jsonSiteStatuses
	for _, site := range allSites {
		var statuses []Status
		db.Model(&site).Related(&statuses)
		if len(statuses) != 0 {
			allStatuses = append(allStatuses, jsonSiteStatuses{site, statuses})
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": allStatuses})
}
