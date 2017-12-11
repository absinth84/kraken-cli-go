package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/beldur/kraken-go-api-client"
	"github.com/kr/pretty"
	"github.com/segmentio/go-prompt"
)

var key = ""
var secret = ""

var options = []string{
	"Balance",
	"Open Orders",
	"Closed Orders",
	"Buy",
	"Sell",
}

func main() {

	for {

		fmt.Printf("\nKraken api cli")
		i := prompt.Choose("Please Select an option", options)
		println("picked: " + options[i])
		switch i {
		case 0:
			fmt.Println("case 0")
			balance()
		case 1:
			fmt.Println("case 1")
			openorders()
		case 2:
			fmt.Println("case 2")
			closedorders()
		case 3:
			buy()

		}

	}

}

func balance() {
	api := krakenapi.New(key, secret)
	result, err := api.Query("Balance", map[string]string{})
	for i := 0; err != nil && i < 3; i++ {
		fmt.Println(i)
		result, err = api.Query("Balance", map[string]string{})
		if err != nil {
			log.Println(err)
		}
	}

	fmt.Printf("%# v", pretty.Formatter(result))

}

func openorders() {
	api := krakenapi.New(key, secret)
	result, err := api.OpenOrders(map[string]string{})
	for i := 0; err != nil && i < 3; i++ {
		fmt.Println(i)
		result, err = api.OpenOrders(map[string]string{})
		if err != nil {
			log.Println(err)
		}
	}

	fmt.Printf("%# v", pretty.Formatter(result))

}

func closedorders() {
	api := krakenapi.New(key, secret)
	result, err := api.ClosedOrders(map[string]string{})
	for i := 0; err != nil && i < 3; i++ {
		fmt.Println(i)
		result, err = api.ClosedOrders(map[string]string{})
		if err != nil {
			log.Println(err)
		}
	}

	fmt.Printf("%# v", pretty.Formatter(result))
	fmt.Printf("Result: %#v\n", result)

}

func buy() {
	var eur, volume, price float64
	var response string
	fmt.Println("Buy Bitcoin/EUR")
	api := krakenapi.New(key, secret)

	//Get the last trade
	ticker, err := api.Ticker(krakenapi.XXBTZEUR)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Last Trade:")
	fmt.Println(ticker.XXBTZEUR.Bid)

	fmt.Println("Add a new order")

	fmt.Println("Insert € to spend:")
	fmt.Scanf("%f", &eur)
	fmt.Println("Insert Price:")
	fmt.Scanf("%f", &price)
	volume = eur / price

	fmt.Println("Confirm buy", strconv.FormatFloat(volume, 'f', 8, 64), "BTC (", eur, "EURO) at", price, " [y/n]")
	fmt.Scanln(&response)
	if response == "y" {

		result, err := api.Query("AddOrder", map[string]string{
			"type":      "buy",
			"ordertype": "limit",
			"price":     strconv.FormatFloat(price, 'f', 8, 64),
			"volume":    strconv.FormatFloat(volume, 'f', 8, 64),
		})
		fmt.Printf("%# v", pretty.Formatter(result))
		if err != nil {
			log.Println(err)
		}
	} else {
		fmt.Println("Aborted")
	}
}

func sell() {
	fmt.Println("Sell Bitcoin/EUR")
	api := krakenapi.New(key, secret)

	//Get the last trade
	ticker, err := api.Ticker(krakenapi.XXBTZEUR)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Last Trade:")
	fmt.Println(ticker.XXBTZEUR.Bid)

	fmt.Println("Add a new order")
	result, err := api.AddOrder("XXBTZEUR", "sell", "limit", "volume", map[string]string{
		"price": "price",
	})
	fmt.Printf("%# v", pretty.Formatter(result))

}
