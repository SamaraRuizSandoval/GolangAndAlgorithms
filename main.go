package main

import (
	"fmt"
	"os"

	dependencyinjection "github.com/SamaraRuizSandoval/GolangAndAlgorithms/fundamentals/dependency_injection"
)

func main() {
	fmt.Println("Hello, World!")

	dependencyinjection.Greet(os.Stdout, "Elodie")
}
