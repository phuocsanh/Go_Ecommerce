package initialize

import (
	"fmt"
	"go_ecommerce/global"
	"go_ecommerce/internal/po"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// checkErrorPanic logs the error and panics if the error is not nil
func CheckErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}


func InitMysql() {
	m := global.Config.Mysql
	// Build the Data Source Name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username, m.Password, m.Host, m.Port, m.Dbname)
	// Open the connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	CheckErrorPanic(err, "Failed to initialize MySQL")

	global.Logger.Info("MySQL Initialized Successfully")
	global.Mdb = db


	// Set connection pool settings
	// A pool is a set of pre-maintained connections that improve performance.
	setPool()


	// Run migrations
	migrateTables()
}

// setPool sets the MySQL connection pool settings
func setPool() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	CheckErrorPanic(err, "Failed to get SQL DB from GORM")

	// Set connection pool configurations
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns) * time.Second)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
}



// migrateTables runs database migrations
func migrateTables() {
err:= global.Mdb.AutoMigrate(
	&po.User{},
	&po.Role{},
 )
 if err != nil {
	fmt.Println("Migration table failed", err)
 }
}
