# go-log

> ⚠️ **Work in Progress**: This library is currently under active development. Some features are still being implemented.

A beautiful, structured logging library for Go with rich styling and theming support using [Charm's Lipgloss](https://github.com/charmbracelet/lipgloss).

## Features

### Current Features ✅
- 🎨 **Multiple Built-in Themes**: Monokai, Dracula, Gruvbox, Nord, Tokyo Night, and more
- 🔍 **Structured Logging**: Log with fields, objects, maps, and errors
- 📝 **Pretty Printing**: Automatic formatting of complex data structures
- 🎯 **Caller Information**: Automatic file, line, and function tracking
- 🚀 **Performance**: Uses sync.Pool for buffer reuse
- 🎭 **Customizable**: Create your own color schemes and formatters

### Planned Features 🚧
- 📄 **File Writers**: Log rotation and file-based logging
- 🗄️ **MongoDB Integration**: Direct logging to MongoDB collections
- 🏭 **Writer Factory**: Decorator pattern for composable log writers
- 🔄 **Formatter Factory**: Easy switching between different log formats (JSON, plain text, etc.)
- 🔌 **Multi-Writer Support**: Log to multiple destinations simultaneously

## Installation

```bash
go get github.com/mohammedaouamri5/go-log
```

## Quick Start

```go
package main

import (
    "os"
    "github.com/mohammedaouamri5/go-log/log"
)

func main() {
    // Initialize the logger
    log.INIT(
        os.Stdout,
        log.NewStructuredLoggerFormatter(nil), // nil uses default theme
    )

    // Basic logging
    log.Info("Application started")
    log.Warn("This is a warning")
    log.Error("Something went wrong")
    log.Fatal("Critical error")
}
```

## Logging Levels

The library supports four logging levels:

- **INFO**: Informational messages
- **WARN**: Warning messages
- **ERROR**: Error messages
- **FATAL**: Critical errors

## Structured Logging

### Log with Fields

```go
log.WithField("user_id", 123).
    WithField("action", "login").
    Info("User logged in")
```

### Log with Maps

```go
data := map[string]interface{}{
    "name": "John Doe",
    "age": 30,
    "active": true,
}
log.WithMap(data).Info("User data")
```

### Log with Objects

```go
type User struct {
    ID    int
    Name  string
    Email string
}

user := &User{
    ID:    1,
    Name:  "Alice",
    Email: "alice@example.com",
}
log.WithObj(user).Info("User created")
```

### Log with Errors

```go
err := errors.New("connection failed")
log.WithErr(err).Error("Database connection error")
```

## Custom ToMap Method

You can implement a `ToMap()` method on your structs to control how they're serialized:

```go
type User struct {
    ID       int
    Name     string
    Password string // sensitive field
}

func (u *User) ToMap() map[string]any {
    return map[string]any{
        "ID":   u.ID,
        "Name": u.Name,
        // Password is excluded for security
    }
}
```

## Available Themes

The library includes several pre-built themes:

- `DefaultTheme()` - Default color scheme
- `MonokaiTheme()` - Monokai inspired
- `DraculaTheme()` - Dracula inspired
- `GruvboxTheme()` - Gruvbox inspired
- `SolarizedDarkTheme()` - Solarized Dark
- `NordTheme()` - Nord inspired
- `OneDarkTheme()` - One Dark inspired
- `TokyoNightTheme()` - Tokyo Night inspired
- `AyuDarkTheme()` - Ayu Dark inspired
- `LightTheme()` - Light theme

### Using a Theme

```go
log.INIT(
    os.Stdout,
    log.NewStructuredLoggerFormatter(log.DraculaTheme()),
)
```

## Custom Themes

Create your own color scheme:

```go
customTheme := &log.Colors{
    Info:  lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")).Bold(true),
    Warn:  lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFF00")).Bold(true),
    Error: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true),
    Fatal: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Background(lipgloss.Color("#000000")).Bold(true),
    
    Time: lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")).Italic(true),
    File: lipgloss.NewStyle().Foreground(lipgloss.Color("#00FFFF")).Bold(true),
    Line: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF")).Bold(true),
    Func: lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFF00")).Italic(true),
    
    KeyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF")).Bold(true),
    TypeStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")).Italic(true),
    StrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")),
    NumStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#00FFFF")),
    BoolStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF8800")).Bold(true),
    NullStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")).Italic(true),
    PtrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#00BFFF")).Italic(true),
    Bracket:   lipgloss.NewStyle().Foreground(lipgloss.Color("#AAAAAA")),
}

log.INIT(os.Stdout, log.NewStructuredLoggerFormatter(customTheme))
```

## Development

### Build

```bash
make build
```

### Run

```bash
make run
```

### Hot Reload

```bash
make hot_reload
```

### Test

```bash
make test
```

### Lint

```bash
make lint
```

## Log Output Format

The logger outputs structured logs with the following format:

<img width="312" height="422" alt="carbon" src="https://github.com/user-attachments/assets/96c4102c-0d65-4622-bf4a-a1960f6b51c5" />

Example output:

<img width="612" height="734" alt="carbon" src="https://github.com/user-attachments/assets/f1b4eb08-b1e1-4667-8b4a-02bb3fbd82bd" />



## Project Structure

```
.
├── log/
│   ├── log.go          # Core logger implementation
│   ├── Functions.go    # Formatting and pretty-printing
│   ├── Themes.go       # Pre-built color themes
│   └── constant.go     # Global logger instance
├── main.go             # Example usage
├── Makefile            # Build commands
└── README.md
```

## License

This project is open source and available under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Author

[Mohammed Aouamri](https://github.com/mohammedaouamri5)   

## Acknowledgments

- [Charm Lipgloss](https://github.com/charmbracelet/lipgloss) for the styling engine
- Inspired by popular logging libraries like logrus and zap


