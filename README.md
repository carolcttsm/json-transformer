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

The output will go always in a file called `input.json` which will be generated with the applied modifications, but the order of keys may be altered due to the mapping function not preserving the order.

## Contact

Carolina Medeiros - @carolcttsm

Project Link: https://github.com/carolcttsm/json-transformer
