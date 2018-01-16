# Beego

## Context

https://github.com/astaxie/beego/blob/master/context/context.go

A wrapper for http.ResponseWriter, but didn't make use of that for logging because size is not tracked https://github.com/astaxie/beego/blob/master/router.go#L884

````go
type Context struct {
	Input          *BeegoInput
	Output         *BeegoOutput
	Request        *http.Request
	ResponseWriter *Response
	_xsrfToken     string
}

type Response struct {
	http.ResponseWriter
	Started bool
	Status  int
}
````