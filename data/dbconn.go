package data

import (
	"fmt"
	"go/models/api"
	"go/models/dbTable"

	_ "github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/task_manager?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	return db
}

// Migrate when init
func Migrate() {
	db := ConnectDB()

	tableList := []interface{}{
		&dbTable.User{},
		&dbTable.Task{},
		&dbTable.Category{},
	}
	db.AutoMigrate(tableList...)
}

// Create a admin user
func CreateInitUser() {
	var adminUser api.User
	var count int64

	db := ConnectDB()
	query := db.Where("account=?", "admin").Find(&adminUser)
	query.Count(&count)

	if count != 0 {
		fmt.Println("Admin exist.")
	} else {
		fmt.Println("No admin user, create one.")
		db.Create(&dbTable.User{
			UserName: "Admin User", Account: "admin", Password: "admin",
		})
	}

}
