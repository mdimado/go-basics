# The Go language 🦦

This repository contains all the fundamental concepts of Go, providing a solid foundation to start working on Go projects

## Go installation and setup:

### Installation

The first step is to install Go on yous system
Follow the [official installation guide](https://go.dev/doc/install) 

### Setting up the Go workspace

First, create a working directory on your system:

```sh
mkdir my_og_project
cd my_og_project
```

Go uses modules for managing dependencies, the `go mod` command is used to initiliaze the Go module, careting a `go.mod` file

```sh
go mod init my_og_project
```

or if you want it to be publicly available

```sh
go mod init github.com/your-username/my_og_project
```

so that, if someone wants to use your project as a library, they can import it as:

```sh
import "github.com/your-username/my_og_project"
```

## [Hello World](./hello_world.go)

Now that a go.mod file has been initialized and the project has been set up successfully, it's time to write your first program

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}

```

### A few important points:
- every go program starts with a package
    **why ?**
    - Go is designed around packages for modularity
    **why `package main`?**
    - for an executable, stand-alone program, use `package main`. Go looks for `package main` and converts it into binary
- the syantax for importing external packages in Go is `import "package-name"` for a single package and for multiple packages, 
```go
import (
    "list",
    "of",
    "packages"
)
```
- to import package with alias
```go
import (
    f "fmt"
)
```
- Go gives an `imported but unused` error if a package is unused after importing
To prevent this, use `Blank import _`
```go
import _ "database/sql"
```
- External packages can be installed using `go get`
- The `fmt` package is an inbuilt package in Go's standard library. It helps in:
    - Printing to the cosole
    - Formatting strings
    - Scanning user input
- The `main` function is the entry point of every executable program in Go
    - `func main()` is where executon begins
    - Go will throw an error if the `func main()` is not found
    - Global variables, constants, and types can be declared outside functions
    - `init()` can be used for setup before `main()`


## Data Types in Go

### Integer Types:
These store whole numbers
- `int` platform-dependent
    - stores 32 bit signed integers on a 32-bit system
    - stores 64 bit signed integers on a 64-bit system
- `intN`
    - stores N bit signed integers
    - example, `int8` stores 8 bit signed integers and `int32` stores 32 bit signed integers

### Unsigned integers
Stores only positive whole numbers
- `uint, uint8, uint16, uint32, uint64`

### Floating point types
Stores decimal numbers
- `float32, float64`

### Boolean
Represents `true` or `false`
- `bool`


## Declaration of varibles

### **Key Differences**
| Feature          | `var` Syntax        | `:=` Syntax |
|-----------------|--------------------|-------------|
| **Type required?** | Optional (can be specified) | No (inferred automatically) |
| **Inside functions?** | ✅ Yes | ✅ Yes |
| **Outside functions?** | ✅ Yes | ❌ No (only `var` works globally) |
| **Allows multiple declarations?** | ✅ Yes | ❌ No |

---

<span style="color:red">Go does not allow variable redecclaration in the same scope</span>
<span style="color:red">Go is strict about not using any declared variable. It throws an error</span>.
