# gusers

`gusers` is a Go-based CLI application that replicates the functionality of the `users` command in Linux. It reads from `utmpx` files to display a list of logged-in users on the system. The application is built using Go and includes a library to interact with `utmpx`.

## Features

- **Lightweight and Fast**: Written in Go for high performance.
- **Library Support**: Includes a reusable library for interacting with `utmpx`.
- **Help and Version Commands**: Supports common CLI arguments `--help` and `--version`.

## Installation

You can install `gusers` using the `go install` command:

```bash
go install github.com/drunkleen/gusers@latest
```

## Usage

### Command-line Interface

```bash
$ gusers [OPTIONS]
```

#### Options:
- `--help`: Display usage information.
- `--version`: Show version information.

Example:
```bash
$ gusers
user1 user2 user3
```

### Library Usage

You can also use the `utmpx` library in your Go projects to interact with `utmpx` files.

#### Install the library using `go get`:
```bash
go get github.com/drunkleen/gusers/utmpx
```

#### Example:
```go
package main

import (
	"fmt"
	"github.com/drunkleen/gusers/utmpx"
)

func main() {
	entries := utmpx.GetAll()
	for _, entry := range entries {
		fmt.Println(entry.User)
	}
}
```

## Project Structure
****
```
.
├── go.mod
├── LICENSE
├── main.go
├── README.md
├── utils
│   └── args.go
└── utmpx
    ├── helpers.go
    ├── libs.go
    ├── types.go
    └── utmpx.go
```

### Key Components:
- **`main.go`**: Entry point for the application.
- **`utils/args.go`**: Handles command-line argument parsing.
- **`utmpx/`**: Contains the core library for interacting with `utmpx` files.

## Development

### Prerequisites

- Go 1.19 or higher

### Build Instructions

1. Ensure Go is installed on your system.

2. Clone the repository:
   ```bash
   git clone https://github.com/drunkleen/gusers.git
   ```
3. Navigate to the project directory:
   ```bash
   cd gusers
   ```
4. Run the following command to build the application:
   ```bash
   go build -o bin/gusers
   ```

### Running Tests

(Currently, no automated tests are included for now. Future versions will address this.)

## License

This project is licensed under the GPLv3+ license. See the [LICENSE](LICENSE) file for more details.

## Acknowledgments

- Inspired by the GNU Coreutils `users` command and also [Jadi](https://github.com/jadijadi).
- Built with love by [DrunkLeen](https://github.com/drunkleen).


## Links

- [GitHub Repository](https://github.com/drunkleen/gusers)
- [GNU Coreutils](https://www.gnu.org/software/coreutils/)
