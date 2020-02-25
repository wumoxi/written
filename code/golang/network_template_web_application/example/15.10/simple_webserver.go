package main

import (
	"io"
	"net/http"
)

const form = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>GoService</title>
</head>
<body>
<form action="#" method="post">
    <input type="text" name="in" />
    <input type="submit" value="submit" />
</form>
</body>
</html>
`

// handle a simple get request
func SimpleServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>hello, world!</h1>")
}

func FormServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		io.WriteString(w, r.FormValue("in"))
	}
}

func main() {
	http.HandleFunc("/test1", SimpleServer)
	http.HandleFunc("/test2", FormServer)
	err := http.ListenAndServe(":8892", nil)
	if err != nil {
		panic(err)
	}
}
