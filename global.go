package ctx

// SuccessCB is the c.Success() callback.
// DO NOT USE DEFAULT, MAKE IT YOURS.
var SuccessCB = func(*Context, interface{}) error { return nil }

// ErrorCB is the c.Error() callback.
// DO NOT USE DEFAULT, MAKE IT YOURS.
var ErrorCB = func(*Context, int, interface{}) error { return nil }

// ErrorHandler is the error handler when error occured.
// DO NOT USE DEFAULT, MAKE IT YOURS.
var ErrorHandler = func(c *Context, err error) {
	if c.StatusCode == 0 {
		c.StatusCode = 500
	}
	c.Error(c.StatusCode, err)
}
