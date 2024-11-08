package main

import (
	"fmt"
	"github.com/aguevs/HelloWorld/funciones"
)

func helloworld() string {
	return "Hello World!!"
}

func main() {
	fmt.Println(helloworld())
	fmt.Println(funciones.Chino())
}
