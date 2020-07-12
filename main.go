package main

import (
	"log"

	"github.com/Richby/tuitting/bd"
	"github.com/Richby/tuitting/handlers"
)

func main() {
	if bd.ChequeoConexion() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
