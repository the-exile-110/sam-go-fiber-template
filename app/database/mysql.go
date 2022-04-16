package database

import (
	"app/database/schema"
	"app/utils/aws_sdk"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var (
	GormDbMysql *gorm.DB
)

// GetSqlDriver Get database driver
func GetSqlDriver() (*gorm.DB, error) {

	secret, err := aws_sdk.CreateAWSSDKFactory().GetSecret()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&charset=utf8mb4&loc=Local", secret.Username, secret.Password, secret.Host, secret.Port, secret.DbName)
	fmt.Println("=====================")
	fmt.Println("start to connect to the database")

	gormDb, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
			Logger:                 logger.Default.LogMode(logger.Info),
		})
	if err != nil {
		// gorm database driver initialization failed
		return nil, err
	}
	fmt.Println("successfully connected to the database")
	fmt.Println("=====================")
	return gormDb, nil
}

func init() {
	if dbMysql, err := GetSqlDriver(); err != nil {
		log.Fatal(err.Error())
	} else {
		GormDbMysql = dbMysql
		err = GormDbMysql.AutoMigrate(
			&schema.Users{},
		)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
