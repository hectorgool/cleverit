package controller

import (
	"cleverit/config"
	"cleverit/currencylayer"
	"cleverit/schema"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	BEERBOX float64 = 6
)

func SearchBeers(c *gin.Context) {

	var (
		rows []schema.BeerItem
	)

	if err := config.GetDB().Find(&rows).Error; err != nil {
		c.JSON(400, gin.H{"cerveza": "Request invalida"})
	} else {
		// Operacion exitosa
		c.JSON(200, gin.H{
			"api":     schema.Beer,
			"cerveza": rows,
		})
	}
}

func AddBeers(c *gin.Context) {

	var (
		beer, row schema.BeerItem
	)

	c.BindJSON(&beer)

	if err := config.GetDB().Where("ID = ?", beer.ID).First(&row).Error; err != nil {
		log.Println(err)
	}

	// El ID de la cerveza ya existe
	if row.ID != 0 {
		c.JSON(409, gin.H{"cerveza": "El ID de la cerveza ya existe"})
	}

	if err := config.GetDB().Create(&beer).Error; err != nil {
		c.JSON(400, gin.H{"cerveza": "Request invalida"})
	}
	// Cerveza creada
	c.JSON(201, gin.H{
		"api":     schema.Beer,
		"cerveza": "Cerveza creada",
	})
}

func SearchBeerById(c *gin.Context) {

	var (
		row schema.BeerItem
	)

	beerID := c.Param("beerID")
	if err := config.GetDB().Where("ID = ?", beerID).First(&row).Error; err != nil {
		// El Id de la cerveza no existe
		c.JSON(404, gin.H{"cerveza": "El Id de la cerveza no existe"})
	} else {
		// Operacion exitosa
		c.JSON(200, gin.H{
			"api":     schema.Beer,
			"cerveza": row,
		})
	}

}

func BoxBeerPriceById(c *gin.Context) {

	var (
		row schema.BeerItem
	)

	beerID := c.Param("beerID")
	if err := config.GetDB().Where("ID = ?", beerID).First(&row).Error; err != nil {
		// El Id de la cerveza no existe
		c.JSON(404, gin.H{"cerveza": "El Id de la cerveza no existe"})
	}

	quanty := c.Query("quanty")
	currency := c.Query("currency")
	price := princeTotal(row.Price, quanty, currency)

	// Operacion exitosa
	c.JSON(200, gin.H{
		"api":     schema.Beer,
		"cerveza": price,
	})

}

func princeTotal(price float64, quanty string, currency string) float64 {

	var priceTotal float64

	if len(quanty) == 0 && len(currency) == 0 {
		priceTotal = price * BEERBOX
	}
	if len(quanty) == 0 && len(currency) > 0 {
		priceTotal = price * currencylayer.GetCurrency(currency) * BEERBOX
	}
	if len(quanty) > 0 && len(currency) == 0 {
		priceTotal = price * StringToFloat64(quanty) * BEERBOX
	}
	if len(quanty) > 0 && len(currency) > 0 {
		priceTotal = price * StringToFloat64(quanty) * currencylayer.GetCurrency(currency) * BEERBOX
	}
	return priceTotal

}

func StringToFloat64(value string) float64 {
	number, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		log.Println(err)
	}
	out := number
	return float64(out)
}
