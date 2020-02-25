package main

import (
	"fmt"
	"log"
	"net/http"
)

const form = `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Statistics</title>
    <style>
        table, th, td {
            border: 1px solid black;
        }
        td {
            text-align: left;
        }
        .container  {
            text-align: center;
        }
    </style>
</head>
<body>
<div class="container">
    <h1>Statistics</h1>
    <h2>Computes basic statistics for a given list of numbers</h2>
    <form action="/" method="post">
        <input type="text" name="numbers" />
        <input type="submit" value="Calculate">
    </form>
	%s
</div>
</body>
</html>
`

const result = `
<table style="width:50%; margin: 2rem auto; ">
    <tr>
        <th colspan="2">Result</th>
    </tr>
    <tr>
        <td width="20%">Numbers</td>
        <td>%v</td>
    </tr>
    <tr>
        <td>Count</td>
        <td>%d</td>
    </tr>
    <tr>
        <td>Mean</td>
        <td>%f</td>
    </tr>
    <tr>
        <td>Count</td>
        <td>%f</td>
    </tr>
</table>
`

func HTML(str string, result string) string {
	return fmt.Sprintf(str, result)
}

func main() {
	http.HandleFunc("/", homePage)
	err := http.ListenAndServe(":8895", nil)
	if err != nil {
		log.Fatalf("Listen and server error: %s", err)
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	request.Header.Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		fmt.Fprint(writer, HTML(form, ""))
	case "POST":
		log.Println(request.ParseForm())
		log.Println(request.Form)
	}
}
