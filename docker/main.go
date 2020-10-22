package main

import (
	"fmt"

	"kargos/docker/container"
)

func main() {
	fmt.Println(`Docker`)
	fmt.Println(`---`)

	container.Generate()
}
