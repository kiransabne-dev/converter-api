package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kiransabne/converter/routers"
)

func main() {
	log.Println("hello wotkd")
	router := routers.Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err : %s\n", err.Error())
	}
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/", router)
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
