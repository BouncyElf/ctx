// CTX is a web framework which simply use httprouter as its router,
// and offer an easier method (or just alias) of the 'net/http'.
// CTX itself dose not integrate websocket, but if you need this, please go
// check https://github.com/gorilla/websocket or a easier one like
// https://github.com/olahol/melody.
package ctx

import (
	"errors"
	"fmt"
	"net/http"
)

// Map is the alias of map[string]interface{}
type Map map[string]interface{}

// ErrNotFound is the NotFound error.
var ErrNotFound = errors.New("404 Not Found")

// ErrMethodNotAllow is the MethodNotAllowed error.
var ErrMethodNotAllow = errors.New("405 Method Not Allow")

// SuccessCB is the c.Success() callback.
// NOTE: DO NOT USE DEFAULT, MAKE IT YOURS.
var SuccessCB = func(*Context, interface{}) error { return nil }

// ErrorCB is the c.Error() callback.
// NOTE: DO NOT USE DEFAULT, MAKE IT YOURS.
var ErrorCB = func(c *Context, code int, msg interface{}) error {
	innerError := ""
	switch msg.(type) {
	case string:
		innerError = msg.(string)
	case error:
		innerError = msg.(error).Error()
	default:
		innerError = "internal server error, unsupported error message type"
	}
	http.Error(c.Res, innerError, code)
	return nil
}

// ErrorHandler is the centralized error handler.
// NOTE: DO NOT USE DEFAULT, MAKE IT YOURS.
var ErrorHandler = func(c *Context, err error) {
	if err == ErrNotFound {
		c.SetStatusCode(http.StatusNotFound)
		c.Error(http.StatusNotFound, err)
	} else if err == ErrMethodNotAllow {
		c.SetStatusCode(http.StatusMethodNotAllowed)
		c.Error(http.StatusMethodNotAllowed, err)
	} else {
		c.SetStatusCode(http.StatusInternalServerError)
		c.Error(c.StatusCode, err)
	}
}

// PanicHandler is the centralized panic handler.
// NOTE: DO NOT USE DEFAULT, MAKE IT YOURS.
var PanicHandler = func(c *Context, msg interface{}) {
	if c.StatusCode == 0 {
		c.StatusCode = 500
	}
	c.Error(c.StatusCode, msg)
}

// e returns a ctx error.
func e(msg string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s %s: %v", "[CTX]", msg, err)
}
