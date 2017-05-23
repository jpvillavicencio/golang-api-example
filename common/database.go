package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

type Database struct {
	DB *gorm.DB
}

func (db *Database) InitDB() {
	var err error
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	url := "tcp("+os.Getenv("MYSQL_URL")+")"
	dbName := os.Getenv("MYSQL_DATABASE")

	sqlUrl := username + ":" + password + "@" + url + "/" + dbName + "?charset=utf8&parseTime=True"
	log.Print(sqlUrl)
	db.DB, err = gorm.Open("mysql", sqlUrl)
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	db.DB.LogMode(true)
}

func (db *Database) InitSchema(m interface{}) {
	db.DB.AutoMigrate(m)
}

func (db *Database) DropSchema(m interface{}) {
	if db.DB.HasTable(m) {
		db.DB.DropTable(m)
	}
}

func (db *Database) CreateModel(m interface{}) {
	if err := db.DB.Create(m).Error; err != nil {
		return
	}
}

func (db *Database) GetAllModels(m interface{}) {
	if err := db.DB.Find(m).Error; err != nil {
		return
	}
}

func (db *Database) GetModel(m interface{}, id int) {
	if db.DB.First(m, id).Error != nil {
		return
	}
}

func (db *Database) DeleteModel(m interface{}, id int) {
	if db.DB.First(m, id).Error != nil {
		return
	}
	if err := db.DB.Delete(m).Error; err != nil {
		return
	}
}
