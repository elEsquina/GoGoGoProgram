package main 

import (
	"project3/controller"
	"net/http"
)

func main(){

	http.HandleFunc("/students", controller.StudentsRouter)
	http.HandleFunc("/students/{id}", controller.StudentsPathParamRouter)
	http.ListenAndServe(":8080", nil)


}