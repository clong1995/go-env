package main

import (
	"fmt"
)

func main() {
	port := env.Env("PORT")
	fmt.Println(port)
}
