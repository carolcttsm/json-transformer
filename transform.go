package main

import (
	"encoding/json" // JSON files handling
	"fmt"           // Printing messagens in the command line
	"os"            // File and error treatment
	"strings"       // String handling
)

func main() {
	data := parseInput("input.json")

	Transform(data, os.Args[1:])

	writeOutput("output.json", data)
}

// Function to open and parse input file and crate input data
func parseInput(fileName string) map[string]interface{} {
	// Opening input.json
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	// Decoding the JSON input file
	decoder := json.NewDecoder(jsonFile)

	// Creating a map to store JSON data
	var data map[string]interface{}

	// Decoding JSON data and storing in the map
	if err := decoder.Decode(&data); err != nil {
		fmt.Printf("Failed to parse JSON: %v\n", err)
	}

	return data
}

// Function to transform the input data in the desired output data
func Transform(data map[string]interface{}, commands []string) {
	// Handling the commands
	for _, command := range commands {
		parts := strings.Split(command, ":") // parsing the arguments
		if len(parts) < 2 {
			fmt.Println("Invalid command:", command)
			os.Exit(1)
		}
		command := parts[0] // if there are arguments in the program execution, then case handling is needed
		switch command {
		case "rename": // for `rename` transformation
			if len(parts) != 3 { // 3 items are needed: the command (rename), the key to change and the new key
				fmt.Println("Invalid command (rename):", command)
				os.Exit(1)
			}
			data[parts[2]] = data[parts[1]] // creates new key with the same value of the old one
			delete(data, parts[1])          // deletes the pair identified by the old key
		case "set": // for `set` transformation
			if len(parts) != 3 { // 3 items are needed: the command (set), the key whose value to change and the new value
				fmt.Println("Invalid command (set):", command)
				os.Exit(1)
			}
			data[parts[1]] = parts[2] // replaces the old value by the new one
		case "delete": // for `delete` transformation
			if len(parts) != 2 { // 2 items are needed: the command (delete), the key whose information to delete
				fmt.Println("Invalid command (delete):", command)
				os.Exit(1)
			}
			key := parts[1]
			delete(data, key)

		default:
			fmt.Println("Unknown command:", command) // Handling error of bad written commands
			os.Exit(1)
		}
	}
}

// Function to create the output file and write the output data in it
func writeOutput(fileName string, data map[string]interface{}) {
	// Opening output file
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer output.Close()

	// Creating an output JSON file
	encoder := json.NewEncoder(output)
	encoder.SetIndent("", "\t")

	// Encoding map data and storing in the file
	if err := encoder.Encode(data); err != nil {
		fmt.Printf("Failed to write JSON: %v\n", err)
		return
	}

	fmt.Println("JSON file succesfully modified!")
}
