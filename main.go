package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"pokemon-api/database"
)

func addNewPokemon(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var pokemon database.Pokemon

	err := json.Unmarshal(requestBody, &pokemon)
	if err != nil {
		fmt.Println("error:", err)
	}

	found := false
	for i := 0; i < len(database.PokemonDb); i++ {
		if database.PokemonDb[i] == pokemon {
			fmt.Println(found)
			found = true
			fmt.Println(found)
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}
	if !found {
		fmt.Println(found)
		database.PokemonDb = append(database.PokemonDb, pokemon)
	}
	w.WriteHeader(http.StatusOK)
}

func getAllPokemons(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(database.PokemonDb)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func handleRequests() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(commonMiddleware)
	myRouter.HandleFunc("/pokemons", getAllPokemons).Methods("GET")
	myRouter.HandleFunc("/pokemons", addNewPokemon).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("Pokemon Rest API")
	handleRequests()
}
