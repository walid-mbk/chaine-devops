package item

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ItemRoutes(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", baseInstance.NewItem)
	router.GET("/all", baseInstance.GetItems)
	router.POST("/search", baseInstance.SearchItems)
	router.PUT("/:id", baseInstance.UpdateItem)
	router.DELETE("/:id", baseInstance.DeleteItem)
}
