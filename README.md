Go Colour Namer
=====================

## Using this project

Install the package

```bash
$ go get github.com/bubblestudent/gocolournamer
```

Then add the import to the top of whichever file you want to use the package in

```go
import "github.com/bubblestudent/gocolournamer"
```

All of the packages functionality is exposed in one method. ToNearestColour is called with a valid hex colour and returns the named struct and an error.

```go
named, err := gocolournamer.ToNearestColour(hexstring)
```

Additional documentation available on godoc: http://godoc.org/github.com/bubblestudent/gocolournamer

## Example program

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/bubblestudent/gocolournamer"
)

func main() {
	hex := "#905E26" //hash is optional
	named, err := gocolournamer.ToNearestColour(hex)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	out, _ := json.Marshal(named)
	fmt.Printf("%s", string(out))
}
```

The above program outputs

```bash
{"hex":"905E26","colour":"Afghan Tan","hue":"Yellow","huehex":"FFFF00"}
```

## Inpsiration

Idea and colour list from http://chir.ag/projects/ntc/ javascript implementation.

## Lisence

Distributed free for any use under an MIT licence.
