package ctx

import (
	"errors"
	"net/http"
)

// ErrNotFound is the NotFound error.
// Use this to distinguish error.
var ErrNotFound = errors.New("404 Not Found")

// ErrMethodNotAllow is the MethodNotAllowed error.
// Use this to distinguish error.
var ErrMethodNotAllow = errors.New("405 Method Not Allow")

// SuccessCB is the c.Success() callback.
// DO NOT USE DEFAULT, MAKE IT YOURS.
var SuccessCB = func(*Context, interface{}) error { return nil }

// ErrorCB is the c.Error() callback.
// DO NOT USE DEFAULT, MAKE IT YOURS.
var ErrorCB = func(*Context, int, interface{}) error { return nil }

// ErrorHandler is the error handler when error occured.
// DO NOT USE DEFAULT, MAKE IT YOURS.
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

// PanicHandler is the panic handler when panic occurred.
// DO NOT USE DEFAULT, MAKE IT YOURS.
var PanicHandler = func(c *Context, msg interface{}) {
	if c.StatusCode == 0 {
		c.StatusCode = 500
	}
	c.Error(c.StatusCode, msg)
}
