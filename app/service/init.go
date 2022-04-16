package service

import (
	"app/database"
	"gorm.io/gorm"
)

func CreateServiceFactory() *Service {
	return &Service{mysqlDB: database.GormDbMysql}
}

type Service struct {
	mysqlDB *gorm.DB
}
