> Navega por la tabla de contenido

- [Condicionales](#condicionales)
  - [Switch](#switch)
    - [Primero, segundo o tercero](#primero-segundo-o-tercero)
    - [Conocer el tiempo, mañana o tarde](#conocer-el-tiempo-mañana-o-tarde)
    - [Fin de semana o día de semana](#fin-de-semana-o-día-de-semana)
    - [Tipo de dato de una variable](#tipo-de-dato-de-una-variable)
    - [Contador de vocales](#contador-de-vocales)
- [Keywords](#keywords)
  - [Defer](#defer)
  - [Break](#break)
  - [Continue](#continue)
  - [Defer + Continue + Break](#defer--continue--break)
- [Recursos](#recursos)

## Condicionales

[Anteriormente](https://github.com/jersonmartinez/30-Days-Of-Go/blob/main/Days/day02.md) hemos visto cómo hacer condicionales con `if`, ahora le toca a nuestro amigo `switch`.

### Switch

Te mostraré por medio de ejemplos prácticos, el uso de este tipo de condicionales.

#### Primero, segundo o tercero

Este es muy fácil de entender, se trata de que un valor inicial `i=2`, se verifique y entre en la segunda condición y seguido imprima un _string_.

```go
package main

import (
    "fmt"
)

func main() {
    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("First")
    case 2:
        fmt.Println("Second")
    case 3:
        fmt.Println("Third")
    }
}
```

```bash
go run first_second_or_third.go

#Output
Write 2 as Second
```

#### Conocer el tiempo, mañana o tarde

Este es divertido, importaremos la librería `time` y con la instrucción `time.Now().Hour()` conoceremos la hora actual, al menos la que tiene el sistema en este momento. Si es menor a 12, es de mañana, de lo contrario, es de tarde, después de medio día.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }
}
```

```bash
go run before_or_after_noon.go

#Output
It's after noon
```

#### Fin de semana o día de semana

Con el siguiente ejercicio buscamos validar si el día actual está dentro del fin de semana o día de semana.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday:", time.Now().Weekday())
    }
}
```

```bash
go run weekend_or_weekday.go

#Output
It's a weekday: Friday
```

#### Tipo de dato de una variable

La idea es pasar una variable al `switch` y que este verifique qué tipo de dato es. Este ejercicio es muy interesante, ya que creamos una función anónima, almacenada en  una variable, donde se le pasa una interfaz cómo parámetro.

```go
package main

import (
    "fmt"
)

func main() {
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }

    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}
```

```bash
go run data_type.go

#Output
I'm a bool
I'm an int
Don't know type string
```

#### Contador de vocales

Este me parece que no necesita explicación, pues se trata de contar las vocales que hay en una frase, además de obtener la cantidad de repetición de cada vocal.

```go
package main

import "fmt"

func vowelsCounter(sentence string) (int, int, int, int, int) {
    counter_a := 0
    counter_e := 0
    counter_i := 0
    counter_o := 0
    counter_u := 0
    for _, value := range sentence {
        variable := string(value)
        switch variable {
        case "a":
            counter_a++
        case "e":
            counter_e++
        case "i":
            counter_i++
        case "o":
            counter_o++
        case "u":
            counter_u++
        }
    }
    return counter_a, counter_e, counter_i, counter_o, counter_u
}

func main() {

    sentence := "For example, this is a litter sentence with a lot of vowels"
    a, e, i, o, u := vowelsCounter(sentence)
    fmt.Printf("Phrase '%s' has: \n", sentence)
    fmt.Printf("%d vowel a \n", a)
    fmt.Printf("%d vowel e \n", e)
    fmt.Printf("%d vowel i \n", i)
    fmt.Printf("%d vowel o \n", o)
    fmt.Printf("%d vowel u \n", u)
}
```

```bash
go run vowel_counter.go

#Output
Phrase 'For example, this is a litter sentence with a lot of vowels' has: 
3 vowel a 
7 vowel e
4 vowel i
4 vowel o
0 vowel u
```

## Keywords

### Defer

La idea es que esta palabra reservada (viene de `diferido`), espere a que el programa finalice y ejecuta los `defers`.

```go
package main

import "fmt"

func main() {
    fmt.Println("Start")
    for i := 0; i < 5; i++ {
        defer fmt.Println(i)
    }
    fmt.Println("End")
}
```

```bash
go run defer.go

#Output
Start
End
4  
3  
2  
1  
0
```

Todo lo que está para ejecutarse con `defer`, este lo guarda y lo resuelve hasta el final, cabe destacar que si hay varios `defer` definidos, tal y como sucede en este ejemplo, este se almacena en una cola `LIFO` (_Last in First Out_ / Último en entrar, primero en salir).

```go
package main

import "fmt"

func multiplication(a, b int) int {
    return (a * b)
}

func show() {
    fmt.Println("Hello, my name is Antonio!")
}

func main() {

    fmt.Println("Result is:", multiplication(10, 10))

    defer fmt.Println("Result is:", multiplication(15, 15))

    show()
}
```

```bash
go run defer_with_functions.go

#Output
Result is: 100
Hello, my name is Antonio!
Result is: 225
```

### Break

Este rompe con el ciclo o condición.

```go
package main

import "fmt"

func main() {
    var i uint8 = 0
    for i < 10 {
        fmt.Println(i)
        i++

        if i == 8 {
            fmt.Println("Break!")
            break
        }
    }
}
```

```bash
go run break.go

#Output
0
1
2
3
4
5
6
7
Break!
```

### Continue

El continue rompe con lo que se encuentra dentro del ciclo y continua con la siguiente iteración, saltando lo que haya dentro del ciclo.

```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        if i == 2 {
            continue
        }
        fmt.Print(i, "\n")
    }
}
```

```bash
go run continue.go

#Output
0
1
3
4
5
6
7
8
9
```

### Defer + Continue + Break

```go
package main

import "fmt"

func main() {
    defer fmt.Println("Hola")
    fmt.Println("Mundo")

    for i := 0; i < 10; i++ {

        if i == 2 {
            fmt.Println("es 2")
            continue
        }

        fmt.Println(i)

        if i == 5 {
            break
        }
    }
}
```

```bash
go run defer_continue_break.go

#Output
Mundo
0
1
es 2
3
4
5
Hola
```

Estos han sido algunos ejemplos del modo de uso de las palabras claves `defer`, `continue` y `break` en Go, además de profundizar en la forma de hacer validaciones con la instrucción `switch`.

## Recursos

[Inicia tu carrera de programador en Go](https://www.crashell.com/estudio/inicia_tu_carrera_de_programador_en_go)

[Go by Example: Switch](https://gobyexample.com/switch)

[Defer Keyword in Golang - GeeksforGeeks](https://www.geeksforgeeks.org/defer-keyword-in-golang/)
