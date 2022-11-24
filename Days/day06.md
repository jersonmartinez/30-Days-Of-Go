> Navega por la tabla de contenido

- [Maps](#maps)
  - [Introducción](#introducción)
    - [Qué son los mapas en Go](#qué-son-los-mapas-en-go)
    - [Características](#características)
    - [Tipos de clave permitidos](#tipos-de-clave-permitidos)
    - [Tipos de valores permitidos](#tipos-de-valores-permitidos)
  - [Ejercicios](#ejercicios)
    - [Declarando un mapa](#declarando-un-mapa)
    - [Inicializando un mapa](#inicializando-un-mapa)
    - [Inicializando un mapa mediante la funcion make()](#inicializando-un-mapa-mediante-la-funcion-make)
    - [Acceso a elementos del mapa](#acceso-a-elementos-del-mapa)
  - [Recursos](#recursos)

# Maps

## Introducción

### Qué son los mapas en Go

Una de las estructuras de datos más útiles en informática es la tabla hash. Existen muchas implementaciones de tablas hash con propiedades variables, pero en general ofrecen búsquedas, adiciones y eliminaciones rápidas. Go proporciona un tipo de mapa integrado que implementa una tabla hash.

**Tabla hash**

Es una estructura de datos no lineal cuyo propósito final se centra en llevar a cabo las acciones básicas (inserción, eliminación y búsqueda de elementos) en el menor tiempo posible, mejorando las cotas de rendimiento respecto a un gran número de estructuras.


### Características

- Los mapas se utilizan para almacenar valores de datos en pares `key`:`value`.
- Cada elemento de un mapa es un par de `key`:`value`.
- Un mapa es una colección desordenada y modificable que no permite duplicados.
- La longitud de un mapa es el número de sus elementos. Este dato lo puedes obtener también utilizando la función `len()`.
- El valor predeterminado de un mapa es `nil`.
- Los mapas contienen referenias a una tabla hash subyacente.
- Go tiene varias formas de crear mapas. 

### Tipos de clave permitidos

La clave de asignación puede ser de cualquier tipo de datos para el que se haya definido el operador de igualdad.

- Booleanos (_Booleans_)
- Números (_Numbers_)
- Cadenas (_Strings_)
- Matrices (_Arrays_)
- Punteros (_Pointers_)
- Estructuras (_Structs_)
- Interfaces (siempre que el tipo dinámico admita la igualdad)

Los tipos de clave no válidos son:

- Rebanadas (_Slices_)
- Mapas (_Maps_)
- Funciones (_Functions_)

Estos tipos no son válidos porque el operador de igualdad () no está definido para ellos.==

### Tipos de valores permitidos

Los valores del mapa pueden ser de cualquier tipo.

Un tipo de mapa Go tiene el siguiente aspecto: 

```go
map[KeyType]ValueType
```

## Ejercicios

### Declarando un mapa

```go
package main

import (
	"fmt"
)

func main() {
	var mapa map[string]string

	fmt.Println(mapa)
}
```

```bash
go run declaration.go

#Output
map[]
```

### Inicializando un mapa

Se inicializan dos mapas, el primero que está constituído como _string_ tanto la clave como el valor. El segundo está definido para que la clave sea de _string_ y los valores, numéricos enteros (_int_).

```go
package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{
		"Username": "jersonmartinez",
		"Name":     "Jerson Antonio",
		"Lastname": "Martínez Moreno",
	}

	b := map[string]int{
		"Day":          6,
		"Month":        4,
		"Year":         1996,
		"Magic number": 9,
	}

	fmt.Printf("a = \t%v\n", a)
	fmt.Printf("b = \t%v\n", b)

}
```

```bash
go run initialization.go

# Output
a =     map[Lastname:Martínez Moreno Name:Jerson Antonio Username:jersonmartinez]
b =     map[Day:6 Magic number:9 Month:4 Year:1996]
```

### Inicializando un mapa mediante la funcion make()

```go
package main

import (
	"fmt"
)

func main() {

	var a = make(map[string]string)

	a["Username"] = "jersonmartinez"
	a["Name"] = "Jerson Antonio"
	a["Lastname"] = "Martínez Moreno"

	b := make(map[string]int)

	b["Day"] = 6
	b["Month"] = 4
	b["Year"] = 1996
	b["Magic number"] = 9

	fmt.Printf("a = \t%v\n", a)
	fmt.Printf("b = \t%v\n", b)

}
```

```bash
go run initialization_with_make.go

#Output
a =     map[Lastname:Martínez Moreno Name:Jerson Antonio Username:jersonmartinez]
b =     map[Day:6 Magic number:9 Month:4 Year:1996]
```

### Acceso a elementos del mapa

```go
package main
import ("fmt")

func main() {
  var a = make(map[string]string)
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"

  fmt.Printf(a["brand"])
}
```

## Recursos

[🔥 Mapas (Maps)](https://apuntes.de/golang-estructuras-de-datos-y-algoritmos/mapas/#gsc.tab=0)
[¿Qué es una tabla hash](https://www.ecured.cu/Tabla_hash)
[Go Maps](https://www.w3schools.com/go/go_maps.php)