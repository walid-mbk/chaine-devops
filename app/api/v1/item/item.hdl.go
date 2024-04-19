package item

import (
	"net/http"
	"os"
	"regexp"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// create new item
func (db Database) NewItem(ctx *gin.Context) {

	// init vars
	var item Item
	//empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	//if empty_reg.MatchString(item.Field) || empty_reg.MatchString(item.Value) {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
	//	return
	//}

	// init new item
	new_item := Item{
		ID:    uuid.NewString(),
		Field: item.Field,
		Value: item.Value,
	}

	// create item
	if _, err := NewItem(db.DB, new_item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all items from database
func (db Database) GetItems(ctx *gin.Context) {

	// get items
	items, err := GetItems(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, items)
}

// search items from database
func (db Database) SearchItems(ctx *gin.Context) {

	// init vars
	var item Item

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search items from database
	items, err := SearchItems(db.DB, item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, items)
}

func (db Database) UpdateItem(ctx *gin.Context) {

	// init vars
	var item Item
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	item_id := ctx.Param("id")

	// check values validity
	if empty_reg.MatchString(item.Field) || empty_reg.MatchString(item.Value) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	item.ID = item_id

	// update item
	if err := UpdateItem(db.DB, item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteItem(ctx *gin.Context) {

	// get id from path
	item_id := uuid.MustParse(ctx.Param("id"))

	// delete item
	if err := DeleteItem(db.DB, item_id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
