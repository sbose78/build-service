// main.go
//go:generate swagger generate spec

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sbose78/build-service/internal/build"
	"github.com/sbose78/build-service/internal/buildstrategy"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/builds", build.ReturnAllBuilds)
	myRouter.HandleFunc("/build", build.CreateNewBuild).Methods("POST")
	myRouter.HandleFunc("/build/{id}", build.DeleteBuild).Methods("DELETE")
	myRouter.HandleFunc("/build/{id}", build.ReturnSingleBuild)

	myRouter.HandleFunc("/buildstrategies", buildstrategy.ReturnAllStrategies)
	myRouter.HandleFunc("/buildstrategy", buildstrategy.CreateNewStrategy).Methods("POST")
	myRouter.HandleFunc("/buildstrategy/{id}", buildstrategy.DeleteBuildStrategy).Methods("DELETE")
	myRouter.HandleFunc("/buildstrategy/{id}", buildstrategy.ReturnSingleStrategy)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	build.Builds = []build.Build{
		build.Build{ID: "1", Title: "Hello", Desc: "Build Description", Content: "Build Content"},
		build.Build{ID: "2", Title: "Hello 2", Desc: "Build Description", Content: "Build Content"},
	}

	buildstrategy.BuildStrategies = []buildstrategy.BuildStrategy{
		buildstrategy.BuildStrategy{ID: "1", Title: "Hello", Desc: "Build Description", Content: "Build Content"},
		buildstrategy.BuildStrategy{ID: "2", Title: "Hello 2", Desc: "Build Description", Content: "Build Content"},
	}
	handleRequests()
}
