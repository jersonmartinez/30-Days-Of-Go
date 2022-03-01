# Introducción

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





**Fuentes**

[Qué es Go o Golang | OpenWebinars](https://openwebinars.net/blog/que-es-go/)

[Go vs Python: Diferencias y puntos fuertes | OpenWebinars](https://openwebinars.net/blog/go-vs-python-diferencias-y-puntos-fuertes/)

[Introduction · GitBook](https://www.pazams.com/Go-for-Javascript-Developers/)

[GoLang Tutorials: Table of Contents](https://golangtutorials.blogspot.com/2011/05/table-of-contents.html)

YouTube Vídeos

[Golang Tutorial for Beginners | Full Go Course TechWorld With Nana](https://www.youtube.com/watch?v=yyUHQIec83I)

[Go / Golang Crash Course - YouTube](https://www.youtube.com/watch?v=SqrbIlUwR0U)
