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

[## Hello World](./hello_world.go)

Now that a go.mod file has been initialized and the project has been set up successfully, it's time to write your first program

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}

```

a few important points:
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
