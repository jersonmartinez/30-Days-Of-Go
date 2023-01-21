> Navega por la tabla de contenido

- [Structs](#structs)
	- [Introducción](#introducción)
		- [Utilidad de structs](#utilidad-de-structs)
		- [Características](#características)
	- [Ejemplos prácticos](#ejemplos-prácticos)
		- [Declarar e instanciar una estructura](#declarar-e-instanciar-una-estructura)
		- [Almacenar e imprimir datos](#almacenar-e-imprimir-datos)
		- [Estructura Global](#estructura-global)
		- [Estructura Local](#estructura-local)
		- [Funciones dentro de estructuras](#funciones-dentro-de-estructuras)
		- [Asociar método dentro de una estructura](#asociar-método-dentro-de-una-estructura)
	- [Preguntas interesantes](#preguntas-interesantes)
	- [Recursos](#recursos)

# Structs

**Un struct en Go es un tipo de datos que permite agrupar varios valores de diferentes tipos en una sola estructura.** Los valores se llaman campos y cada campo tiene un nombre y un tipo. Puedes declarar un struct vacío o inicializarlo con valores al momento de su declaración. También puedes crear un struct anónimo y asignarlo a una variable. Los campos de un struct se acceden utilizando el operador `.`.

## Introducción

### Utilidad de structs

Los structs en Go son útiles para agrupar datos relacionados en una sola estructura de datos y proporcionar una manera de trabajar con ellos de manera organizada y eficiente. Al utilizar structs, puedes crear tipos de datos personalizados que se ajusten mejor a tus necesidades y proporcionen una mayor legibilidad y claridad en tu código.

Además, los structs en Go son útiles para crear estructuras de datos complejas, como árboles, listas enlazadas y grafos, así como para representar objetos y entidades en una aplicación. También son útiles para trabajar con datos de una base de datos o una API, ya que puedes crear structs que se ajusten a la estructura de los datos devueltos.

**Ventajas de usar structs en Go:**

- Mejora la legibilidad del código: Al tener datos relacionados agrupados en un solo objeto, es más fácil de entender y mantener.
- Reutilización de código: Una vez que se define un struct, puede ser reutilizado en diferentes partes del código.
- Mejora la seguridad del código: Los structs permiten limitar el acceso a los datos mediante la encapsulación de los campos.
- Mejora la eficiencia del código: Al tener todos los datos relacionados en un solo objeto, se reducen las llamadas a funciones y la memoria utilizada.

En resumen, los structs en Go son una herramienta muy importante para organizar y manejar datos de manera eficiente y legible en tu código, y son esenciales para crear aplicaciones complejas y escalables.

### Características

**Los structs en Go tienen las siguientes características:**

- Campos: Los structs en Go tienen campos, que son valores de diferentes tipos agrupados en una sola estructura. Cada campo tiene un nombre y un tipo.
- Inicialización: Los structs pueden ser declarados vacíos o inicializados con valores al momento de su declaración.
- Acceso a campos: Los campos de un struct se acceden utilizando el operador ".".
- Métodos: Los structs en Go pueden tener métodos asociados a ellos, lo cual permite a los structs tener comportamiento y poder ser tratados como objetos.
- Composición: Los structs en Go pueden contener otros structs, lo cual permite crear estructuras de datos complejas.
- Encapsulamiento: Los structs en Go son privados por defecto, es decir, no se pueden acceder a sus campos desde fuera de la paquete donde se declaran.
- Interfaz: Los structs en Go pueden implementar una interfaz, lo cual permite tratar diferentes tipos de structs de la misma manera.
- Valores o punteros: Los structs en Go pueden ser pasados como valores o como punteros, lo cual tiene implicaciones en el comportamiento y rendimiento.

## Ejemplos prácticos

### Declarar e instanciar una estructura

Para declarar un struct en Go, primero debes definir el tipo de struct utilizando la palabra clave "type" seguida del nombre del struct, seguido de la estructura entre llaves {}. Los campos del struct se definen dentro de las llaves, con su nombre y tipo.

La sintaxis básica sería la siguiente: 

```go
type StructureName struct {
  // structure definition 
}
```

Aquí tienes un ejemplo de cómo declarar un struct de un punto en un plano cartesiano:

```go
type Point struct {
    x float64
    y float64
}
```

En este ejemplo, se define un struct llamado "Point" con dos campos, "x" y "y", ambos del tipo float64.

Una vez que se define el struct, puedes crear variables de ese tipo utilizando la sintaxis "nombre_del_tipo{valor_del_campo_1, valor_del_campo_2, ...}", como se ilustra a continuación:

```go
p := Point{x: 3.0, y: 4.0}
```

En este caso se crea una variable "p" de tipo "Point" con los valores 3.0 y 4.0 para los campos "x" y "y" respectivamente.

```go
package main

import "fmt"

type Point struct {
	x float64
	y float64
}

func main() {
	p := Point{x: 3.0, y: 4.0}

	fmt.Println("X:", p.x)
	fmt.Println("Y:", p.y)
}
```

```bash
go run struct_declaration.go 

# Output
X: 3
Y: 4
```

También es posible crear una variable de un struct vacío utilizando la sintaxis "nombre_del_tipo{}", y asignar valores a sus campos más adelante.

### Almacenar e imprimir datos

Almacenar nombre y edad de una persona en una estructura e imprimirla.

```go
package main

import "fmt"

type Person struct {
    name string
    age int
}

func main() {
    p := Person{name: "Jerson Martínez", age: 26}

    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
}
```

En este caso, se inicializa la variable "p" con los valores "Jerson Martínez" y 26 para los campos "name" y "age" respectivamente.

```bash
go run add_data_and_print.go 

# Output
Name: Jerson Martínez
Age: 26
```

### Estructura Global

En este ejemplo, primero se define un `struct` global `Person` con dos campos, *name* y *age*, ambos del tipo `string` y `int` respectivamente.

Se crea la estructura fuera de la función main, lo que pasa a ser global, logrando que la función `printPerson` pueda acceder instanciándola desde sus parámetros por medio de (`person Person`).

```go
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var p Person
	p.name = "Jerson Martínez"
	p.age = 26
	fmt.Println("Información de persona en la función main:")
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)

	fmt.Println("\nInformación de persona en la función printPerson:")
	printPerson(p)
}

func printPerson(person Person) {
	fmt.Println("Name:", person.name)
	fmt.Println("Age:", person.age)
}
```

```bash
go run global_struct.go 

# Output
Información de persona en la función main:
Name: Jerson Martínez
Age: 26

Información de persona en la función printPerson:
Name: Jerson Martínez
Age: 26
```

### Estructura Local

En este ejemplo, el `struct` *Person* se define dentro de la función `main`, por lo que es un `struct` local y su alcance solo es dentro de la función `main`. En la función `main` se crea una variable `p` de tipo *Person* y se inicializan sus campos *name* y *age* con los valores "Jerson Martínez" y 26 respectivamente.

Es importante destacar que el `struct` *Person* solo es accesible y utilizable dentro de la función `main` y no estará disponible en otras funciones o archivos del paquete.

```go
package main

import "fmt"

func main() {
	type Person struct {
		name string
		age  int
	}

	p := Person{name: "Jerson Martínez", age: 26}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
}
```

```bash
go run local_struct.go 

# Output
Name: Jerson Martínez
Age: 26
```

En general, es recomendado declarar `structs` fuera de las funciones para tener un mejor control y organización de la aplicación y poder reutilizar el `struct` en diferentes partes del código.

### Funciones dentro de estructuras

En Go, las estructuras (`structs`) pueden tener funciones asociadas, conocidas como **métodos**. Estos métodos son funciones que tienen acceso a los campos de la estructura a la que están asociadas.

```go
package main

import "fmt"

// Define la estructura "Person" con dos campos, "name" y "age"
type Person struct {
	name string
	age  int
}

// Define un método "PrintInfo" asociado a la estructura "Person"
// El método recibe una variable de tipo "Person" y utiliza los campos de esa variable para imprimir en pantalla
func (p Person) PrintInfo() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
}

func main() {
	// Crea una variable "p" de tipo "Person" y asigna valores a los campos "name" y "age"
	p := Person{name: "Jerson Martínez", age: 26}
	// Llama al método "PrintInfo" de la variable "p"
	p.PrintInfo()
}
```

```bash
go run functions_inside_struct.go 

# Output
Name: Jerson Martínez
Age: 26
```

### Asociar método dentro de una estructura

Cómo asociar una función anónima a un campo dentro de una estructura en Go.

La estructura `Person` tiene tres campos: *name*, *age* e `Info`, del tipo `string`, `int` y `PrintInfo` respectivamente. El campo `Info` es una función del tipo `PrintInfo` que recibe dos argumentos de tipo `string` y `int` y devuelve un `string`.

En la función *main*, se crea una variable *result* de tipo `Person` y se asignan valores a sus campos, incluyendo una función anónima al campo `Info` que utiliza la función `fmt.Sprintf` para concatenar el valor de los campos *name* y *age* y devuelve un `string` con el nombre y la edad del usuario.

Finalmente, en la función *main* se imprimen el nombre y el nombre y edad del usuario mediante el uso del campo *name* y el campo `Info` respectivamente.

```go
package main

import "fmt"

// Declare a new type "PrintInfo" as a function
// that takes two arguments, a string and an int
// and returns a string
type PrintInfo func(string, int) string

// Declare a struct "Person" with fields "name", "age"
// and "Info" of type string, int and PrintInfo respectively
type Person struct {
	name string
	age  int
	Info PrintInfo
}

func main() {

	// Create a variable "result" of type "Person" and
	// initialize its fields with values
	result := Person{
		name: "Jerson Martínez",
		age:  26,
		// An anonymous function is assigned to the "Info" field
		// that takes the "name" and "age" fields of the "Person" struct and
		// concatenates them using the fmt.Sprintf function
		Info: func(name string, age int) string {
			return fmt.Sprintf("%s %d", name, age)
		},
	}

	// Print the "name" and "Info" fields of the "result" variable
	fmt.Println("Name: ", result.name)
	fmt.Println("Name and age:", result.Info(result.name, result.age))
}
```

```bash
go run associate_functions_inside_struct.go 

# Output
Name:  Jerson Martínez
Name and age: Jerson Martínez 26
```

En resumen, el algoritmo define una estructura Person con un campo `Info` de tipo `PrintInfo`, que es una función anónima que recibe dos argumentos de tipo `string` y `int` y devuelve un `string` con la concatenación de esos dos valores. Luego se asigna una instancia de `Person` con valores y se utiliza el campo `Info` para obtener el `string` concatenado de nombre y edad.

## Preguntas interesantes

**¿Se pueden definir structs dentro y fuera de la función main en Go?**

**R:** Sí, es posible definir structs tanto dentro como fuera de la función main en Go.

**¿Cuál es la diferencia de declarar un struct dentro y fuera de la función main en Go?**

**R:** La diferencia entre declarar un struct dentro y fuera de la función `main` en Go es la **visibilidad y alcance** de ese struct.

Cuando se declara un struct fuera de cualquier función, se dice que es un struct global y su alcance es todo el paquete. Esto significa que cualquier función o archivo dentro del mismo paquete puede acceder y utilizar ese struct. Por otro lado, si se declara un struct dentro de una función, se dice que es un struct local y su alcance solo es dentro de esa función. Esto significa que solo se puede acceder y utilizar ese struct dentro de esa función donde se declaró.

En general, es recomendado declarar structs fuera de las funciones para tener un mejor control y organización de la aplicación y poder reutilizar el struct en diferentes partes del código. Sin embargo, en algunos casos específicos puede ser útil utilizar structs locales para limitar el alcance de la información o para trabajar con un conjunto de datos temporal.

## Recursos
