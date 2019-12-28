# Random String Generator
This package generate random string.

## Installation
```sh
$ go get github.com/dkuye/random
```

## Usage
```go
package main

import (
    "fmt"
    "github.com/dkuye/random"
)

func main(){
    str := random.String{Lower: true, Upper: true}.Gen(8)
    fmt.Print(str)
}
```

## Available Option
-   Lower
-  	Upper
-  	Number
-  	Special
