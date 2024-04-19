package item

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	ID    string `gorm:"column:id" json:"id"`
	Field string `gorm:"column:field" json:"field"`
	Value string `gorm:"column:value" json:"value"`
	gorm.Model
}

// create new item
func NewItem(db *gorm.DB, item Item) (Item, error) {
	return item, db.Create(&item).Error
}

// get all items
func GetItems(db *gorm.DB) (items []Item, err error) {
	return items, db.Find(&items).Error
}

// check if item exists
func CheckItemExists(db *gorm.DB, id uint) bool {

	// init vars
	item := &Item{}

	// check if row exists
	check := db.Where("id=?", id).First(&item)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// search items
func SearchItems(db *gorm.DB, item Item) (items []Item, err error) {
	return items, db.Where(&item).Find(&items).Error
}

// update item
func UpdateItem(db *gorm.DB, item Item) error {
	return db.Where("id=?", item.ID).Updates(&item).Error
}

// delete item
func DeleteItem(db *gorm.DB, item_id uuid.UUID) error {
	return db.Where("id=?", item_id).Delete(&Item{}).Error
}
