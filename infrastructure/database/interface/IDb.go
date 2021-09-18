package Database

import "github.com/jinzhu/gorm"

type Db interface {
	begin() *gorm.DB
	connect() *gorm.DB
}
