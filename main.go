/*
El objetivo de este programa es realizar la ejecucion de un programa, levantarlo en docker, para posteriormente correrlo en paralelo con otro en kubernetes, para practicar el despliegue en docker, y el mismo en kubernetes
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola mundo con Go + Docker")
}

func escribirMensaje(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Nuevo mensaje que se vera si se refresco el contenedor")
}

func main() {
	http.HandleFunc("/", hello) // funcion que si se ingresa a la ruta "/", se utiliza la funcion hello
	http.HandleFunc("/NuevoMensaje", escribirMensaje)

	fmt.Println("Servidor ejecutandose en el puerto:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
