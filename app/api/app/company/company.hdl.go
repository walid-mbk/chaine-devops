package company

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"template_rest_api/api/app/user"
	"template_rest_api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// create new company
func (db Database) NewCompany(ctx *gin.Context) {

	// init vars
	var company Company
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	log.Println("ok")

	// check json validity
	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	log.Println("ok")
	// check fields
	if empty_reg.MatchString(company.Name) || empty_reg.MatchString(company.Email) || empty_reg.MatchString(company.Phone) || empty_reg.MatchString(company.Address) || company.ManagedBy < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	log.Println("ok")
	// check user exists
	if exists := user.CheckUserExists(db.DB, company.ManagedBy); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "manager does not exist"})
		return
	}
	log.Println("ok")

	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// init new company
	new_company := Company{
		Name:      company.Name,
		Email:     company.Email,
		Phone:     company.Phone,
		Address:   company.Address,
		ManagedBy: company.ManagedBy,
		CreatedBy: session.UserID,
	}

	// create new company
	_, err := NewCompany(db.DB, new_company)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all companies from database
func (db Database) GetCompanies(ctx *gin.Context) {

	// get companies
	companies, err := GetCompanies(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, companies)
}

// search companies from database
func (db Database) SearchCompanies(ctx *gin.Context) {

	// init vars
	var company Company

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search companies
	companies, err := SearchCompanies(db.DB, company)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, companies)
}

// update company
func (db Database) UpdateCompany(ctx *gin.Context) {

	// init vars
	var company Company
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(company.Name) || empty_reg.MatchString(company.Email) || empty_reg.MatchString(company.Phone) || empty_reg.MatchString(company.Address) || company.ManagedBy < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// check subsidiary_of exists
	if exists := CheckCompanyExists(db.DB, company.ManagedBy); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid manager id"})
		return
	}

	// check user exists
	if exists := user.CheckUserExists(db.DB, company.ManagedBy); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "manager does not exist"})
		return
	}

	if company.ManagedBy != 0 {

		// check user exists
		if exists := user.CheckUserExists(db.DB, company.ManagedBy); !exists {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
			return
		}
	}

	// get id value from path
	company_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// init update user
	update_company := Company{
		ID:        uint(company_id),
		Name:      company.Name,
		Email:     company.Email,
		Phone:     company.Phone,
		Address:   company.Address,
		ManagedBy: company.ManagedBy,
	}

	// update company
	if err = UpdateCompany(db.DB, update_company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// delete company from database
func (db Database) DeleteCompany(ctx *gin.Context) {

	// get id value from path
	company_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete company
	if err := DeleteCompany(db.DB, uint(company_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
