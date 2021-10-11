package Infrastructure

import (
	"CSMSite.Backend/Controllers"
	_ "CSMSite.Backend/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
	r.Gin.GET("/users", func(c *gin.Context) { usersController.GetAll(c) })
	r.Gin.POST("/user", func(c *gin.Context) { usersController.Post(c) })
	r.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
