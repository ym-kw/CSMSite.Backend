package Repositories

import (
	"CSMSite.Backend/Repositories/IRepositories"
	"github.com/jinzhu/gorm"
)

type DbRepository struct {
	DbRepository IRepositories.IDbRepository
}

func (db *DbRepository) Begin() *gorm.DB {
	return db.DbRepository.Begin()
}

func (db *DbRepository) Connect() *gorm.DB {
	return db.DbRepository.Connect()
}
