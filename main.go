package main

import (
	"fmt"
)

// SSL Labs API entry point
const entryPoint = "https://api.ssllabs.com/api/v2/"

//

func main() {
	// If info is true, then the server is working ()
	info, err := getInfo(entryPoint)
	if err != nil {
		fmt.Println("Error:", err)
	} else if info != "" {
		fmt.Println("Servidor de SSL Labs activo")
		fmt.Println(info)
	}

	// Analizamos un dominio usando la función analyzeDomain
	var domain string
	fmt.Print("Ingrese el dominio: ")
	fmt.Scanln(&domain)
	fmt.Println("Dominio ingresado:", domain)

	grade, err := analyzeDomain(domain)

	if err != nil {
		fmt.Println("Error analyzing domain:", err)
	} else {
		fmt.Println("¡Completed!")
		fmt.Printf("The SSL grade for %s is: %s\n", domain, grade)
	}
}
