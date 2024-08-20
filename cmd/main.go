package main

import (
	"fmt"
	"github.com/JuanRulliansyah/pgklone/pkg"
	"os"
)

func main() {
	var sourceURL, targetURL string

	fmt.Print("Enter Source DB URL: ")
	_, err := fmt.Scanln(&sourceURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading source DB URL: %v\n", err)
		os.Exit(1)
	}

	fmt.Print("Enter Target DB URL: ")
	_, err = fmt.Scanln(&targetURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading target DB URL: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Source DB URL: %s\n", sourceURL)
	fmt.Printf("Target DB URL: %s\n", targetURL)

	if err := pkg.CloneDatabase(sourceURL, targetURL); err != nil {
		fmt.Fprintf(os.Stderr, "Error cloning database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Database cloned successfully!")
}
