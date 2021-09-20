package main

import (
	"e-commerce/cmd/app/container"
	"fmt"
)

func main() {
	container := container.BuildContainer()
	fmt.Print(container)
}
