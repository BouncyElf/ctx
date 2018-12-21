package ctx

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Context is context of the current http request.
type Context struct {
	Res        http.ResponseWriter
	Req        *http.Request
	urlValue   url.Values
	formValue  url.Values
	StatusCode int
	done       bool
	m          Map
	params     map[string]string

	routerParamsParsed bool
}

// NewContext returns a empty context.
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Res:       w,
		Req:       r,
		urlValue:  nil,
		formValue: nil,
		m:         make(Map),
		params:    make(map[string]string),
	}
}

// Set set a couple of k-v.
func (c *Context) Set(k string, v interface{}) {
	c.m[k] = v
}

// Get get value from the given k, Get only get the value you Set.
func (c *Context) Get(k string) (interface{}, bool) {
	v, ok := c.m[k]
	return v, ok
}

// Request Method

// Params get the router param with the specific k.
func (c *Context) Params(k string) string {
	if !c.routerParamsParsed {
		ctx := c.Req.Context()
		ps := httprouter.ParamsFromContext(ctx)
		for _, p := range ps {
			c.params[p.Key] = p.Value
		}
		c.routerParamsParsed = true
	}
	return c.params[k]
}

// Exists returns if the k exists in query string or form value.
func (c *Context) Exists(k string) bool {
	_ = c.Query("")
	formValue := map[string][]string(c.formValue)
	urlValue := map[string][]string(c.urlValue)
	if len(formValue[k]) != 0 {
		return true
	}
	if len(urlValue[k]) != 0 {
		return true
	}
	return false
}

// Query returns a string value of k. If the specific k not exists, returns "".
func (c *Context) Query(k string) string {
	if c.Method() == "GET" {
		if c.urlValue == nil {
			c.urlValue = c.Req.URL.Query()
		}
		return c.urlValue.Get(k)
	} else {
		if c.formValue == nil || c.urlValue == nil {
			c.Req.ParseForm()
			c.formValue = c.Req.PostForm
			c.urlValue = c.Req.Form
		}
		v := c.formValue.Get(k)
		if v != "" {
			return v
		}
		return c.urlValue.Get(k)
	}
	return ""
}

// QueryInt returns a int value and error if atoi wrong.
func (c *Context) QueryInt(k string) (int, error) {
	sv := c.Query(k)
	return strconv.Atoi(sv)
}

// QueryInt64 returns a int64 value and error if ParseInt wrong.
func (c *Context) QueryInt64(k string) (int64, error) {
	sv := c.Query(k)
	return strconv.ParseInt(sv, 10, 64)
}

// QueryBool returns a bool value and error if ParseBool wrong.
func (c *Context) QueryBool(k string) (bool, error) {
	sv := c.Query(k)
	return strconv.ParseBool(sv)
}

// Cookie returns the http Cookie with the specific name.
func (c *Context) Cookie(name string) (*http.Cookie, error) {
	return c.Req.Cookie(name)
}

// Cookies returns all http cookies in request.
func (c *Context) Cookies() []*http.Cookie {
	return c.Req.Cookies()
}

// SetCookie set a cookie in response.
func (c *Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(c.Res, cookie)
}

// File returns the formfile with the specific name.
func (c *Context) File(name string) (multipart.File, *multipart.FileHeader, error) {
	return c.Req.FormFile(name)
}

// Method returns the method of the current request.
func (c *Context) Method() string {
	return c.Req.Method
}

// URI returns the uri of the current request.
func (c *Context) URI() string {
	return c.Req.URL.RequestURI()
}

// Host returns the host of the current request.
func (c *Context) Host() string {
	return c.Req.Host
}

// Response Method

// Json response the current request with json.
func (c *Context) Json(data interface{}) error {
	if ct := c.Res.Header().Get("Content-Type"); ct == "" {
		c.Res.Header().Set("Content-Type", "application/json")
	}
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.Write(j)
}

// Redirect response the current request and tell the host to request other url.
func (c *Context) Redirect(location string, code ...int) error {
	if len(code) != 0 {
		c.SetStatusCode(code[0])
	} else {
		c.SetStatusCode(303)
	}
	c.done = true
	http.Redirect(c.Res, c.Req, location, c.StatusCode)
	return nil
}

// SetStatusCode set the StatusCode with the specific code.
func (c *Context) SetStatusCode(code int) {
	c.StatusCode = code
	c.Res.WriteHeader(code)
}

// Success response the current request with the specific format of data. The type
// is json, and you can change format by setting ctx.SuccessJson.
func (c *Context) Success(data interface{}) error {
	c.SetStatusCode(200)
	SuccessJson[SuccessKey] = data
	return c.Json(SuccessJson)
}

// Error response the current request with the specific format of data. The type
// is json, and you can change format by setting ctx.ErrorJson.
func (c *Context) Error(code int, msg string) error {
	ErrorJson[ErrorCodeKey] = code
	ErrorJson[ErrorKey] = msg
	return c.Json(ErrorJson)
}

// String response the current request with the specific value s.
func (c *Context) String(s string) error {
	return c.Write([]byte(s))
}

// Write response the current request with data in its body.
func (c *Context) Write(data []byte) error {
	if ct := c.Res.Header().Get("Content-Type"); ct == "" {
		c.Res.Header().Set("Content-Type", "text/plain")
	}
	if !c.done {
		if c.StatusCode == 0 {
			c.SetStatusCode(200)
		}
		c.Res.Write(data)
		c.done = true
	}
	return nil
}
