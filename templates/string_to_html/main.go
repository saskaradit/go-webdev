package main

import "fmt"

func main() {
	name := "Raditya Saskara"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
		<body>
			<h1>` + name + `</h1> 
		</body>
	</html>`

	fmt.Println(tpl)
}

// go run main.go > index.html
