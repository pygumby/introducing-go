package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<!DOCTYPE html>
		<html>
			<head>
				<title>Hello, world.</title>
			</head>
			<body>
				<p>Hello, world.</p>
			</body>
		</html>`,
	)
}

func main() {
	// To test go to http://127.0.0.1:9000/hello
	http.HandleFunc("/hello", hello)
	// To test go to http://127.0.0.1:9000/assets
	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets")),
		),
	)
	http.ListenAndServe(":9000", nil)
}
