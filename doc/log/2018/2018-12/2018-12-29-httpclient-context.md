# 2018-12-29 HTTP Client context

This doc describes what is needed for context and client

Both

- [x] header
- retry
- encoding

Context only

- [x] query parameter
- is a stream request

Client only

- [x] bath path `/api/v1`

Basic features

- [x] apply header
- [x] apply query parameter
- encode request body as json if it is not `io.Reader`

Advanced features

- decode response body when 2xx
- decode response body when 4xx
  - this is why we always needed extra wrapper ...
- save response and request body

## Error handling

The basic error handling flow is 

````text
if ErrGetResponse {
   return ErrGetResponse
}
// now we have response, normally application error
if IsSuccessCode() {
    DecodeToStruct or Return raw response
} else {
    DecodeToErrorStruct 
    return ErrorStruct
}
````

There are two things we allow user to do

- error detection, `ErrorDetector(status, res) bool`
  - default just check status code range
- error handler, `ErrorHandler(status, body []byte, res) error`

Since error detection is quite common, might shrink the size of interface or just 
make it a function instead of interface, the drawback is you can't list implementations