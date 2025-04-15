# IO Library

The `io` library provides input/output operations for JotLang.

## Import

```jt
import "io"
```

## Functions

### Output Functions

```jt
fn print(message string) : void
```
Prints a message to the standard output without a newline.

```jt
fn println(message string) : void
```
Prints a message to the standard output with a newline.

### Input Functions

```jt
fn readLine() : string
```
Reads a line from standard input.

```jt
fn readInt() : int
```
Reads a line from standard input and converts it to an integer.

```jt
fn readFloat() : float
```
Reads a line from standard input and converts it to a float.

### File Operations

```jt
fn readFile(path string) : string
```
Reads the entire contents of a file.

```jt
fn writeFile(path string, content string) : void
```
Writes content to a file, overwriting if it exists.

```jt
fn appendFile(path string, content string) : void
```
Appends content to a file.

```jt
fn fileExists(path string) : bool
```
Checks if a file exists.

### Formatting

```jt
fn format(template string, ...args any) : string
```
Formats a string using the provided arguments.

## Examples

```jt
import "io"

// Basic output
io.println("Hello, World!")

// String interpolation
name = "JotLang"
io.println("Welcome to {name}!")

// File operations
content = "Some content to write"
io.writeFile("example.txt", content)

if io.fileExists("example.txt") {
    fileContent = io.readFile("example.txt")
    io.println("File content: {fileContent}")
}

// User input
io.print("Enter your name: ")
userName = io.readLine()
io.println("Hello, {userName}!")

io.print("Enter your age: ")
age = io.readInt()
io.println("You are {age} years old")

// Formatting
message = io.format("Hello, {0}! You are {1} years old.", userName, age)
io.println(message)
```

## Best Practices

1. Always check if a file exists before reading it
2. Use string interpolation for simple string formatting
3. Use `format` for complex string formatting with multiple arguments
4. Close files after use (handled automatically by the library)
5. Use `println` for general output and `print` when you need to control line breaks 