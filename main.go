package main

import "fmt"

const helloWorld string = "Hola %s %s, bienvenido al fascinante mundo de go"
const testConst = "Test"

func main() {
	var name string
	name = "Pablo"
	lastname := "<Modificar con mi apellido>"
	var number = 100
	var (
		a = 1
		b = 2
		c = 3
	)

	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	fmt.Printf(helloWorld, name, lastname)

	fmt.Println("Hola mundo")
	fmt.Println(number, a, b, c)
}
