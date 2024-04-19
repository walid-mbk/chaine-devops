package company

import (
	"gorm.io/gorm"
)

type Company struct {
	ID        uint   `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Name      string `gorm:"column:name;not null;unique" json:"name"`
	Email     string `gorm:"column:email;not null;unique;" json:"email"`
	Phone     string `gorm:"column:phone;not null;" json:"phone"`
	Address   string `gorm:"column:address;" json:"address"`
	ManagedBy uint   `gorm:"column:managed_by;" json:"managed_by"`
	CreatedBy uint   `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// create new company
func NewCompany(db *gorm.DB, company Company) (Company, error) {
	return company, db.Create(&company).Error
}

// get all companies
func GetCompanies(db *gorm.DB) (company []Company, err error) {
	return company, db.Find(&company).Error
}

// check company exists
func CheckCompanyExists(db *gorm.DB, id uint) bool {

	// init vars
	company := &Company{}

	// check if row exists
	check := db.Where("id=?", id).First(&company)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// search companies
func SearchCompanies(db *gorm.DB, company Company) (companies []Company, err error) {
	return companies, db.Where(&company).Find(&companies).Error
}

// update company
func UpdateCompany(db *gorm.DB, company Company) error {
	return db.Where("id=?", company.ID).Updates(&company).Error
}

// delete company
func DeleteCompany(db *gorm.DB, company_id uint) error {
	return db.Where("id=?", company_id).Delete(&Company{}).Error
}
