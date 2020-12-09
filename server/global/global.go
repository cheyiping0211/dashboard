package global

import (
	"github.com/robfig/cron/v3"
	"github.com/jinzhu/gorm"
)

var (
	GVA_CRON   *cron.Cron
	GVA_DB     *gorm.DB
)