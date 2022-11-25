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
		- [Inicializando un mapa mediante la función _make()_](#inicializando-un-mapa-mediante-la-función-make)
		- [Acceso a elementos del mapa](#acceso-a-elementos-del-mapa)
		- [Cambiar valores dentro de un mapa](#cambiar-valores-dentro-de-un-mapa)
		- [Agregar elementos al mapa](#agregar-elementos-al-mapa)
		- [Eliminar elemento en el mapa](#eliminar-elemento-en-el-mapa)
		- [Recorriendo un mapa por medio de bucles](#recorriendo-un-mapa-por-medio-de-bucles)
		- [Validar elementos específicos en el mapa](#validar-elementos-específicos-en-el-mapa)
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
- Los mapas contienen referencias a una tabla hash subyacente.
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

Estos tipos no son válidos porque el operador de igualdad () no está definido para ellos.

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

Se inicializan dos mapas, el primero que está constituido como _string_ tanto la clave como el valor. El segundo está definido para que la clave sea de _string_ y los valores, numéricos enteros (_int_).

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

### Inicializando un mapa mediante la función _make()_

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

# Output
a =     map[Lastname:Martínez Moreno Name:Jerson Antonio Username:jersonmartinez]
b =     map[Day:6 Magic number:9 Month:4 Year:1996]
```

### Acceso a elementos del mapa

Una vez creado al mapa, la forma de acceder es la misma que la de un arreglo; por medio del índice en la variable que contiene el mapa, ya sea que este índice sea compuesto de caracteres o numérica (depende de cómo se define el mapa). Observa el siguiente ejemplo donde se almacenan 3 datos y se imprimen dichos datos.

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

	fmt.Printf("Your Username is: " + a["Username"] + "\n")
	fmt.Printf("Your Name is: " + a["Name"] + "\n")
	fmt.Printf("Your Lastname is: " + a["Lastname"])
}
```

```bash
go run get_data.go

# Output
Your Username is: jersonmartinez
Your Name is: Jerson Antonio
Your Lastname is: Martínez Moreno
```

### Cambiar valores dentro de un mapa

```go
package main

import (
	"fmt"
)

func main() {
	// Create a map
	var a = map[string]string{
		"Username": "jersonmartinez",
		"Name":     "Jerson Antonio",
		"Lastname": "Martínez Moreno",
	}

	fmt.Printf("=> Mapa inicial\n")
	fmt.Printf("Your Username is: " + a["Username"] + "\n")
	fmt.Printf("Your Name is: " + a["Name"] + "\n")
	fmt.Printf("Your Lastname is: " + a["Lastname"] + "\n\n")

	// Change values of the map
	a["Username"] = "CoreStack"
	a["Name"] = "Antonio"
	a["Lastname"] = "Moreno"

	fmt.Printf("=> Mapa modificado\n")
	fmt.Printf("Your Username is: " + a["Username"] + "\n")
	fmt.Printf("Your Name is: " + a["Name"] + "\n")
	fmt.Printf("Your Lastname is: " + a["Lastname"])
}

```

```bash
go run change_values.go

# Output
=> Mapa inicial
Your Username is: jersonmartinez
Your Name is: Jerson Antonio
Your Lastname is: Martínez Moreno

=> Mapa modificado
Your Username is: CoreStack
Your Name is: Antonio
Your Lastname is: Moreno
```

### Agregar elementos al mapa

```go
package main

import (
	"fmt"
)

func main() {
	// Create a map
	var User = map[string]string{
		"Username": "jersonmartinez",
		"Name":     "Jerson Antonio",
		"Lastname": "Martínez Moreno",
	}

	// Add element
	User["Age"] = "26"

	// Add other element
	User["Career"] = "DevOps Engineer"

	fmt.Printf("Your Username is: " + User["Username"] + "\n")
	fmt.Printf("Your Name is: " + User["Name"] + "\n")
	fmt.Printf("Your Lastname is: " + User["Lastname"] + "\n")

	// New elements
	fmt.Printf("Your Age is: " + User["Age"] + "\n")
	fmt.Printf("Your Career is: " + User["Career"] + "\n\n")

}
```

```bash
go run add_element.go

# Output
Your Username is: jersonmartinez
Your Name is: Jerson Antonio
Your Lastname is: Martínez Moreno
Your Age is: 26
Your Career is: DevOps Engineer
```

### Eliminar elemento en el mapa

Go tiene una función integrada (`delete()`), que elimina un elemento respecto a la clave especificada en un mapa. Si la clave no está disponible en el mapa, la eliminación no es operativa; en caso de que exista la clave, este la elimina del mapa.

```go
package main

import (
	"fmt"
)

func main() {
	// Create a map
	var User = map[string]string{
		"Username": "jersonmartinez",
		"Name":     "Jerson Antonio",
		"Lastname": "Martínez Moreno",
	}

	// Add element
	User["Age"] = "26"

	// Add other element
	User["Career"] = "DevOps Engineer"

	fmt.Printf("Your Username is: " + User["Username"] + "\n")
	fmt.Printf("Your Name is: " + User["Name"] + "\n")
	fmt.Printf("Your Lastname is: " + User["Lastname"] + "\n")

	// Delete element "Age" key of the map "User".
	delete(User, "Age")

	// New elements
	fmt.Printf("Your Age is: " + User["Age"] + "\n")
	fmt.Printf("Your Career is: " + User["Career"] + "\n\n")

}
```

```bash
go run remove_element.go

# Output
Your Username is: jersonmartinez
Your Name is: Jerson Antonio     
Your Lastname is: Martínez Moreno
Your Age is: 
Your Career is: DevOps Engineer
```

### Recorriendo un mapa por medio de bucles

Se recorre un mapa por medio de iteraciones en sus elementos. Usando la estructura repetitiva `for`.

```go
package main

import (
	"fmt"
)

func main() {
	// Create a map
	var User = map[string]string{
		"Username": "jersonmartinez",
		"Name":     "Jerson Antonio",
		"Lastname": "Martínez Moreno",
	}

	// Looping through the map
	for key, value := range User {
		fmt.Printf("The %s is: %s \n", key, value)
	}

}
```

```bash
go run looping_through.go

# Output
The Username is: jersonmartinez 
The Name is: Jerson Antonio 
The Lastname is: Martínez Moreno
```

### Validar elementos específicos en el mapa

Es posible realizar un conjunto de validaciones en los elementos que integran el mapa. Si solo desea comprobar la existencia de una determinada clave, puede usar el identificador en blanco (`_`) en lugar de la variable que almacena el valor.

```go
package main

import "fmt"

func main() {
	// Create a map
	var User = map[string]string{
		"Username": "jersonmartinez",
		"Name":     "Jerson Antonio",
		"Lastname": "Martínez Moreno",
		"Age":      "",
	}

	// Check if exists key and its value
	k1, v1 := User["Username"]

	// Check if non-exists key and its value
	k2, v2 := User["Phone"]

	// Check if exists key and its value
	k3, v3 := User["Age"]

	// Only check if exists key and not its value
	_, v4 := User["Name"]

	fmt.Println(k1, v1)
	fmt.Println(k2, v2)
	fmt.Println(k3, v3)
	fmt.Println(v4)

}
```

```bash
go run check_specific_elemetns.go

# Output
jersonmartinez true
 false
 true 
true
```

Se ha podido introducir Mapas en Go, pasando por su concepto, características, tipos de clave permitidos y de valores permitidos. Además, por medio de ejercicios, aprendes a declarar un mapa, inicializar, acceso a elementos, cambiar valores, agregar elementos, eliminar y recorrer los elementos, asimismo, también verificar dichos elementos.

## Recursos

[🔥 Mapas (Maps)](https://apuntes.de/golang-estructuras-de-datos-y-algoritmos/mapas/#gsc.tab=0)

[¿Qué es una tabla hash](https://www.ecured.cu/Tabla_hash)

[Go Maps](https://www.w3schools.com/go/go_maps.php)

[Progamiz - Maps in Go](https://www.programiz.com/golang/map)