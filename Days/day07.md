> Navega por la tabla de contenido

- [Structs](#structs)
	- [Introducción](#introducción)
		- [Utilidad de structs](#utilidad-de-structs)
		- [Características](#características)
	- [Ejercicios](#ejercicios)
	- [Recursos](#recursos)

# Structs

**Un `struct` en Go es un tipo de datos que permite agrupar varios valores de diferentes tipos en una sola estructura.** Los valores se llaman campos y cada campo tiene un nombre y un tipo. Puedes declarar un struct vacío o inicializarlo con valores al momento de su declaración. También puedes crear un struct anónimo y asignarlo a una variable. Los campos de un struct se acceden utilizando el operador `.`.

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

**Declarar una estructura**

Para declarar un `struct` en Go, primero debes definir el tipo de struct utilizando la palabra clave "type" seguida del nombre del struct, seguido de la estructura entre llaves {}. Los campos del struct se definen dentro de las llaves, con su nombre y tipo.

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

También es posible crear una variable de un struct vacío utilizando la sintaxis "nombre_del_tipo{}", y asignar valores a sus campos más adelante.

## Ejercicios

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

## Recursos
