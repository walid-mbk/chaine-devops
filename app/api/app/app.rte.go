package app

import (
	"template_rest_api/api/app/company"
	permission "template_rest_api/api/app/permission"
	"template_rest_api/api/app/role"
	"template_rest_api/api/app/user"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// declare app routes
func RoutesApps(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	// user routes
	user.UserRoutes(router.Group("/user"), db, enforcer)

	// role routes
	role.RoutesRoles(router.Group("/role"), db, enforcer)

	// permission routes
	permission.RoutesPermissions(router.Group("/permission"), db, enforcer)

	// company routes
	company.RoutesCompanies(router.Group("/company"), db, enforcer)
}
