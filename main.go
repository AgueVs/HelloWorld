package main

import (
	"fmt"
//	"log"
	"unicode/utf8"
	"math/rand"
	"time"
	"github.com/aguevs/HelloWorld/funciones"
)

// Funcion que devuelve Hello World
func helloworld() string {
	return "Hello World!!"
}

//funcion que devuelva una cadena de caracteres con n bytes tomados de una cadena input de entrada
func RandBytes (length int, cadena string) string {
        var substr string

        buf := make([]byte, length) 
        for i:= range buf {
             buf[i] = cadena[rand.Intn(len(cadena))]
	     substr = string(buf)}
     return substr
}

//funcion que devuelva una cadena de caracteres con n runes tomados de una cadena input de entrada
func RandRunes (length int, cadena string) string {

	//longitud del string en runes
        runes := utf8.RuneCountInString(cadena)
        substr := make([]rune, length)
	
	// Inicializar el generador de números aleatorios
        rand.Seed(time.Now().UnixNano())

        for i:= 0; i < length; i++ {
	     number := rand.Intn(runes)
	     for pos, runa := range cadena {
		   // pos es la posicion de byte.
		   //como podemos comprobar que esta posicon de byte corresponde a la posicion number de runa que queremos.
                   if pos == number {
			   substr[i] = runa    
                   }
	     }
	}
	return string(substr)
}

//Funcion hecha por Javi Valdez
func RandRunes2(n int, cadena string) string {

    cadenaRune := []rune(cadena)

    rand.Seed(time.Now().UnixNano())

    result := make([]rune, n)

    for i := 0; i < n; i++ {
        result[i] = cadenaRune[rand.Intn(len(cadenaRune))]
    }

    return string(result)
}

// Funcion para obtener una cadena de n runas aleatorias de la cadena de entrada
func runasAleatorias(n int, input string) (string, error) {

	// Convertir la cadena de entrada a un slice de runas
	runas := []rune(input)

	// Verificar que n no sea mayor que la longitud de la cadena de runas
	if n > len(runas) {
		return "", fmt.Errorf("El número de runas solicitado es mayor que la longitud de la cadena")
	}

	// Inicializar el generador de números aleatorios
	rand.Seed(time.Now().UnixNano())

	// Crear un mapa para almacenar índices únicos seleccionados
	indicesSeleccionados := make(map[int]struct{})

	// Seleccionar n índices únicos aleatoriamente
	for len(indicesSeleccionados) < n {
		indice := rand.Intn(len(runas))
		indicesSeleccionados[indice] = struct{}{}
	}

	// Crear una cadena con las runas seleccionadas
	var seleccionadas []rune
	for indice := range indicesSeleccionados {
		seleccionadas = append(seleccionadas, runas[indice])
	}

	// Convertir las runas seleccionadas de nuevo a una cadena
	resultado := string(seleccionadas)

	return resultado, nil
}


/*Requirio ayuda de internet y explicaciones matematicas para entender dicho metodo de estadisticas.

Generación de Puntos Aleatorios:

Generamos puntos aleatorios (x, y) dentro de un cuadrado que va de -radio a radio en ambas dimensiones.
Esto asegura que el círculo esté completamente dentro del cuadrado.

Determinación de Puntos Dentro del Círculo:
Verificamos si cada punto (x, y) cae dentro del círculo de radio radius centrado en el origen.
Esto se hace comprobando si x*x + y*y <= radio*radio.

Cálculo de π:
La proporción de puntos dentro del círculo con respecto al total de puntos generados nos proporciona una estimación de π multiplicando por 4,
ya que el área del círculo es π*radio^2 y el área del cuadrado es (2*radio)^2 = 4*radio^2.

Precisión:
La precisión de esta estimación mejora con un mayor número de puntos numPoints.
A más puntos, más precisa será la aproximación, pero también requerirá más tiempo de cómputo.

Este método es flexible y permite cambiar el radio del círculo, lo que ilustra cómo el método de Monte Carlo puede adaptarse a diferentes escenarios.
*/

func  πMonteCarlo(rad int) float64 {

	rand.Seed(time.Now().UnixNano())
	insideCircle := 0
	numPoints :=100000
	radio := float64(rad)

	for i := 0; i < numPoints; i++ {
		// Generar un punto aleatorio (x, y) en el cuadrado [-radio, radio] x [-radio, radio]
		x := rand.Float64()*2*radio-radio
		y := rand.Float64()*2*radio-radio

		// Comprobar si el punto está dentro del círculo de radio `radio`
		if x*x+y*y <= float64(radio*radio) {
			insideCircle++
		}
	}

	// La proporción de puntos dentro del círculo respecto al total
	// nos da π * (radio^2) / (4 * radio^2), por lo tanto multiplicamos por 4 para obtener π.
	return 4.0 * float64(insideCircle) / float64(numPoints)

}

/* Funcion Main;

  Ejercicio1: nos imprime por consola "Hola Mundo" en Ingles y Chino.
                Hola Mundo en ingles usando una funcion creada en este paquete main.
                Hola Mundo en chino usando una funcion creada en otro paquete llamado funciones creado por mi.
  Ejercicio2: funciones que nos devuelva una cadena de caracteres
                  Con n bytes tomados de una cadena input de entrada
                  Con n runes tomados de una cadena input de entrada (pendiente de hacer este paso)
  Ejericcio3: Programar la función π que devuelve una aproximación del valor del
              número irracional π usando el método de Monte-Carlo sobre cuadrados de dimensión arbitraria.

*/
func main() {
	// Ejercicio1:
	fmt.Println("----------------------------------------")
	fmt.Println("EJERCICIO 1:")
	fmt.Println(helloworld())
	fmt.Println(funciones.Chino())
	fmt.Println("----------------------------------------")
	// Ejercicio2:
	fmt.Println("EJERCICIO 2:")
	fmt.Println("Estos son los 10 bytes de esta cadena: abcdefghijklmnopqrstuvwxyz")
	fmt.Println(RandBytes(10,"abcdefghijklmnopqrstuvwxyz"))
	// Aqui son 10 bytes pero tenemos caracters especiales que ocupan mas de 1 byte.
	// A veces falla con una ? porque coge un byte que realmente es de una runa que ocupa mas de 1 byte.
	fmt.Println("Estos son los 10 bytes de esta cadena: a你cd$%&efghijklmnopq&/()%$·")
	fmt.Println(RandBytes(10,"a你cd$%&efghijklmnopq&/()%$·"))
	fmt.Println("Estos son los 100 bytes de esta cadena: abcdefghijklmnopqrstuvwxyz")
	fmt.Println(RandBytes(100,"abcdefghijklmnopqrstuvwxyz"))
    	fmt.Println("Estas son las 10 runes de esta cadena: a你cd$%&efghijklmnopq&/()%$·")
	fmt.Println("Usando el método RandRunes pero no es 100% exacto: ", RandRunes(10,"a你cd$%&efghijklmnopq&/()%$·"))
	fmt.Println("Usando el método RandRunes2 generado por Javi Valdez: ", RandRunes2(10,"a你cd$%&efghijklmnopq&/()%$·"))
        substring, err := runasAleatorias(10,"a你cd$%&efghijklmnopq&/()%$·")
	if err!=nil { 
		fmt.Println("RandRunes fallo usando el método runasAleatorias?", err)
        }else{
	      fmt.Println("RandRunes usando el método runasAleatorias:", substring)
        }

        substring2, err2 := runasAleatorias(100,"a你cd$%&efghijklmnopq&/()%$·")
        if err2!=nil { 
                fmt.Println("RandRunes fallo usando el método runasAleatorias?", err2)
        }else{
              fmt.Println("RandRunes usando el método runasAleatorias:", substring2)
        }

    	fmt.Println("----------------------------------------")
    	// Ejercicio3:
	fmt.Println("EJERCICIO 3:")
	// Normalmente calculamos π en 2 dimensiones el parametro de entrada es el tamaño de circulo (radio).
	fmt.Printf("Aproximación de π en cuadrados con %d de radio: %f\n",1, πMonteCarlo(1))
    fmt.Printf("Aproximación de π en cuadrados con %d de radio: %f\n",10, πMonteCarlo(10))
	fmt.Println("----------------------------------------")
}
