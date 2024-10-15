package global

import (
	"go_ecommerce/pkg/logger"
	"go_ecommerce/pkg/setting"

	"gorm.io/gorm"
)

var(
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)