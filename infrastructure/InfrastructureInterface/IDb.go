package InfrastructureInterface

import "github.com/jinzhu/gorm"

type IDb interface {
	begin() *gorm.DB
	connect() *gorm.DB
}
