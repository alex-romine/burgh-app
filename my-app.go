package main

//We import 4 important libraries
//1. “net/http” to access the core go http functionality
//2. “fmt” for formatting our text
//3. “html/templates” a library that allows us to interact with our html file.
//4. "time" - a library for working with date and time.

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
)

type Money struct {
	Gold   string
	Silver string
	Copper string
}

func main() {
	Initial_funds := Money{"0", "0", "0"}
	Spent_funds := Money{"0", "0", "0"}
	Final_funds := Money{"0", "0", "0"}

	templates := template.Must(template.ParseFiles("templates/templates.html"))

	http.Handle("/static/", //final url can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	//This method takes in the URL path "/" and a function that takes in a response writer, and a http request.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//Takes the name from the URL query e.g ?name=Martin, will set welcome.Name = Martin.
		money_types := []string{"Gold", "Silver", "Copper"}
		for _, metal := range money_types {
			if initial_amount := r.FormValue(metal); initial_amount != "" {
				fmt.Println(metal)
				initial_amount_int, _ := strconv.Atoi(initial_amount)
				reflect.ValueOf(&Initial_funds).Elem().FieldByName(metal).SetString(initial_amount)

				spent_metal_int := rand.Intn(initial_amount_int + 1)
				spent_metal := strconv.Itoa(spent_metal_int)
				reflect.ValueOf(&Spent_funds).Elem().FieldByName(metal).SetString(spent_metal)

				final_metal_int := initial_amount_int - spent_metal_int
				final_metal := strconv.Itoa(final_metal_int)
				reflect.ValueOf(&Final_funds).Elem().FieldByName(metal).SetString(final_metal)
			}
		}
		if err := templates.ExecuteTemplate(w, "templates.html", map[string]Money{
			"initial": Initial_funds,
			"spent":   Spent_funds,
			"final":   Final_funds,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	//Start the web server, set the port to listen to 8080. Without a path it assumes localhost
	//Print any errors from starting the webserver using fmt
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8888", nil))
}
