package main

import (
	"log"
	"net/http"

	"github.com/xubiosueldos/framework/configuracion"
)

func main() {
	configuracion := configuracion.GetInstance()
	router := newRouter()

	server := http.ListenAndServe(":"+configuracion.Puertomicroserviciolegajo, router)

	log.Fatal(server)

}
