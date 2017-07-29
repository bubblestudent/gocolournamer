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

## Inpsiration

Idea and colour list from http://chir.ag/projects/ntc/ javascript implementation.

## Lisence

Distributed free for any use under an MIT licence.