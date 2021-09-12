package main

//We import 4 important libraries
//1. “net/http” to access the core go http functionality
//2. “fmt” for formatting our text
//3. “html/templates” a library that allows us to interact with our html file.
//4. "time" - a library for working with date and time.

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
)

type Money struct {
	Gold   string
	Silver string
	Copper string
	CoinsToSpend string
}

type Percents struct {
	Gold   float64
	Silver float64
	Copper float64
}

var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)

func main() {
	InitialFunds := Money{"0", "0", "0", "0"}
	SpentFunds := Money{"0", "0", "0", "0"}
	FinalFunds := Money{"0", "0", "0", "0"}

	templates := template.Must(template.ParseFiles("templates/templates.html"))

	http.Handle("/static/", //final url can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	//This method takes in the URL path "/" and a function that takes in a response writer, and a http request.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		moneyTypes := []string{"Gold", "Silver", "Copper"}
		percentages := getMetalPercentages(r)
		coinsToReturn, _ := strconv.ParseFloat(r.FormValue("Spend"), 64)
		log.Printf("coinsToReturn: %v", coinsToReturn)

		gold := math.Floor(coinsToReturn * percentages.Gold)
		log.Printf("initial gold spent: %v", gold)
		silver := math.Floor(coinsToReturn * percentages.Silver)
		log.Printf("initial silver spent: %v", silver)
		copper := math.Floor(coinsToReturn * percentages.Copper)
		log.Printf("initial copper spent: %v", copper)

		gold, silver, copper = manageExpectedAndCalculated(gold, silver, copper, coinsToReturn, percentages)

		for _, metal := range moneyTypes {
			if initialAmount := r.FormValue(metal); initialAmount != "" {
				CalculateMetalSpent(metal, initialAmount, &InitialFunds, &SpentFunds, &FinalFunds)
			}
		}

		log.Printf("Initial funds before render: %v", InitialFunds)
		log.Printf("Spent funds before render: %v", SpentFunds)
		log.Printf("Final funds before render: %v", FinalFunds)
		if err := templates.ExecuteTemplate(w, "templates.html", map[string]Money{
			"initial": InitialFunds,
			"spent":   SpentFunds,
			"final":   FinalFunds,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		
		log.Println("")
	})

	//Start the web server, set the port to listen to 8080. Without a path it assumes localhost
	//Print any errors from starting the webserver using fmt
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8888", nil))
}

func manageExpectedAndCalculated(gold float64, silver float64, copper float64, coinsToReturn float64, percentages Percents) (float64, float64, float64) {
	if gold+silver+copper != coinsToReturn {
		randomNumber := float64(rand.Intn(1))
		if randomNumber <= percentages.Gold {
			gold = gold + 1
			log.Printf("gold was added to")
		} else if randomNumber <= percentages.Gold+percentages.Silver {
			silver = silver + 1
			log.Printf("silver was added to")
		} else {
			copper = copper + 1
			log.Printf("copper was added to")
		}
	}

	log.Printf("new metal values; gold, silver, copper: %v, %v, %v", gold, silver, copper)
	return gold, silver, copper
}

func getMetalPercentages(r *http.Request) Percents {
	gold, _ := strconv.ParseFloat(r.FormValue("Gold"), 64)
	log.Printf("gold form value: %v", gold)
	silver, _ := strconv.ParseFloat(r.FormValue("Silver"), 64)
	log.Printf("silver form value: %v", silver)
	copper, _ := strconv.ParseFloat(r.FormValue("Copper"), 64)
	log.Printf("copper form value: %v", copper)
	total := gold + silver + copper

	if total > 0 {
		goldPercent := gold / total
		silverPercent := silver / total
		copperPercent := copper / total

		percentages := Percents{goldPercent, silverPercent, copperPercent}
		log.Printf("metal percent: %v", percentages)

		return percentages
	} else {
		percentages := Percents{0, 0, 0}
		return percentages
	}

}

func CalculateMetalSpent(metal string, initialAmount string, InitialFunds *Money, SpentFunds *Money, FinalFunds *Money) {
	initialAmountInt, _ := strconv.Atoi(initialAmount)
	setMetalValue(metal, "initial", InitialFunds, initialAmount)

	spentMetalInteger := rand.Intn(initialAmountInt + 1)
	spentMetal := strconv.Itoa(spentMetalInteger)
	setMetalValue(metal, "spent", SpentFunds, spentMetal)

	finalMetalInteger := initialAmountInt - spentMetalInteger
	finalMetal := strconv.Itoa(finalMetalInteger)
	setMetalValue(metal, "final", FinalFunds, finalMetal)
}

func setMetalValue(metal string, fund string, fundType *Money, newValue string) {
	log.Printf("setting %v %v to %v", fund, metal, newValue)
	reflect.ValueOf(fundType).Elem().FieldByName(metal).SetString(newValue)
}
