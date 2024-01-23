package main

import (
	"fmt"

	"github.com/igor-cotrim/exercism-go/easy"
)

func main() {
	fmt.Println("RAINDROPS:", easy.Convert(21))
	result, _ := easy.CollatzConjecture(21)
	fmt.Println("CollatzConjecture:", result)
}
