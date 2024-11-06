package main

import (
	"fmt"
	"github.com/AgueVs/HelloWorld/funciones"

func helloworld() string {
	return "Hello World!!"
}

func main() {
	fmt.Println(helloworld())
	fmt.Println(funciones.chino("Hello World!!"))
}
