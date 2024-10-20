# Mapify

## Overview

The `Mapify` function is a utility that converts a Go structure (or any value of type `any`) into a `map[string]any`. It leverages JSON marshaling and unmarshaling to achieve the conversion. This can be useful for serializing complex structures into a map format, which can then be manipulated more easily or used with other libraries that require data in `map[string]any` format.

## Function Signature

```go
func Mapify(v any) (map[string]any, error)
```

### Parameters

- `v`: The input value of type `any` that you want to convert into a map. This is usually a struct or any data type that can be marshaled to JSON.

### Returns

- `(map[string]any)`: The converted data in the form of a `map[string]any`.
- `(error)`: Returns an error if there are issues during marshaling or unmarshaling.

## How It Works

1. **Marshal**: The input value `v` is first marshaled into a JSON-encoded byte array using `json.Marshal`.
2. **Unmarshal**: The JSON-encoded byte array is then unmarshaled into a `map[string]any` using `json.Unmarshal`.
3. **Error Handling**: If there are any errors during the marshaling or unmarshaling process, they are returned. If everything succeeds, the resulting map is returned along with `nil` for the error.

## Example Usage

Here’s a quick example of how to use the `Mapify` function:

```go
package main

import (
    "fmt"
    "log"
    "encoding/json"
)

type Person struct {
    Name   string
    Age    int
    Active bool
}

func main() {
    // Create an instance of the struct
    person := Person{Name: "John Doe", Age: 30, Active: true}

    // Convert the struct to a map
    data, err := Mapify(person)
    if err != nil {
        log.Fatalf("Error converting struct to map: %v", err)
    }

    // Print the resulting map
    fmt.Println(data)
}
```

### Output:

```
map[Name:John Doe Age:30 Active:true]
```

## Error Handling

If the input value cannot be marshaled into JSON (e.g., if it contains unexported fields or circular references), the function will return an error.

For example:

```go
data, err := Mapify(invalidStruct)
if err != nil {
    fmt.Println("Error:", err)
}
```

## Dependencies

- The function relies on Go's built-in `encoding/json` package.
