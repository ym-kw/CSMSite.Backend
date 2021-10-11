package InfrastructureInterface

import "github.com/jinzhu/gorm"

type IDb interface {
	Begin() *gorm.DB
	Connect() *gorm.DB
}
