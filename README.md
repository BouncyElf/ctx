# CTX
🏝A Go web framework.

## Feature
- Singleton.
- Centralized error handling.
- Centralized panic recover & handling.
- Gracefully shutdown.
- Websocket support.

## Install
```bash
$ go get -u -v github.com/BouncyElf/ctx
```

## Example
```Go
package main

import "github.com/BouncyElf/ctx"

func main() {
	ctx.GET("/", func(c *ctx.Context) error {
		return c.String("hello, world")
	})
	ctx.Run()
}
```

## Doc
[Doc Here](https://godoc.org/github.com/BouncyElf/ctx)
