# Go Microservice: Syntax and Constructs Crash Course

This guide explains the Go concepts and syntax used in this microservice project.

## 1. Project Structure and Packages
Go code is organized into **packages**. Every file must start with a `package` declaration.
- `package main`: The entry point for the executable.
- `import`: Used to include code from other packages.
- **Exporting**: Identifiers (functions, structs, variables) starting with a **Capital Letter** are exported (public) to other packages.

## 2. Variables and Data Types
Go is statically typed.
- `var DB *sql.DB`: Package-level variable with explicit type.
- `tokenStr := parts[1]`: Short declaration (type inference) inside functions.
- `float64`, `string`, `bool`, `int`: Common primitive types.

## 3. Structs and JSON Tags
Structs are collections of fields, used like classes without inheritance.
```go
type Product struct {
    ID    string  `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price,string"`
}
```
- **Tags**: `` `json:"id"` `` metadata tells the JSON encoder/decoder how to map fields.

## 4. Functions and Multiple Returns
Functions can return multiple values, often used for error handling.
```go
func GenerateToken(username string) (string, error) {
    // ... logic
    return token, nil
}
```

## 5. Error Handling
Go doesn't use try/catch. instead, errors are returned as values.
```go
if err != nil {
    log.Fatal(err)
}
```
Always check the `err` variable immediately after calling a function that returns it.

## 6. Slices
Slices are dynamic arrays.
- `parts := strings.Split(authHeader, " ")`: Returns a slice of strings.
- `var products []models.Product`: A slice of `Product` structs.

## 7. Pointers
Pointers (`*`) store the memory address of a value.
- `claims := &Claims{}`: Creates a pointer to a new `Claims` struct.
- Used for performance (avoid copying) or to allow a function to modify a variable.

## 8. Concurrency (Goroutines & Channels)
The "Consumer" pattern uses Go's unique concurrency model.
- **Goroutine**: `go func() { ... }()` starts a lightweight thread.
- **Channel**: `MessageBus := make(chan models.Product, 10)` creates a pipe for goroutines to communicate.
- **Select/Range**: `for product := range MessageBus` waits for and receives data from a channel until it's closed.

## 9. Flow Control
- **If**: `if !limiter.Allow() { ... }`
- **For**: `for rows.Next() { ... }` (Go only has `for` loops, no `while`).
- **Defer**: `defer rows.Close()` ensures a function call runs at the very end of the current function (usually for cleanup).

## 10. Blank Identifier `_`
Used to ignore values you don't need.
- `_ "github.com/mattn/go-sqlite3"`: Side-effect import (runs the package's `init()` function).
- `statement, _ := DB.Prepare(...)`: Ignores the error returned by `Prepare`.
