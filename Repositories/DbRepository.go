package Repositories

import (
	Repositories "CSMSite.Backend/Repositories/Interface"
	"github.com/jinzhu/gorm"
)

type DbRepository struct {
	DbRepository Repositories.IDbRepository
}

func (db *DbRepository) Begin() *gorm.DB {
	return db.DbRepository.Begin()
}

func (db *DbRepository) Connect() *gorm.DB {
	return db.DbRepository.Connect()
}
