# LogExtended (Educational Project)

**LogExtended** is a small educational logging library written in Go.
This project was created for learning purposes only and **will not be maintained or further developed**.
It was implemented as part of the course ["Go Basics"](https://github.com/buzzismaloy/gonalg-basics)

The goal of the project is to demonstrate how to extend Go’s standard `log.Logger` with basic log levels.

---

## Features

- Simple log levels: `ERROR`, `WARNING`, `INFO`
- Custom logger wrapper around Go standard `log` package
- Runtime log level filtering
- Minimal and easy-to-read implementation

---

## Installation

```bash
go get github.com/buzzismaloy/go-logextended
```

---

## Usage example

```go
func main() {
    logger := logextended.NewLogExtended()
    logger.SetLogLevel(logextended.LogLevelWarning)

    logger.Infoln("It must not be printed")
    logger.Warnln("Hello")
    logger.Errorln("World")
    logger.Println("Debug")
}
```

## Log Levels

```go
const (
    LogLevelError LogLevel = iota
    LogLevelWarning
    LogLevelInfo
)
```

- `LogLevelError` — prints only errors

- `LogLevelWarning` — prints warnings and errors

- `LogLevelInfo` — prints info, warnings, and errors (default)

## Project Structure

```
.
├── go.mod
└── logextended
    └── logextended.go
```

## Educational Notice

This project is for **educational purposes only**.
It is not intended for production use and **will not be actively maintained or extended**.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/buzzismaloy/go-logextended/blob/main/LICENSE) file for details.
