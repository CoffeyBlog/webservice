package main

import (
	"github.com/webservice/controllers"
	"net/http"
)


func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)


}


