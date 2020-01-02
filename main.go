package main

import (
	"net/http"

	"github.com/vibhavsharma341/polyglot-go/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":5000", nil)
}
