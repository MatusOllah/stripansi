# stripansi

[![Go Reference](https://pkg.go.dev/badge/github.com/MatusOllah/stripansi.svg)](https://pkg.go.dev/github.com/MatusOllah/stripansi) [![Go Report Card](https://goreportcard.com/badge/github.com/MatusOllah/stripansi)](https://goreportcard.com/report/github.com/MatusOllah/stripansi)

**stripansi** is a Go package for removing [ANSI escape sequences](https://en.wikipedia.org/wiki/ANSI_escape_code).

## Basic Usage

```go
package main

import (
    "fmt"
    "os"

    "github.com/MatusOllah/stripansi"
)

func main() {
    s := "\x1b[38;5;140mhello\x1b[0m world"

    fmt.Println(stripansi.String(s))                // "hello world"
    fmt.Println(string(stripansi.Bytes([]byte(s)))) // "hello world"
    stripansi.NewWriter(os.Stdout).Write([]byte(s)) // "hello world"
}
```
