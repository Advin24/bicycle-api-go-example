package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gitlab.asynchrony.com/brandon.adams/bicycle-api-go/bicycle"

	// "github.com/spf13/viper"
	"github.com/gorilla/mux"
)

var bicycles []bicycle.Bicycle

// GetBicycles get all the bicycles
func GetBicycles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(bicycles)
}

// GetBicycle get a single bicycle
func GetBicycle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range bicycles {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&bicycle.Bicycle{})
}

// CreateBicycle create a new bicycle
func CreateBicycle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var bicycle bicycle.Bicycle
	_ = json.NewDecoder(r.Body).Decode(&bicycle)
	bicycle.ID = params["id"]
	// bicycle = append(bicycles, bicycle)
	json.NewEncoder(w).Encode(bicycles)
}

// DeleteBicycle delete a bicycle
func DeleteBicycle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range bicycles {
		if item.ID == params["id"] {
			bicycles = append(bicycles[:index], bicycles[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(bicycles)
	}
}

// func boostrapViper() (Viper, error) {
// 	viper.SetConfigName("config")        // name of config file (without extension)
// 	viper.AddConfigPath("/etc/appname/") // path to look for the config file in
// 	viper.AddConfigPath(".")             // optionally look for config in the working directory
// 	err := viper.ReadInConfig()          // Find and read the config file
// 	if err != nil {                      // Handle errors reading the config file
// 		panic(fmt.Errorf("Fatal error config file: %s \n", err))
// 	}
// }

func main() {

	// viper, err := bootstrapViper()
	// if err != nil {
	// 	log.Fatalf("Failed: %s", err)
	// }

	router := mux.NewRouter()
	bicycles = append(bicycles, bicycle.Bicycle{ID: "1", Make: "Specialized", Model: "Tarmac SL3", Size: 54, FrameMaterial: &bicycle.FrameMaterial{Type: "Carbon"}})
	bicycles = append(bicycles, bicycle.Bicycle{ID: "2", Make: "ALl City", Model: "Nature Boy", Size: 53, FrameMaterial: &bicycle.FrameMaterial{Type: "Steel"}})
	router.HandleFunc("/bicycle", GetBicycles).Methods("GET")
	router.HandleFunc("/bicycle/{id}", GetBicycle).Methods("GET")
	router.HandleFunc("/bicycle/{id}", CreateBicycle).Methods("POST")
	router.HandleFunc("/bicycle/{id}", DeleteBicycle).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
