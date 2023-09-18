package main

import (
	"fmt"
	"runtime"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)*/
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es Windows")
	} else {
		fmt.Println("Esto es", os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es", os)
	case "darwin":
		fmt.Println("Esto es", os)
	default:
		fmt.Printf("%s \n", os)
	}
}
