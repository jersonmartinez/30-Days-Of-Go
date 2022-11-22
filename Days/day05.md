![image](https://d3nykx067fw7ke.cloudfront.net/pages/estudio/articles/43/uso-de-slices-en-go.png)

> Navega por la tabla de contenido

- [Slices](#slices)
  - [Introducción](#introducción)
  - [Uso de Slices](#uso-de-slices)
    - [Ejemplos básicos](#ejemplos-básicos)
    - [Append](#append)
      - [Añadiendo nuevos elementos](#añadiendo-nuevos-elementos)
      - [Tipos de datos numéricos y strings](#tipos-de-datos-numéricos-y-strings)
      - [Insertar datos usando loop for](#insertar-datos-usando-loop-for)
    - [Remove](#remove)
      - [Eliminar valores](#eliminar-valores)
      - [Creando función para remover elemento](#creando-función-para-remover-elemento)
    - [Copy](#copy)
      - [Copiar valores de un slice a otro](#copiar-valores-de-un-slice-a-otro)
  - [Slice of slices](#slice-of-slices)
  - [Recursos](#recursos)

# Slices

## Introducción

Si bien los arreglos son una parte importante del lenguaje, es mas común el uso de slices. Un slice en cierta forma es un segmento de un arreglo.

El uso de arreglos tiene ciertas limitaciones, la mas importante es que una vez que se alcanza la capacidad del arreglo, no es posible agregar mas valores a este. En ese sentido los slices son mas flexibles. Es posible agregar, remover y copiar elementos de un slice a otro.

## Uso de Slices

### Ejemplos básicos

Asignando un rango de datos en un slice.

```go
package main

import "fmt"

func main() {
    primes := [6]int{2, 3, 5, 7, 11, 13}

    var s []int = primes[0:4]
    fmt.Println(s)
}
```

```bash
go run numeric.go

#Output
[2 3 5 7]
```

Crea un slice, usando la función `make`, anunciando que serán strings y que el tamaño a reservar es de 2. Se le añaden valores y se imprime.

```go
slice_name := make([]type, length, capacity)
```

```go
package main

import "fmt"

func main() {
    var marcasDeCoches = make([]string, 2)

    marcasDeCoches[0] = "Aprendiendo"
    marcasDeCoches[1] = "Go"

    fmt.Println(marcasDeCoches)
}
```

```bash
go run add_values.go

#Output
[Aprendiendo Go]
```

### Append

#### Añadiendo nuevos elementos

```go
package main

import "fmt"

func main() {
    var marcasDeCoches = make([]string, 2)
    marcasDeCoches[0] = "Aprendiendo"
    marcasDeCoches[1] = "Go"

    marcasDeCoches = append(marcasDeCoches, "en", "la", "nube")
    fmt.Println(marcasDeCoches)
}
```

```bash
go run append.go

#Output
[Aprendiendo Go en la nube]
```

#### Tipos de datos numéricos y strings

```go
package main

import (
    "fmt"
)

func main() {
    // Creating int and string slices
    s1 := []int{10, 20, 30}
    s2 := []string{"Hello", "World"}

    // Printing types and values of slices
    fmt.Printf("%T, %v\n", s1, s1)
    fmt.Printf("%T, %q\n", s2, s2)

    // Appending the elements
    s1 = append(s1, 40, 50)
    s2 = append(s2, "How are you?", "Boys", "and", "Girls")

    fmt.Println("After appending...")
    fmt.Printf("%T, %v\n", s1, s1)
    fmt.Printf("%T, %q\n", s2, s2)
}
```

```bash
go run append_adding_slices.go

#Output
[]int, [10 20 30]
[]string, ["Hello" "World"]
After appending...
[]int, [10 20 30 40 50]
[]string, ["Hello" "World" "How are you?" "Boys" "and" "Girls"
```

#### Insertar datos usando loop for

Como ya mencionamos, un arreglo es una colección de valores del mismo  tipo que tiene una longitud fija y no puede tener un método para cambiar su tamaño.

Por otro lado un slice es una abstracción de un arreglo, es decir que utiliza un arreglo tras bambalinas para funcionar, pero provee ciertas funciones adicionales.

```go
package main

import "fmt"

func main() {

    mySlice := []int{}

    for i := 0; i < 100; i += 1 {
        mySlice = append(mySlice, i+1)
    }

    fmt.Printf("%+v", mySlice)
}
```

```bash
go run append_using_loop_for.go

#Output
[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100]
```

### Remove

#### Eliminar valores

Como ya mencionamos, un arreglo es una colección de valores del mismo tipo que tiene una longitud fija y no puede tener un método para cambiar su tamaño.

Por otro lado un slice es una abstracción de un arreglo, es decir que utiliza un arreglo tras bambalinas para funcionar, pero provee ciertas funciones adicionales.

```go
package main

import "fmt"

func main() {
    OS := []string{"Ubuntu", "Debian", "Windows 10", "Windows 11", "Fedora", "Mac OS X"}
    fmt.Println(OS)

    indice := 2 // Windows 10

    OS = append(OS[:indice], OS[indice+1:]...)
    fmt.Println(OS)
}
```

```bash
go run delete_slice.go

#Output
[Ubuntu Debian Windows 10 Windows 11 Fedora Mac OS X]
[Ubuntu Debian Windows 11 Fedora Mac OS X]
```

#### Creando función para remover elemento

```go
package main

import "fmt"

func remove(slice []string, s int) []string {
    return append(slice[:s], slice[s+1:]...)
}

func main() {
    OS := []string{"Ubuntu", "Debian", "Windows 10", "Windows 11", "Fedora", "Mac OS X"}
    fmt.Println(OS)

    OS = remove(OS, 2)

    fmt.Println(OS)
}
```

```bash
go run delete_slice_function.go

#Output
[Ubuntu Debian Windows 10 Windows 11 Fedora Mac OS X]
[Ubuntu Debian Windows 11 Fedora Mac OS X]
```

### Copy

Para copiar todos los elementos de un slice existe la función `copy`.  Para copiar valores de un slice a otro, un slice puede ser inicializado mediante otro del mismo tipo, ya que no es posible copiar los valores de un slice de enteros a otro de cadenas de strings. Una copia implica que los valores permanecen aun si el slice original es modificado o destruido.

#### Copiar valores de un slice a otro

- Se crea un slice llamado `OS`.
- Se genera un contenedor para crear una copia, con la dimensión de `OS` mediante `make([]string, len(OS))`.
- Se copian los valores de `OS` en `OS_Copy` mediante `copy(OS_Copy, OS)`.
- Se elimina el elemento `Debian` de `OS`.
- Los valores de `OS_Copy` no se han alterado.

```go
package main

import "fmt"

func main() {
    OS := []string{"Ubuntu", "Debian", "Windows 10", "Windows 11", "Fedora", "Mac OS X"}
    var OS_Copy = make([]string, len(OS))

    fmt.Println(OS)

    copy(OS_Copy, OS)

    OS = append(OS[:1], OS[2:]...)

    fmt.Println(OS_Copy)
    fmt.Println(OS)
}
```

```bash
go run copy_slices.go

#Output
[Ubuntu Debian Windows 10 Windows 11 Fedora Mac OS X]
[Ubuntu Debian Windows 10 Windows 11 Fedora Mac OS X]
[Ubuntu Windows 10 Windows 11 Fedora Mac OS X]
```

## Slice of slices

```go
package main

import (
    "fmt"
)

// main function
func main() {

    /*Declarando un slice bidimensional*/
    slice_of_slices := make([][]int, 3)

    for i := 0; i < 3; i++ {
        /*Declarando un slice dentro de cada elemento del slice padre*/
        slice_of_slices[i] = make([]int, 3)

        /*Asignando valores a cada elemento del slice*/
        for j := 0; j < 3; j++ {
            slice_of_slices[i][j] = i * j
        }
    }

    /*Imprimiendo los slices hijos*/
    for i := 0; i < 3; i++ {
        fmt.Println(slice_of_slices[i])
    }

    /*Imprimiendo todo el slice*/
    fmt.Println("Slice of slices: ", slice_of_slices)
}
```

```bash
go run slice_of_slices.go

#Output
[0 0 0]
[0 1 2]
[0 2 4]
Slice of slices:  [[0 0 0] [0 1 2] [0 2 4]]
```

## Recursos

[Go by Example: Slices](https://gobyexample.com/slices)

[A Tour of Go](https://go.dev/tour/moretypes/7)

https://www.includehelp.com/golang/append-function-with-examples.aspx
