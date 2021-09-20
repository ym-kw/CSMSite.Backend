package Repositories

import "github.com/jinzhu/gorm"

type IDbRepository interface {
	Begin() *gorm.DB
	Connect() *gorm.DB
}
