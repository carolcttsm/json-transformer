# JSON Transformer

JSON Transformer is a command-line tool written in Go that allows you to load JSON data from a file, modify it according to the passed arguments, and save the changes to a new JSON files.

## Features

- Load JSON data from a file.
- Delete specified keys.
- Set values for specified keys.
- Rename specified keys.
- Save the modified JSON data to a new file, preserving the order of keys.

## Requirements

- [Go](https://golang.org/doc/install) 1.16 or higher

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/carolcttsm/json-transform.git
    cd json-transform
    ```

2. Build the project:

    ```sh
    go build transform.go
    ```

## Usage

To use JSON Transformer, run the program with the following arguments:

```sh
./transform [operations]
```

## Examples

Preserve all information:

    ```sh
    ./transform
    ```
Delete a key:

    ```sh
    ./transform delete:key
    ```

Set the value of a key:

    ```sh
    ./transform set:newKey:Value
    ./transform set:existingKey:newValue
    ```

Rename a key:

    ```sh
    ./transform rename:key:newKey
    ```

Combine multiple operations:

    ```sh
    ./transform delete:key_1 set:key_2:newValue rename:key_3:newKey
    ```

## Input JSON Format

To keep it simple, the input is always in a file called `input.json` in the current directory.

The input file should be in valid JSON format, assuming that each key or value consists of one or more letters or numbers, both for the file and the keys/values given on the command-line (this project doesn't handle cases such as keys that contain a colon `:` or a space character ` `).

For example:

```json
    {
        "color": "blue",
        "number": "two",
        "pet": "cat"
    }
```

## Output JSON Format

The output will go always in a file called `output.json` which will be generated with the applied modifications, but the order of keys may be altered due to the mapping function not preserving the order.

## Development

### Motivation

The motivation behind creating JSON Transformer was to provide a simple and efficient command-line tool for manipulating JSON data. This tool can be useful for developers, data engineers, and anyone who frequently works with JSON data.

### Technologies Used

    Go: The programming language chosen for its performance, simplicity, and powerful standard library for handling JSON.

### Development Process

1. Initial Setup:

    Set up the Go environment and initialize a new Go module.
    Create a basic project structure with transform.go as the main file.

2. Loading JSON Data:

    Implement functionality to read JSON data from a file using json.Decoder.
    Ensure the JSON data is loaded into a structure that can be updated.

3. Modifying JSON Data:

    Hadles cases to delete keys, set values for keys, and rename keys within the JSON structure.
   
5. Saving Modified JSON Data:

    Implement functionality to write the modified JSON data back to a file using json.Encoder.

6. Command-Line Interface:

    Implement the corresponding changes to modify the JSON data based on the provided arguments.

### Key Functions

- `parseInput`: opens and parses input file and create input data.

- `Transform`: changes the input data with the desired specifications for the output data.

- `writeOutput`: creates the output file and writes the output data in it.

### Challenges and Solutions

    Key Order: Go maps do not preserve the order of keys.
    Command-Line Parsing: Handling various operations via command-line arguments required careful parsing and validation.

## Contact

Carolina Medeiros - @carolcttsm

Project Link: https://github.com/carolcttsm/json-transformer
