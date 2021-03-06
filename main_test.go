package main

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", getAllPokemons).Methods("GET")
	return router
}

// Expect a 200 code
func TestGetAllPokemonsCode(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	log.Print(response.Body)
	assert.Equal(t, 200, response.Code, "OK")
}

// Expect only two pokemon in the body and 200 code meaning it was successful
func TestGetAllPokemonsBody(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	log.Println(response.Body.String())
	assert.Equal(t, 200, response.Code, "OK")
	assert.JSONEq(t, "[{\"Id\":\"1\",\"Name\":\"Pikachu\",\"Type\":\"Electric\"},{\"Id\":\"2\",\"Name\":\"Charmeleon\",\"Type\":\"Fire\"}]", response.Body.String(), "OK")
}

func Router2() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", addNewPokemon).Methods("POST")
	return router
}

// Expect an empty response to a POST request and 200 code meaning it was successful
func TestAddNewPokemonPOST(t *testing.T) {
	r := strings.NewReader("{\"Id\":\"3\",\"Name\":\"Charizard\",\"Type\":\"Fire\"}")
	log.Println(r)
	request, _ := http.NewRequest("POST", "/", r)
	response := httptest.NewRecorder()
	Router2().ServeHTTP(response, request)
	log.Println(response.Body.String())
	assert.Equal(t, 200, response.Code, "OK")
	assert.Equal(t, "", response.Body.String())
}
