# CTX
A Go web framework.

## Install
```bash
$ go get -u -v github.com/BouncyElf/ctx
```

## Example
```Go
package main

import "github.com/BouncyElf/ctx"

func main() {
	r := ctx.New()
	r.GET("/", func(c *ctx.Context) error {
		return c.String("hello, world")
	})
	r.Run()
}
```

## Doc
[Doc Here](https://godoc.org/github.com/BouncyElf/ctx)
