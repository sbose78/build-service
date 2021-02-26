package build

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Build - Our struct for all builds
type Build struct {
	ID      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Builds is the total list of builds ever done!
var Builds []Build

func ReturnAllBuilds(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBuilds")
	json.NewEncoder(w).Encode(Builds)
}

func ReturnSingleBuild(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, build := range Builds {
		if build.ID == key {
			json.NewEncoder(w).Encode(build)
		}
	}
}

func CreateNewBuild(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Build struct
	// append this to our Builds array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var build Build
	json.Unmarshal(reqBody, &build)
	// update our global Builds array to include
	// our new Build
	Builds = append(Builds, build)

	json.NewEncoder(w).Encode(build)
}

func DeleteBuild(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, build := range Builds {
		if build.ID == id {
			Builds = append(Builds[:index], Builds[index+1:]...)
		}
	}

}
