> Navega por la tabla de contenido

- [Paquete fmt](#paquete-fmt)
  - [Println](#println)
  - [Printf](#printf)
  - [Sprintf](#sprintf)
- [Funciones](#funciones)
  - [Retornar más de un valor](#retornar-más-de-un-valor)
  - [Recibir solo un valor](#recibir-solo-un-valor)
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

Para saber el tipo del dato con la librería fmt:

`fmt.Printf("%T\n", variable)`

## Funciones

Se utilizan para reutilizar código y que todo sea más corto y limpio para leer.

De hecho la función `main`, es la primera que definimos. Dentro de la misma también hemos utilizado otras funciones, como por ejemplo, las de impresión de contenido, por medio de las funciones, `Println`, `Printf `o `Sprintf`, todo esto, con el objetivo de reutilizar la función para evitar escribir las salidas manualmente por medio de métodos complejos.

Ejemplo de funciones: 

```go
package main

import "fmt"

func suma(A, B int) int           { return A + B }
func resta(A, B int) int          { return A - B }
func multiplicacion(A, B int) int { return A * B }
func division(A, B int) int       { return A / B }

func main() {
    var A int = 20
    var B int = 10

    Suma := suma(A, B)
    Resta := resta(A, B)
    Multiplicacion := multiplicacion(A, B)
    Division := division(A, B)

    fmt.Println("Suma:", Suma)
    fmt.Println("Resta:", Resta)
    fmt.Println("Multiplicación:", Multiplicacion)
    fmt.Println("División:", Division)
}
```

```bash
go run calc_aritmetica.go

#Output
Suma: 30
Resta: 10
Multiplicacion: 200
Division: 2
```

Se declara la función con la palabra reservada `func`, seguido se le agrega un nombre, además, entre parentesis se le pasan parámetros, donde si más de un parámetro, en una secuencia tienen el mismo tipo de dato, este solo se resume a una vez.

```go
func suma(A, B int)
```

Si esta función retorna algún valor, se definen la cantidad de variables y su tipo de dato, por ejemplo: 

```go
func cuadrado(n int) (x, y int)
```

### Retornar más de un valor

Se deja el ejemplo del cuadrado, donde se pasa un valor y se busca retornar ese mismo valor y el cuadrado a la vez.

```go
package main

import "fmt"

func cuadrado(n int) (x, y int) {
    return n, n * n
}

func main() {
    var N int = 20

    Value, Cuadrado := cuadrado(N)
    fmt.Printf("Value: %d, Cuadrado %d\n", Value, Cuadrado)

    /*Solo obtiene un valor*/
    _, Value = cuadrado(N)
    fmt.Printf("Cuadrado: %d\n", Value)
}
```

```bash
go run double_return.go

#Output 
Value: 20, Cuadrado 400
Cuadrado: 400
```

Se muestra como obtener los dos valores, cada uno en la variable correspondiente.

```go
Value, Cuadrado := cuadrado(N)
fmt.Printf("Value: %d, Cuadrado %d\n", Value, Cuadrado)
```

### Recibir solo un valor

Si tenemos funciones donde retornarmos más de un valor, pero solo necesitamos recibir uno de ellos, es tan sencillo como omitir el valor que viene en cierto orden con el caracter `_`.

```go
/*Solo obtiene un valor*/
_, Value = cuadrado(N)
fmt.Printf("Cuadrado: %s\n", Value)
```

**Ejercicio**

Obtener el área de un círculo, rectángulo y trapecio usando funciones.

```go
package main

import (
    "fmt"
    "math"
)

func áreaCírculo(radio float64) float64 {
    return math.Pi * radio * radio
}
func áreaRectángulo(base, altura float64) float64 {
    return base * altura
}

func áreaTrapecio(B, b, h float64) float64 {
    return h * (B + b) / 2
}

func main() {
    fmt.Printf("Círculo: %.2f \n", áreaCírculo(2))
    fmt.Printf("Rectángulo: %.2f \n", áreaRectángulo(5, 10))
    fmt.Printf("Trapecio: %.2f \n", áreaTrapecio(10, 5, 3))
}
```

```bash
go run area_rectangulo_circulo_trapecio.go

#Output
Círculo: 12.57 
Rectángulo: 50.00 
Trapecio: 22.50
```

## Recursos

[fmt package - fmt - pkg.go.dev](https://pkg.go.dev/fmt)

`YouTube Vídeos`

[✅¿Cual es el mejor Framework web de Golang? - YouTube](https://www.youtube.com/watch?v=Pq6rkq7iuHM)

[CÓMO hacer un CRUD ► 🎁 crud GOlang MySql PASO a PASO - YouTube](https://www.youtube.com/watch?v=G58gN0lIbyI)
