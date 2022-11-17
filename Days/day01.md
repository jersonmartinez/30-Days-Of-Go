> Navega por la tabla de contenido

- [Paquete fmt](#paquete-fmt)
  - [Println](#println)
  - [Printf](#printf)
  - [Sprintf](#sprintf)
- [Variables de entorno](#variables-de-entorno)
  - [Escribir y leer variables de entorno](#escribir-y-leer-variables-de-entorno)
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

## Variables de entorno

Son las que están definidas en el sistema y que desde código Go pueden ser accedidas, siempre y cuando se importen las librerías necesarias, en este caso, se utilizará `os`.

```go
package main

import (
    "fmt"
    "os"
)

func main() {

    var (
        home   = os.Getenv("HOME")
        user   = os.Getenv("USER")
        gopath = os.Getenv("GOPATH")
    )

    fmt.Println("HOME:", home)
    fmt.Println("USER:", user)
    fmt.Println("GOPATH:", gopath)
}
```

```bash
go run environment.go

#Output 
HOME: C:\Users\Root
USER: 
GOPATH: C:\Users\Root\go
```

### Escribir y leer variables de entorno

Se utiliza la librería `os` y `log`, para registrar la actividad.

```go
package main

import (
    "log"
    "os"
)

func main() {

    key := "NEW_VAR"
    value := "This is the first value in this environment value!"

    os.Setenv(key, value)

    val := GetEnvDefault(key, value)

    log.Println("\nEl valor es: " + val)

    os.Unsetenv(key)

    val = GetEnvDefault(key, value)

    log.Println("\nEl valor por default es: " + val)

}

// GetEnvDefault buscar las vars
func GetEnvDefault(key, defVal string) string {
    val, ex := os.LookupEnv(key)
    if !ex {
        return defVal
    }
    return val
}
return defVal
    }
    return val
}
```

```bash
go run set_and_get_environment.go

#Output
2022/03/02 15:33:41 
El valor es: This is the first value in this environment value!
2022/03/02 15:33:41 
El valor por default es: This is the first value in this environment value!
```

## Funciones

Se utilizan para reutilizar código y que todo sea más corto y limpio para leer.

De hecho, la función `main`, es la primera que definimos. Dentro de la misma también hemos utilizado otras funciones, como, por ejemplo, las de impresión de contenido, por medio de las funciones, `Println`, `Printf `o `Sprintf`, todo esto, con el objetivo de reutilizar la función para evitar escribir las salidas manualmente por medio de métodos complejos.

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

Si tenemos funciones donde retornaremos más de un valor, pero solo necesitamos recibir uno de ellos, es tan sencillo como omitir el valor que viene en cierto orden con el caracter `_`.

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

[Inicia tu carrera de programador en Go](https://www.crashell.com/estudio/inicia_tu_carrera_de_programador_en_go)

[fmt package - fmt - pkg.go.dev](https://pkg.go.dev/fmt)

[Golang: Obtener y establecer variables de entorno con valores predeterminados | zeroidentidad](https://awebytes.wordpress.com/2019/08/25/golang-obtener-y-establecer-variables-de-entorno-con-valores-predeterminados/)

[Effective Go - The Go Programming Language](https://go.dev/doc/effective_go#constants)

`YouTube Vídeos`

[¿Cual es el mejor Framework web de Golang?](https://www.youtube.com/watch?v=Pq6rkq7iuHM)

[CÓMO hacer un CRUD - crud GOlang MySql PASO a PASO](https://www.youtube.com/watch?v=G58gN0lIbyI)
