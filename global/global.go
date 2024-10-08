package global

import (
	"go_ecommerce/pkg/logger"
	"go_ecommerce/pkg/setting"
)

var(
	Config setting.Config
	Logger *logger.LoggerZap
)