package ctx

// SuccessJson is the success json format.
var SuccessJson = Map{}

// SuccessKey is the key of the success data.
var SuccessKey = ""

// ErrorJson is the error json format.
var ErrorJson = Map{}

// ErrorKey is the key of the error data.
var ErrorKey = ""

// ErrorCodeKey is the key of the error code.
var ErrorCodeKey = ""

// ErrorHandler is the error handler when error occured.
// DO NOT USE DEFAULT, Make it your owns.
var ErrorHandler = func(c *Context, err error) {
	if c.StatusCode == 0 {
		c.StatusCode = 500
	}
	c.Error(c.StatusCode, err.Error())
}
