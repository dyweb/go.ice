package openapi_test

import (
	"net/http"
	"testing"
)

// redoc_test.go servers a openapi spec with redoc
// TODO: skip the test when playground is not required ...

var redocTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <title>ReDoc</title>
    <!-- needed for adaptive design -->
    <meta charset="utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">

    <!--
    ReDoc doesn't change outer page styles
    -->
    <style>
      body {
        margin: 0;
        padding: 0;
      }
    </style>
  </head>
  <body>
    <redoc spec-url='/petstore.yaml'></redoc>
    <script src="https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js"> </script>
  </body>
</html>
`

func TestServe(t *testing.T) {
	t.Skip("uncomment it manually")

	mux := http.NewServeMux()
	mux.HandleFunc("/petstore.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "petstore.yaml")
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(redocTmpl))
	})
	http.ListenAndServe(":3000", mux)
}
