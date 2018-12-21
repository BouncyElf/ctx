package ctx

var SuccessJson = Map{}
var SuccessKey = ""
var ErrorJson = Map{}
var ErrorKey = ""
var ErrorCodeKey = ""

var ErrorHandler = func(c *Context, err error) {
	if c.StatusCode == 0 {
		c.StatusCode = 500
	}
	c.Error(c.StatusCode, err.Error())
}
