package main

import (
	"math/rand"
	"testing"
	"time"

	"unicode/utf8"

	"github.com/aguevs/HelloWorld/funciones"
)

func TestEjercicio1(t *testing.T) {
	if Helloworld() != "Hello World!!" {
		t.Fatal("No me esta imprimiendo Hola mundo en ingles")
	}

	if funciones.Chino() != "你好世界 ！！" {
		t.Fatal("No me esta imprimiento Hola mundo en chino")
	}
}

func TestEjercicio2(t *testing.T) {

	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(26)

	subcadena := RandBytes(randomInt, "abcdefghijklmnopqrstuvwxyz")
	if len(subcadena) != randomInt {
		t.Fatal("No estamos sacando 10 elementos de la subcadena en bytes, usando el metodo RandBytes")
	}

	subcadena2 := RandBytes(randomInt, "a你cd$%&efghijklmnopq&/()%$·")
	if len(subcadena2) != randomInt {
		t.Fatal("No estamos sacando 10 elementos de la subcadena en bytes, usando el metodo RandBytes")
	}

	subcadena4 := RandRunes2(randomInt, "a你cd$%&efghijklmnopq&/()%$·")
	if utf8.RuneCountInString(subcadena4) != randomInt {
		t.Fatal("No estamos sacando 10 elementos de la subcadena en runes, usado metodo RandRunes2")
	}

	subcadena5 := RandRunes(randomInt, "a你cd$%&efghijklmnopq&/()%$·")
	if utf8.RuneCountInString(subcadena5) != randomInt {
		t.Fatal("No estamos sacando 10 elementos de la subcadena en runes, esta funcion tiene un error porque transforma una variable en bytes y a veces no cuenta bien. Usado el metodo RunRunes")
	}

	substring, err := RunasAleatorias(randomInt, "a你cd$%&efghijklmnopq&/()%$·")
	if err != nil {
		t.Fatal("RandRunes fallo usando el método runasAleatorias porque no se puede sacar una subcadena de la cadena", substring, err)
	}

}

/*
fmt.Println("----------------------------------------")
// Ejercicio3:
fmt.Println("EJERCICIO 3:")
// Normalmente calculamos π en 2 dimensiones el parametro de entrada es el tamaño de circulo (radio).
fmt.Printf("Aproximación de π en cuadrados con %d de radio: %f\n", 1, πMonteCarlo(1))
fmt.Printf("Aproximación de π en cuadrados con %d de radio: %f\n", 10, πMonteCarlo(10))
fmt.Println("----------------------------------------")
*/
func TestEjercicio4(t *testing.T) {

	conjunto := NuevoConjunto()
	conjunto.Añadir(byte(1))
	conjunto.Añadir(byte(3))

	if !(conjunto.Existe(byte(1))) {
		t.Fatal("La funcion Añadir no esta haciendo su trabaj. No encuentra un elemento añadido")
	}

	if conjunto.Existe(byte(2)) {
		t.Fatal("La funcion Añadir no esta haciendo su trabajo. Encuentra un elemento que no fue añadido")
	}

	conjunto.Borrar(byte(1))

	if conjunto.Existe(byte(1)) {
		t.Fatal("La funcion Borrar no esta haciendo su trabajo. Encuentra un elemento que no fue borrado")
	}
}
