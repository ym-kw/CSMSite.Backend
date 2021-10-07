package Infrastructure

import (
	"CSMSite.Backend/Controllers"
	"github.com/gin-gonic/gin"
)

type Routing struct {
	DB   *Db
	Gin  *gin.Engine
	Port string
}

func NewRouting(db *Db) *Routing {
	r := &Routing{
		DB:   db,
		Gin:  gin.Default(),
		Port: NewConfig().Routing.Port,
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	usersController := Controllers.NewUserController(r.DB)
	r.Gin.GET("/user/:id", func(c *gin.Context) { usersController.Get(c) })
}
