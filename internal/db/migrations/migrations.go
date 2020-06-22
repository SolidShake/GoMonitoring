package migrations

import (
	"github.com/SolidShake/GoMonitoring/internal/db"
	"github.com/SolidShake/GoMonitoring/internal/db/models"
)

func Migrate() {
	db.GetDB().AutoMigrate(models.Site{}, models.Status{})
}
