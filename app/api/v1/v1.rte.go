package v1

import (
	"template_rest_api/api/v1/item"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesV1(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	// Item Routes
	item.ItemRoutes(router.Group("/item"), db, enforcer)
}
