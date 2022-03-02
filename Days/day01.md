> Navega por la tabla de contenido

- [Paquete fmt](#paquete-fmt)
  - [Println](#println)
  - [Printf](#printf)
  - [Sprintf](#sprintf)
- [Condicionales...](#condicionales)
- [Recursos](#recursos)

## Paquete fmt

Ciertamente Go se parece mucho a C, puesto que el paquete fmt tiene bastante similitud con stdio.h de C. En ambos casos, implementan el formato de entrada y salida de datos. Igualmente, te sorprenderás de las similitudes de Go con C.

Revisaremos las 3 posibles funciones más importantes de `fmt`.

- Println

- Printf

- Sprintf

### Println

Imprime y agrega un saldo de línea.

```go
package main

import "fmt"

func main() {

    Fistname := "Jerson"
    Lastname := "Martínez"

    fmt.Println("Hola,", Fistname, Lastname)
}
```

```bash
go run Println.go

#Output
Hola, Jerson Martínez
```

### Printf

Imprime con un formato.

```go
package main

import "fmt"

func main() {

    Firstname := "Jerson"
    Lastname := "Martínez"

    Year := 2030

    /*
        Imprime en pantalla el mensaje y los valores de las variables
        %s: String
        %d: Entero
        %f: Flotante
        %t: Booleano

        %v: Valor de la variable
        %T: Tipo de dato de la variable

        %p: Dirección de memoria de la variable

        %b: Binario
        %c: Caracter
        %o: Octal
        %x: Hexadecimal

        %q: Cadena de caracteres
        %e: Notación científica
        %g: Notación científica
        %f: Notación científica

        %+v: Valor de la variable con su tipo de dato
        %#v: Valor de la variable con su tipo de dato
    */

    fmt.Printf("Hola, %s %s\n", Firstname, Lastname)
    fmt.Printf("Vengo desde el %d\n", Year)

    fmt.Printf("Dirección de memorias de las variables: Firstname %p, Lastname %p, Year %p\n", &Firstname, &Lastname, &Year)
}
```

```bash
go run Printf.go

#Output
Hola, Jerson Martínez
Vengo desde el 2030
Dirección de memorias de las variables: Firstname 0xc000088220, Lastname 0xc000088230, Year 0xc0000aa058
```

### Sprintf

Aplica un formato a una cadena de caracteres y devuelve la cadena formateada para ser almacenada en una variable y posteriormente esta pueda ser manejada.

```go
package main

import "fmt"

func main() {

    Fistname := "Jerson"
    Lastname := "Martínez"

    Message := fmt.Sprintf("Hola, %s %s", Fistname, Lastname)

    fmt.Println(Message)
}
```

```bash
go run Sprintf.go

#Output
Hola, Jerson Martínez
```

## Condicionales...

## Recursos

[fmt package - fmt - pkg.go.dev](https://pkg.go.dev/fmt)

`YouTube Vídeos`

[✅¿Cual es el mejor Framework web de Golang? - YouTube](https://www.youtube.com/watch?v=Pq6rkq7iuHM)

[CÓMO hacer un CRUD ► 🎁 crud GOlang MySql PASO a PASO - YouTube](https://www.youtube.com/watch?v=G58gN0lIbyI)
