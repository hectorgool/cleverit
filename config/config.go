package config

import (
	"cleverit/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init() {

	dsn := "cleverit:asdf@tcp(db:3306)/cervezas?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=True"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&schema.BeerItem{}, &schema.BeerBox{})
	createData()
}

// GetDB gets mysql conection
func GetDB() *gorm.DB {
	return db
}

func createData() {

	var beerBox = []schema.BeerBox{
		{ID: 1, PriceTotal: 22},
		{ID: 2, PriceTotal: 33},
	}
	GetDB().Create(&beerBox)

	var beerItem = []schema.BeerItem{
		{ID: 1, Name: "Golden", Brewery: "Kross", Country: "Chile", Price: 10.5, Currency: "EUR"},
		{ID: 2, Name: "Negra", Brewery: "Modelo", Country: "Mexico", Price: 1, Currency: "USD"},
	}
	GetDB().Create(&beerItem)

}
