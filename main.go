package main

import "fmt"

const helloWorld string = "Hola %s %s, bienvenido al fascinante mundo de go\n"
const testConst = "Test"

func main() {
	lastname := "<Modificar con mi apellido>"
	name := getName()

	number := sum(50, 50)
	a, b, c := getVariables()
	f32, f64 := getFloat()
	stringUTF8 := getUnicode()

	fmt.Printf(helloWorld, name, lastname)

	fmt.Println("Hola mundo")
	fmt.Println(number, a, b, c)
	fmt.Println(f32, f64)
	fmt.Println("Cadena con UTF8: ", stringUTF8)
	fmt.Println(string("hola"[1]), ". Codigo ASCII: ", "hola"[1])
	fmt.Println("Cantidad de letras que tiene hola: ", len("hola"))

	fmt.Print("Ingresa dos numeros: ")
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(calculate(a, b))
}

func getName() string {
	var name string
	name = "Sin nombre"

	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)

	return name
}

func getVariables() (int, int, int) {
	return 1, 2147000000, 903131313131311313
}

func getFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

func getUnicode() string {
	return "もしもし！"
}

func calculate(a int, b int) (int, int) {
	return sum(a, b), difference(a, b)
}

func sum(a int, b int) int {
	return a + b
}

func difference(a int, b int) int {
	return a - b
}
