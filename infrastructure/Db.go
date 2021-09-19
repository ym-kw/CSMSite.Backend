package Infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Db struct {
	Host       string
	Username   string
	Password   string
	DbName     string
	Connection *gorm.DB
}

func NewDb() *Db {
	c := NewConfig()
	return newDb(&Db{
		Host:     c.Db.Production.Host,
		Username: c.Db.Production.Username,
		Password: c.Db.Production.Password,
		DbName:   c.Db.Production.DBName,
	})
}

func newDb(d *Db) *Db {
	db, err := gorm.Open("mysql", d.Username+":"+d.Password+"@tcp("+d.Host+")/"+d.DbName+"?charset=utf-8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	d.Connection = db
	return d
}

func (db *Db) Begin() *gorm.DB {
	return db.Connection.Begin()
}

func (db *Db) Connect() *gorm.DB {
	return db.Connection
}
