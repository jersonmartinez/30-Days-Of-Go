> Navega por la tabla de contenido

- [Arrays](#arrays)
  - [Introducción](#introducción)
  - [Características](#características)
- [Ejercicios](#ejercicios)
  - [Elementos numéricos (Int)](#elementos-numéricos-int)
  - [Elementos de cadenas (Strings)](#elementos-de-cadenas-strings)
  - [Sin definición de elementos](#sin-definición-de-elementos)
  - [Asignar valores desde la definición](#asignar-valores-desde-la-definición)
  - [Obtener la longitud](#obtener-la-longitud)
  - [Imprimir arreglos mediante ciclos](#imprimir-arreglos-mediante-ciclos)
  - [Aplicando filtros](#aplicando-filtros)
- [Recursos](#recursos)

## Arrays

### Introducción

Una **array**, *arreglo o matriz*, es una colección de elementos del mismo tipo con un tamaño pre-fijado, aunque no siempre tiene que ser así tal y como veremos más adelante. 

### Características

- En **Go**, un **array** es un **tipo por referencia**.

- Los elementos que forman parte de un array, son almacenados de forma secuencial. 

- Una matriz no puede ser re-dimensionada de forma directa.

- Para obtener un elemento de un array, lo haremos a través de su índice.

- Los índices de un **array** en **Go** empiezan por 0, así que 0 sería su primer elemento, 1 el segundo, etc.

## Ejercicios

### Elementos numéricos (Int)

Se define un arreglo, indicando la cantidad de elementos que este tendrá, sin embargo, si en la asociación de los valores hay alguno que no se le agrega, Go por omisión lo asigna a 0.

```go
package main

import "fmt"

func main() {
    array := [5]int{10, 20, 30, 40}

    fmt.Println("Elements:", array)
    fmt.Println("Element [3]:", array[3])
}
```

```bash
go run numeric_elements.go

#Output
Elements: [10 20 30 40 0]
Element [3]: 40
```

### Elementos de cadenas (Strings)

Definimos un primer arreglo de 2 elementos, donde el primero contendrá el valor "Hello" y el segundo "world"; esto mismo para el segundo arreglo, solo que con algunas diferencias en su definición.

```go
package main

import "fmt"

func main() {
    /*First array*/
    var array1 [2]string
    array1[0] = "Hello"
    array1[1] = "World"

    fmt.Println(array1[0], array1[1])
    fmt.Println(array1)

    /*Second array*/
    array2 := [2]string{"Hello", "World"}

    fmt.Println("\n", array2[0], array2[1])
    fmt.Println(array2)
}
```

```bash
go run string_elements.go

#Output
Hello World
[Hello World]

 Hello World
[Hello World]
```

### Sin definición de elementos

Se definen 3 arreglos, dos donde se almacenan números enteros y el tercero donde se almacenan strings. La idea principal es ver que en el índice, donde se anuncia de cuántos elementos será el arreglo, está señalado con puntos suspensivos `...`, indicando que se desconoce el número de elementos, por ende, tendrá que ajustarse a los valores que se pasan.

```go
package main

import (
    "fmt"
)

func main() {
    var array1 = [...]int{1, 2, 3}
    array2 := [...]int{4, 5, 6, 7, 8}
    array_string := [...]string{"Hello", "World"}

    fmt.Println(array1)
    fmt.Println(array2)
    fmt.Println(array_string)
}
```

```bash
go run without_elements.go

#Output
[1 2 3]
[4 5 6 7 8]
[Hello World]
```

### Asignar valores desde la definición

Se declara un arreglo de 5 elementos, por ende serán: [0, 1, 2, 3, 4]. Además, se agregan valores tanto para el elemento 1 y 2. Esto se hace directamente como un sistema clave - valor: `1: 10, 2: 40`.

```go
package main

import "fmt"

func main() {
    array := [5]int{1: 10, 2: 40}

    fmt.Println(array)
}
```

```bash
go run initial_values.go

#Output
[0 10 40 0 0]
```

### Obtener la longitud

Con la función `len` es posible obtener el número delementos que tiene un arreglo, independientemente si este dato fue definido, de si es de un tipo de dato o de otro.

```go
package main

import "fmt"

func main() {
    arr1 := [4]string{"First", "Second", "Third", "Fourth"}
    arr2 := [...]int{1, 2, 3, 4, 5, 6}

    fmt.Println(len(arr1))
    fmt.Println(len(arr2))
}
```

```bash
go run get_length.go

#Output
4
6
```

### Imprimir arreglos mediante ciclos

El ciclo for es especial en Go, porque se puede implementar de muchas formas, sustituyendo así algunos ciclos como el while, do while, foreach que si existen en otros lenguajes, pero que acá con el for, basta y sobra.

```go
package main

import "fmt"

func main() {
    intArray := [5]int{10, 20, 30, 40, 50}

    fmt.Println("\n---------------Example 1--------------------\n")
    for i := 0; i < len(intArray); i++ {
        fmt.Println(intArray[i])
    }

    fmt.Println("\n---------------Example 2--------------------\n")
    for index, element := range intArray {
        fmt.Println(index, "=>", element)

    }

    fmt.Println("\n---------------Example 3--------------------\n")
    for _, value := range intArray {
        fmt.Println(value)
    }

    j := 0
    fmt.Println("\n---------------Example 4--------------------\n")
    for range intArray {
        fmt.Println(intArray[j])
        j++
    }
}
```

```bash
go run print_with_for.go

#Output
---------------Example 1--------------------
10
20
30
40
50

---------------Example 2--------------------
0 => 10
1 => 20
2 => 30
3 => 40
4 => 50

---------------Example 3--------------------
10
20
30
40
50

---------------Example 4--------------------
10
20
30
40
50
```

### Aplicando filtros

```go
package main

import "fmt"

func main() {
    countries := [...]string{"India", "Canada", "Japan", "Germany", "Italy"}

    fmt.Printf("Countries: %v\n", countries)

    fmt.Printf(":2 %v\n", countries[:2])

    fmt.Printf("1:3 %v\n", countries[1:3])

    fmt.Printf("2: %v\n", countries[2:])

    fmt.Printf("2:5 %v\n", countries[2:5])

    fmt.Printf("0:3 %v\n", countries[0:3])

    fmt.Printf("Last element: %v\n", countries[len(countries)-1])

    fmt.Printf("All elements: %v\n", countries[0:len(countries)])
    fmt.Println(countries[:])
    fmt.Println(countries[0:])
    fmt.Println(countries[0:len(countries)])

    fmt.Printf("Last two elements: %v\n", countries[len(countries)-2:len(countries)])
}
```

```bash
go run filter_elements.go

#Output
Countries: [India Canada Japan Germany Italy]
:2 [India Canada]        
1:3 [Canada Japan]       
2: [Japan Germany Italy] 
2:5 [Japan Germany Italy]
0:3 [India Canada Japan] 
Last element: Italy      
All elements: [India Canada Japan Germany Italy]
[India Canada Japan Germany Italy]
[India Canada Japan Germany Italy]
[India Canada Japan Germany Italy]
Last two elements: [Germany Italy]
```

## Recursos

[Go Arrays](https://www.w3schools.com/go/go_arrays.php)

[Go by Example: Arrays](https://gobyexample.com/arrays)

[Arrays in Go with examples - golangprograms.com](https://www.golangprograms.com/go-language/arrays.html)
