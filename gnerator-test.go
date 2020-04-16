package main

import (
"fmt"

"github.com/moznion/gowrtr/generator"
)

func main() {
	generator := generator.NewRoot(generator.NewRawStatement(`type Model struct {
	name string
}`),)

	generated, err := generator.Generate(0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(generated)
}
