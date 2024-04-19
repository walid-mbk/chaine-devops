package company

import (
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// declare company routes
func RoutesCompanies(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("companies", "write", enforcer), baseInstance.NewCompany)
	router.GET("/all", middleware.Authorize("companies", "read", enforcer), baseInstance.GetCompanies)
	router.POST("/search", middleware.Authorize("companies", "read", enforcer), baseInstance.SearchCompanies)
	router.PUT("/:id", middleware.Authorize("companies", "write", enforcer), baseInstance.UpdateCompany)
	router.DELETE("/:id", middleware.Authorize("companies", "write", enforcer), baseInstance.DeleteCompany)
}
