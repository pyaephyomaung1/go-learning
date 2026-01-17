package main

import "fmt"

func main() {
	var port = 8080
	if port == 8080 {
		fmt.Println("Running on default port")
	} else {
		fmt.Println("Running on custom port")
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	env := "dev"
	switch env {
	case "dev":
		fmt.Println("Development")
	case "prod":
		fmt.Println("Production")
	}

}
