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
	getArray()
	getSlice()
	ifTest()
	forTest()

	s := 1000
	while(func() bool {
		return s > 0
	}, func() {
		s--
	})

	fmt.Print("Ingresa dos numeros: ")
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(calculate(a, b))

	guessTheNumber()
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

func getArray() {
	var arr1 [2]string
	arr2 := [3]int{1, 2, 3}

	arr1[0] = "array"
	arr1[1] = "array2"

	fmt.Println(arr1, arr2)
}

func ifTest() {
	var number = 0

	fmt.Print("Ingresa un número del 1 al 10: ")
	fmt.Scanf("%d", &number)

	if number%2 == 0 {
		fmt.Println("El número es par")
	} else {
		fmt.Println("El número es impar")
	}

	if number2 := 3; number2 == 3 {
		fmt.Println("Entró al condicional")
	}
}

func forTest() {
	for i := 0; i < 5; i++ {
		fmt.Println("[FOR] ", i)
	}

	c := 100
	for c > 0 {
		c -= 10
		fmt.Println("[FOR] Solo con una condición de c > 0, ", c)
	}

	s := 1000
	for {
		s--
		if s == 0 {
			// Termina el for "infinito"
			break
		}
	}
}

func while(cond func() bool, body func()) {
	for cond() {
		body()
	}
}

func getSlice() {
	var slice1 []string

	slice1 = append(slice1, "Mi", "slice", "1")

	fmt.Println(slice1)
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

func guessTheNumber() {
	var guess int
	var number = 8

	fmt.Print("Ingrese un número entre 1 y 10: ")
	fmt.Scanf("%d", &guess)

	if guess == number {
		fmt.Println("Ha adivinado el número.")
	} else {
		fmt.Println("No ha adivinado el número.")
	}
}
