package main

import (
	"log"
	"net/http"

	"github.com/lwldcr/validator"
	"github.com/urfave/negroni"
)

func handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello there"))
}

func main() {
	n := negroni.New()
	ch := validator.Sha1Checker{
		Key:"test_key_demo",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handle)

	v := validator.NewValidator(&ch)
	n.Use(v)
	n.Use(negroni.NewLogger())
	n.UseHandler(mux)

	log.Fatal(http.ListenAndServe(":8088", n))
}
