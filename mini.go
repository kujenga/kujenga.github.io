package main

import (
	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir(".")),
	)
	http.ListenAndServe(":8080", n)
}
