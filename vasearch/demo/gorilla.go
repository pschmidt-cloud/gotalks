package main
import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// Set up routes
	router := mux.NewRouter()

	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/static/").Handler(s)

	http.Handle("/", router)
}