package main

import ("net/http"
 		"log"
 		"github.com/xubiosueldos/conexionBD"
)

func main() {

	db := conexionBD.connectBD()
	db.CreateTable(&Legajo{})

	router := newRouter()
	
	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)

}
