package ctx

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"strings"
)

type Map map[string]interface{}

type Handler func(*Context)

func (h Handler) NewHttpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := NewContext(w, r)
		h(ctx)
	}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handle := h.NewHttpHandler()
	handle(w, r)
}

type Context struct {
	Res       http.ResponseWriter
	Req       *http.Request
	urlValue  url.Values
	formValue url.Values
	done      bool
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Res:       w,
		Req:       r,
		urlValue:  nil,
		formValue: nil,
	}
}

// Request Method

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

func (c *Context) QueryInt(k string) (int, error) {
	sv := c.Query(k)
	return strconv.Atoi(sv)
}

func (c *Context) QueryInt64(k string) (int64, error) {
	sv := c.Query(k)
	return strconv.ParseInt(sv, 10, 64)
}

func (c *Context) QueryBool(k string) (bool, error) {
	sv := c.Query(k)
	return strconv.ParseBool(sv)
}

func (c *Context) Cookie(name string) (*http.Cookie, error) {
	return c.Req.Cookie(name)
}

func (c *Context) Cookies() []*http.Cookie {
	return c.Req.Cookies()
}

func (c *Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(c.Res, cookie)
}

func (c *Context) File(name string) (multipart.File, *multipart.FileHeader, error) {
	return c.Req.FormFile(name)
}

func (c *Context) Method() string {
	return c.Req.Method
}

func (c *Context) URI() string {
	return c.Req.URL.RequestURI()
}

func (c *Context) Host() string {
	return c.Req.Host
}

func (c *Context) IsAjaxReq() bool {
	s := c.Req.Header.Get("HTTP_X_REQUESTED_WITH")
	s = strings.ToLower(s)
	return s == "xmlhttprequest"
}

func (c *Context) AcceptJson() bool {
	accept := c.Req.Header.Get("Accept")
	return strings.Contains(accept, "application/json")
}

// Response Method

func (c *Context) Json(data interface{}) {
	if ct := c.Res.Header().Get("Content-Type"); ct == "" {
		c.Res.Header().Set("Content-Type", "application/json")
	}
	j, err := json.Marshal(data)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		c.Write([]byte(fmt.Sprintf("%s:line %d, json marshal error:%v", f, l, err)))
		return
	}
	c.Write(j)
}

func (c *Context) Redirect(location string, code ...int) {
	// TODO
}

func (c *Context) Success(data interface{}) {
	c.Json(map[string]interface{}{
		"errno":  0,
		"errmsg": "",
		"data":   data,
	})
}

func (c *Context) Error(code int, msg string) {
	c.Json(map[string]interface{}{
		"errno":  code,
		"errmsg": msg,
		"data":   "",
	})
}

func (c *Context) String(s string) {
	c.Write([]byte(s))
}

func (c *Context) Write(data []byte) {
	if ct := c.Res.Header().Get("Content-Type"); ct == "" {
		c.Res.Header().Set("Content-Type", "text/plain")
	}
	if !c.done {
		c.Res.Write(data)
		c.done = true
	}
}
