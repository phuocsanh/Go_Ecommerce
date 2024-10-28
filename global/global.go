package global

import (
	"go_ecommerce/pkg/logger"
	"go_ecommerce/pkg/setting"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config  setting.Config
	Logger  *logger.LoggerZap
	Rdb    *redis.Client
	Mdb    	*gorm.DB
)