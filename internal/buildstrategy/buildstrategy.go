package buildstrategy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// BuildStrategy - Our struct for all buildstrategies
type BuildStrategy struct {
	ID      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// BuildStrategies is the total list of builds ever done!
var BuildStrategies []BuildStrategy

func ReturnAllStrategies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBuilds")
	json.NewEncoder(w).Encode(BuildStrategies)
}

func ReturnSingleStrategy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, buildstrategy := range BuildStrategies {
		if buildstrategy.ID == key {
			json.NewEncoder(w).Encode(buildstrategy)
		}
	}
}

func CreateNewStrategy(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Build struct
	// append this to our BuildStrategies array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var buildstrategy BuildStrategy
	json.Unmarshal(reqBody, &buildstrategy)
	// update our global BuildStrategies array to include
	// our new Build
	BuildStrategies = append(BuildStrategies, buildstrategy)

	json.NewEncoder(w).Encode(buildstrategy)
}

func DeleteBuildStrategy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, buildstrategy := range BuildStrategies {
		if buildstrategy.ID == id {
			BuildStrategies = append(BuildStrategies[:index], BuildStrategies[index+1:]...)
		}
	}

}
