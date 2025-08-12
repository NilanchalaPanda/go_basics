##  1. Basics of GO

```go
package main

import "fmt"

func main(){
    fmt.Print("Hello");
}
```
- To run this file : **`go run app.go`**
- If you try to run **`go build`**, it will throw error of **ModuleNotFound**


- Go is **statically typed** language. So, types are important to be mentioned > *Int | Float64* 
- All Go value types come with a so-called "null value" > *int = 0 | float64 = 0.0 | string = "" | bool = false*

- `var investmentAmount, years float64 = 1000, 10` > These single line declaration can have multple datatypes as well.

### 1.1) PACKAGES -
![Packages in Go](assets/image.png) 
- `package main` is the **ENTRY point** of the execution for GO.
- Each go file is to be included inside a PACKAGE.
- GO can't have mulitple packages.
- Mulitple files can have the same package.

## 1.2) MODULES -
- Every GO package is under any single Module.
- `go init mod {module_path}` > module_path: example.com/go-path.