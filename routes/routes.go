package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/leocaversan/micro-services-goLang/controllers"
)

func HandlerRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalities", controllers.ExibitionAllPersonalities).Methods("GET")
	r.HandleFunc("/personalities/{id}", controllers.ExibitionOnePersonalitie).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
