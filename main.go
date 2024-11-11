package main

import (
	"fmt"
	"github.com/aguevs/HelloWorld/funciones"
)
// Funcion que devuelve Hello World
func helloworld() string {
	return "Hello World!!"
}

// Funcion Main;
//  Ejercicio1: nos imprime por consola "Hola Mundo" en Ingles y Chino.
//                Hola Mundo en ingles usando una funcion creada en este paquete main.
//                Hola Mundo en chino usando una funcion creada en otro paquete llamado funciones creado por mi.
//  Ejercicio2:
func main() {
	fmt.Println(helloworld())
	fmt.Println(funciones.Chino())
}
