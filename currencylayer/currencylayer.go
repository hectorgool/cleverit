package currencylayer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	Currencylayer struct {
		Success   bool   `json:"success"`
		Terms     string `json:"terms"`
		Privacy   string `json:"privacy"`
		Timestamp int    `json:"timestamp"`
		Source    string `json:"source"`
		Quotes    Quotes
	}
	Quotes struct {
		USDUSD float64 `json:"USDUSD"`
		USDAUD float64 `json:"USDAUD"`
		USDCAD float64 `json:"USDCAD"`
		USDPLN float64 `json:"USDPLN"`
		USDMXN float64 `json:"USDMXN"`
		USDCHF float64 `json:"USDCHF"`
	}
)

func GetCurrencyLayer(currencie string) *Currencylayer {

	endpoint := fmt.Sprintf("http://api.currencylayer.com/live?access_key=%v&currencies=USD,AUD,CAD,PLN,MXN,CHF&format=1",
		"a1b51af6245aa52ad1c1e6801109b6db")

	// Build the request
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record Currencylayer

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	return &record
}

func GetCurrency(currency string) float64 {
	currencyAPI := GetCurrencyLayer(currency)
	switch currency {
	case "AUD":
		return currencyAPI.Quotes.USDAUD
	case "CAD":
		return currencyAPI.Quotes.USDCAD
	case "PLN":
		return currencyAPI.Quotes.USDPLN
	case "MXN":
		return currencyAPI.Quotes.USDMXN
	case "CHF":
		return currencyAPI.Quotes.USDCHF
	default:
		return currencyAPI.Quotes.USDUSD
	}
}
