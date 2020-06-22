package models

import (
	"net/http"

	sqlDB "github.com/SolidShake/GoMonitoring/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Site struct {
	gorm.Model
	Name string
	SUrl string
}

func CreateSite(c *gin.Context) {
	newSite := Site{
		Name: c.PostForm("name"),
		SUrl: c.PostForm("url"),
	}
	db := sqlDB.GetDB()
	db.Create(newSite)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "site successfuly created"})
}

func GetAllSites(c *gin.Context) {
	db := sqlDB.GetDB()
	var allSites []Site
	db.Find(&allSites)

	if len(allSites) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no sites found"})
		return
	}

	type jsonSite struct {
		Id   uint
		Name string
		Url  string
	}
	var jsonSites []jsonSite

	for _, site := range allSites {
		jsonSites = append(jsonSites, jsonSite{site.ID, site.Name, site.SUrl})
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": jsonSites})
}
