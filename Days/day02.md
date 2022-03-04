> Navega por la tabla de contenido

- [Forma de ver documentación](#forma-de-ver-documentación)
  - [Go doc](#go-doc)
- [Loops](#loops)
- [Condicionales](#condicionales)
  - [Operadores lógicos y de comparación](#operadores-lógicos-y-de-comparación)
    - [Operadores de comparación](#operadores-de-comparación)
    - [Operadores lógicos](#operadores-lógicos)
      - [Operador AND:](#operador-and)
      - [Operador OR:](#operador-or)
      - [Operador NOT:](#operador-not)
  - [Primeros pasos](#primeros-pasos)
  - [Condicionales usando funciones](#condicionales-usando-funciones)
- [Recursos](#recursos)

## Forma de ver documentación

### Go doc

Este es un apartado donde muestro enlaces a fuentes oficinales de Go para documentarse sobre el lenguaje. [Go Docs](https://pkg.go.dev/).

![Go Docs](C:\Users\Root\OneDrive\Documentos\Projects\Repositories\30-Days-Of-Go\Days\02_Day\src\2022-03-03-21-58-48-image.png)

Desde ahí es posible documentarse sobre cualquier paquete de Go. Por ejemplo, buscando sobre el paquete `fmt`.

![](C:\Users\Root\OneDrive\Documentos\Projects\Repositories\30-Days-Of-Go\Days\02_Day\src\2022-03-03-22-03-49-image.png)

## Loops

Dicen las malas lenguas que Go solo tiene for para aplicar ciclos, no como en otros lenguajes que existe en while y el do while, foreach y otros.

El ciclo más sencillo sería el siguiente:

```go
package main

func main() {
    for i := 0; i < 10; i++ {
        println(i)
    }
}
```

```bash
go run *.go

#Output
0
1
2
3
4
5
6
7
8
9
```

Este que se acaba de mostrar el es `for` de toda la vida.

**For / While**

```go
package main

func main() {
    i := 0
    for i < 10 {
        println(i)
        i++
    }
}
```

**For / Forever**

```go
package main

import "fmt"

func main() {
    counter := 0
    for {
        fmt.Println("Forever: ", counter)
        counter++
    }
}
```

**For combinado**

```go
package main

import "fmt"

func main() {
	/*
		Imprime los números del 1 al 3
	*/
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	/*
		Imprime los números del 7 al 9
	*/
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	/*
		Realiza un ciclo infinito, sin embargo, lo detiene a la primera.
	*/
	for {
		fmt.Println("loop")
		break
	}

	/*
		Imprime solo números impares, ya que omite los números pares por medio del módulo.
	*/
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
```

```bash
go run for_combinado.go

#Output
2
3
7
8
9
loop
1
3
5
```



## Condicionales

### Operadores lógicos y de comparación

Son operadores que nos permiten hacer una comparación de condiciones y
 en caso de cumplirse como sino ejecutarán un código determinado. Si se 
cumple es `VERDADERO/TRUE` y si no se cumple son `FALSO/FALSE`.

Empecemos con los operadores de comparación:

#### Operadores de comparación

Son aquellos que retornan TRUE o FALSE en caso de cumplirse o no una expresión. Son los siguientes:

- *valor1 == valor2*: Retorna TRUE si valor1 y valor2 son exactamente iguales.
- *valor1 != valor2*: Retorna TRUE si valor1 es diferente de valor2.
- *valor1 < valor2*: Retorna TRUE si valor1 es menor que valor2
- *valor1 > valor2*: Retorna TRUE si valor1 es mayor que valor2
- *valor1 >= valor2*: Retorna TRUE si valor1 es igual o mayor que valor2
- *valor1 <= valor2*: Retorna TRUE si valor1 es menor o igual que valor2.

#### Operadores lógicos

Son aquellos que retorna TRUE o FALSE si cumplen o no una condición.

##### Operador AND:

Este operador indica que todas las condiciones declaradas deben 
cumplirse para poderse marcar como TRUE. En Go, se utiliza este símbolo `&&`.

Ejemplo1: `1>0 && 2 >0` Esto retornará TRUE porque tanto la primera como la segunda condición son verdaderas.

Ejemplo2: `2<0 && 1 > 0` Esto retornará FALSE porque una de las condiciones no es verdadera.

##### Operador OR:

Este operador indica que al menos una de las condiciones debe 
cumplirse para marcarse como TRUE. En Go, se representa con el símbolo `||`.

Ejemplo: `2<0 || 1 > 0` Esto retornará TRUE porque la segunda condición se cumple, a pesar que la primera no.

##### Operador NOT:

Este operador retornará el opuesto al boleano que está dentro de la variable. Ejemplo:

```Go
myBool :=  true
fmt.Println(!myBool) // Esto retornará false
```

### Primeros pasos

```go
package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 es par")
	} else {
		fmt.Println("7 es impar")
	}

	if 8%4 == 0 {
		fmt.Println("8 es divisible de 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "es negativa")
	} else if num < 10 {
		fmt.Println(num, "tiene 1 dígito")
	} else {
		fmt.Println(num, "tiene multiples dígitos")
	}
}
```

```bash
go run conditionals.go 

#Output
7 es impar
8 es divisible de 4
9 tiene 1 dígito
```

### Condicionales usando funciones

```go
package main

import "fmt"

func main() {
	if isPair(6) {
		fmt.Println("Número par")
	} else {
		fmt.Println("Número impar")
	}
	if isValidUser("Antonio5", "Password") {
		fmt.Println("Credentiales válidas")
	} else {
		fmt.Println("Credentiales inválidas")
	}
}

func isPair(num int) bool {
	return num%2 == 0
}

func isValidUser(userName, pass string) bool {
	return userName == "Antonio" && pass == "Password"
}
```

```bash
go run conditionals_with_functions.go 

#Output
Número par
Credentiales inválidas
```

Continuará...

## Recursos

[fmt package - fmt - pkg.go.dev](https://pkg.go.dev/fmt#example-Println)

[Go cheatsheet](https://devhints.io/go)

https://awesome-go.com/

https://gobyexample.com/

`YouTube Vídeos`

[justforfunc: Programming in Go - YouTube](https://www.youtube.com/c/JustForFunc/)
