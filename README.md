# code-golf

## Overview
`code-golf` is a Go project that includes a command-line tool for parsing data paths. The tool supports various features and can be extended with additional functionality.

## Features
- **Data Path Parser**: Parses and validates data paths with support for object and array references.
- **Data Getter**: Extracts values specified by a given path from an arbitrary structure built out of `map[string]any` or `[]any`.

## Installation
1. Clone the repository:

    ```sh
    git clone <repository-url>
    cd code-golf
    ```

2. Build the project:

    ```sh
    go build cmd/main.go
    ```

## Usage
Run the command-line tool with the desired feature and path:
```sh
./main -feature=data_path_parser -path='."multi"."path"'
```

```sh
./main -feature=data_getter -path='."people"."bob"."tags".-1."types"'
```
