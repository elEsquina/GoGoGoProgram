package main 

import (
	"net/http"
	"bearerAuth/controller"
)

func main() {
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/", controller.Authenticate(controller.GetAllUsers))
	http.ListenAndServe(":8080", nil)
}
