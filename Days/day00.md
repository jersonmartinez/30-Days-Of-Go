# Introducción

- [Introducción](#Introducción)  
	- [Concepto](#Concepto)
	- [Instalando Go con Choco para Windows](#Instalando Go con Choco para Windows)

## Concepto

Golang es un lenguaje de programación de código abierto desarrollado por Google. Es un lenguaje compilado de tipo estático. Este lenguaje admite programación concurrente y también permite ejecutar múltiples procesos simultáneamente.

En el transcurso de esta lectura irás comprendiendo más acerca de ello, sin embargo, que mejor que si deseas profundizar sobre Golang revises [qué es Go](https://openwebinars.net/blog/que-es-go/).

## Instalando Go con Choco para Windows

En mi caso, estoy usando Windows como sistema operativo base, por lo que me será útil instalar Go de la manera más sencilla con Chocolatey. 

**Instalar Chocolatey**

Esta se instala desde PowerShell.

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
```

La instalación de Go mediante Choco se hace de la siguiente manera:

```powershell
choco install golang
```

También puede actualizar y desinstalar golang con Choco de la siguiente manera:

```powershell
choco upgrade golang
```

```
choco uninstall golang
```

## Instalando Go en GNU/Linux (Ubuntu)

Las siguientes instrucciones son obtenidas de la documentación oficial de [go.dev]([Download and install - The Go Programming Language](https://go.dev/doc/install)).

```bash
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.17.7.linux-amd64.tar.gz
```

Se agrega la variable de entorno al `$PATH`

```bash
export PATH=$PATH:/usr/local/go/bin
```

En ambos casos, para comprobar que la instalación se haya resuelto correctamente, se mira la versión de Go.

```bash
 go version
```

## Basic Syntax

Escribiéndo el `Hello, World!` como el clásico ejercicio inicial.

Los objetivos son simples:

- Escribir una función que retorne el string "Hello, World!".

- Compilar el script y ejecutarlo.

```go
package main

import "fmt"

func main() {
		fmt.Println("Hello, World!")
}
```

Compilación

```bash
go build hello_world.go
```

Ejecución

```bash
# Ejecutar sin compilar.
go run hello_world.go

# Ejecutar después de haber compilado.
./hello_world.exe

#Output
Hello, World!
```

## Variables and declaration

Las variables sirven para asociar una tira de datos con un sector de la memoria, siendo esta, un identificador. En palabras sencillas, las variables guardan datos.

Definir variables con las palabra reservada `var` y `const`.

**Declarando constantes:**

```go
package main

import "fmt"

func main() {
		/* Palabra reservada + Nombre de variable + Tipo de dato + Valor */
		const PI float64 = 3.14
		const PI2 = 3.14

		fmt.Println("Pi:", PI, "\nPi2: ", PI2)

}
```

```bash
go run constantes.go

#Output
Pi: 3.14 
Pi2:  3.14
```

**Declarando variables enteras**

```go
package main

import "fmt"

func main() {
		/* Declara la variable "Base", identifica el tipo de dato respecto al valor */
		Base := 10          //Primera forma
		var Altura int = 20 //Segunda forma
		var Area int        //Go no compila si las variables no son usadas

		fmt.Println("Base: ", Base, " | Altura: ", Altura, " | Area: ", Area)

		/*Si no hay valores definidos, Go le asigna un vacío.*/
		var a int     //0
		var b float64 //0.0
		var c string  //""
		var d bool    //false

		fmt.Println("a: ", a, " | b: ", b, " | c: ", c, " | d: ", d)
}
```

```bash
go run variables.go

#Output
Base:  10  | Altura:  20  | Area:  0
a:  0  | b:  0  | c:    | d:  false
```

**Ejercicio**

Calcular el cuadrado e imprimirlo.

```go
package main

import "fmt"

func main() {
		const BaseCuadrado = 10
		AreaCuadrado := BaseCuadrado * BaseCuadrado

		fmt.Println("El área del cuadrado es:", AreaCuadrado)
}
```

```bash
go run obtener_cuadrado.go

#Output
El área del cuadrado es: 100
```

**Ejercicio**

Operación aritmética.

```go
package main

import "fmt"

func main() {
		/* Palabra reservada + Nombre de variable + Tipo de dato + Valor */
		var A int = 10
		var B int = 20

		Suma := A + B
		Resta := A - B
		Multiplicacion := A * B
		Division := A / B
		Modulo := A % B

		A++
		B--

		fmt.Println("Suma:", Suma)
		fmt.Println("Resta:", Resta)
		fmt.Println("Multiplicacion:", Multiplicacion)
		fmt.Println("Division:", Division)
		fmt.Println("Modulo:", Modulo)
		fmt.Println("Incremental:", A)
		fmt.Println("Decremental:", B)
}
```

```bash
go run operaciones_aritmeticas.go 

#Output
Suma: 30
Resta: -10
Multiplicacion: 200
Division: 0
Modulo: 10
Incremental: 11
Decremental: 19
```

**Ejercicio**

Calcular el área de un rectángulo, trapecio y círculo.

```go
package main

import (
		"fmt"
		"math"
)

func main() {
		/* Declaracion de variables base */
		Base := 10
		BaseMayor := 20
		Altura := 30
		Radio := 20

		fmt.Println("El área del rectángulo es:", Base*Altura)
		fmt.Println("El área del trapecio es:", ((BaseMayor+Base)*Altura)/2)
		fmt.Println("El área del círculo es:", math.Pi*math.Pow(float64(Radio), 2))
}
```

```bash
go run calcular_area_rectangulo_trapecio_circulo.go

#Output
El área del rectángulo es: 300
El área del trapecio es: 450
El área del círculo es: 1256.6370614359173
```

**Ejercicio**

Crear un saludo modular. 

Para ello, crearemos 3 ficheros: 

- main.go

```go
package main

import "fmt"

func main() {
		fmt.Println("Hello, World!")

		hello()
		bye()
}
```

- hello.go

```go
package main

import "fmt"

func hello() {
		fmt.Println("Hello, friend!")
}
```

- bye.go

```go
package main

import "fmt"

func bye() {
		fmt.Println("Goodbye, friend!")
}
```

Ejecucion: 

```bash
go mod init Package/main
go run .

#Output
Hello, World!
Hello, friend!
Goodbye, friend!
```

Generará un fichero `go.mod` que contiene lo siguiente: 

```go
module Package/main

go 1.17
```

## Data Types

Estos son los tipos de datos y tipos compuestos que podemos encontrar en Go.

- Basic Types
	- Integers
		- Signed
			- int
			- int8
			- int16 
			- int32 
			- int64
		- Unsigned
			- uint
			- uint8
			- uint16
			- uint32
			- uint64
			- uintptr
	- Floats
		- float32
		- float64
	- Complex Numbers
		- complex64
		- complex128
	- Byte
	- Rune
	- String
	- Boolean
- Composite Types
	- Collection/Aggregation or Non-Reference Types
		- Arrays
		- Structs
	- Reference Types
		- Slices
		- Maps
		- Channels
		- Pointers
		- Function/Methods
	- Interface
		- Special case of empty Interface

**Números enteros**

```go
//int = Depende del OS (32 o 64 bits)
//int8 = 8bits = -128 a 127
//int16 = 16bits = -2^15 a 2^15-1
//int32 = 32bits = -2^31 a 2^31-1
//int64 = 64bits = -2^63 a 2^63-1
```

**Optimizar memoria cuando sabemos que el dato siempre será positivo**

```go
//uint = Depende del OS (32 o 64 bits)
//uint8 = 8bits = 0 a 127
//uint16 = 16bits = 0 a 2^15-1
//uint32 = 32bits = 0 a 2^31-1
//uint64 = 64bits = 0 a 2^63-1
```

**Números decimales**

```go
// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308
```

**Strings y booleanos**

```go
//string = ""
//bool = true or false
```

**Números complejos**

```go
//Complex64 = Real e Imaginario float32
//Complex128 = Real e Imaginario float64
//Ejemplo : c:=10 + 8i
```

**Recursos**

[Qué es Go o Golang | OpenWebinars](https://openwebinars.net/blog/que-es-go/)

[Go vs Python: Diferencias y puntos fuertes | OpenWebinars](https://openwebinars.net/blog/go-vs-python-diferencias-y-puntos-fuertes/)

[Introduction · GitBook](https://www.pazams.com/Go-for-Javascript-Developers/)

[GoLang Tutorials: Table of Contents](https://golangtutorials.blogspot.com/2011/05/table-of-contents.html)

Revisar este: [All data types in Golang with examples - Welcome To Golang By Example](https://golangbyexample.com/all-data-types-in-golang-with-examples/)

`YouTube Vídeos`

[Golang Tutorial for Beginners | Full Go Course TechWorld With Nana](https://www.youtube.com/watch?v=yyUHQIec83I)

[Go / Golang Crash Course - YouTube](https://www.youtube.com/watch?v=SqrbIlUwR0U)
