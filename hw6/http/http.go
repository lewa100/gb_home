package http

import (
	"fmt"
	"io"
	"net/http"
)

// reqParser get params from request.
func reqParser(r *http.Request) string{
	query := r.URL.Query()
	p := ""
	for k,v := range query{
		p += fmt.Sprintf("%s = %s<br>",k ,v)
	}
	return p
}

// Hello create response.
func Hello(w http.ResponseWriter, r *http.Request) {
	params := reqParser(r)
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w,
		`<doctype html>
<html>
	<head>
   	<title>Hello World!</title>
	</head>
	<body>
   	Hello World!
	<br />`+
	params+`</body>
</html>`)
}
