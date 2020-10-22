package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"pokemon-api/database"
)

func addNewPokemon(w http.ResponseWriter, r *http.Request) {
	var pokemon database.Pokemon
	requestBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(requestBody, &pokemon)
	database.PokemonDb = append(database.PokemonDb, pokemon)
	w.WriteHeader(http.StatusOK)
}

func getAllPokemons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.PokemonDb)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(commonMiddleware)
	myRouter.HandleFunc("/pokemons", getAllPokemons).Methods("GET")
	myRouter.HandleFunc("/pokemon/add", addNewPokemon).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
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
