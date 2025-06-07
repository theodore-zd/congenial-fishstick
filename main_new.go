package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <html-file>")
		fmt.Println("Example: go run . testme.html")
		os.Exit(1)
	}

	htmlFile := os.Args[1]

	// Read the HTML file
	htmlContent, err := ioutil.ReadFile(htmlFile)
	if err != nil {
		log.Fatalf("Error reading HTML file: %v", err)
	}

	// Create HTML parser
	parser := NewHTMLParser()

	// Process the HTML content
	result, err := parser.ProcessHTMLContent(string(htmlContent))
	if err != nil {
		log.Fatalf("Error processing HTML: %v", err)
	}

	// Print statistics
	fmt.Printf("=== Tailwind CSS Generator Results ===\n")
	fmt.Printf("Total classes found: %d\n", result.ClassCount)
	fmt.Printf("Unique classes: %d\n", result.UniqueClassCount)

	// Print some of the extracted classes
	fmt.Printf("\n=== Sample Extracted Classes ===\n")
	count := 0
	for class := range result.ExtractedClasses.UniqueClasses {
		if count >= 10 { // Show first 10 classes
			break
		}
		fmt.Printf("- %s\n", class)
		count++
	}

	if len(result.ExtractedClasses.UniqueClasses) > 10 {
		fmt.Printf("... and %d more classes\n", len(result.ExtractedClasses.UniqueClasses)-10)
	}

	// Print class statistics
	stats := result.ExtractedClasses.GetClassStatistics()
	fmt.Printf("\n=== Class Type Statistics ===\n")
	for category, count := range stats {
		fmt.Printf("%s: %d\n", category, count)
	}

	// Generate and save CSS
	fmt.Printf("\n=== Generated CSS ===\n")
	if len(result.GeneratedCSS) > 1000 {
		fmt.Printf("%s...\n[CSS truncated - full CSS saved to output.css]\n", result.GeneratedCSS[:1000])
	} else {
		fmt.Printf("%s\n", result.GeneratedCSS)
	}

	// Write CSS to file
	err = ioutil.WriteFile("output.css", []byte(result.GeneratedCSS), 0644)
	if err != nil {
		log.Printf("Warning: Could not write output.css: %v", err)
	} else {
		fmt.Printf("\nGenerated CSS saved to output.css\n")
	}

	// Demonstrate individual class processing
	fmt.Printf("\n=== Individual Class Examples ===\n")
	exampleClasses := []string{"flex", "text-red-500", "bg-blue-100", "p-4", "m-2", "hover:bg-gray-200"}

	generator := NewCSSGenerator()
	for _, className := range exampleClasses {
		rules, err := generator.GenerateCSS(className)
		if err != nil {
			fmt.Printf("%s: Error - %v\n", className, err)
			continue
		}

		for _, rule := range rules {
			css := generator.FormatCSSRule(rule)
			fmt.Printf("%s: %s\n", className, css)
		}
	}
}

// demonstrateUsage shows various ways to use the library
func demonstrateUsage() {
	fmt.Println("=== Tailwind CSS Generator Demo ===")

	// Example HTML content
	htmlContent := `
	<div class="flex items-center justify-between p-4 bg-white shadow-lg rounded-lg">
		<h1 class="text-2xl font-bold text-gray-800">Hello World</h1>
		<button class="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600">
			Click me
		</button>
	</div>
	`

	// Create parser and process HTML
	parser := NewHTMLParser()
	result, err := parser.ProcessHTMLContent(htmlContent)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Found %d unique classes\n", result.UniqueClassCount)
	fmt.Printf("Generated CSS:\n%s\n", result.GeneratedCSS)
}
