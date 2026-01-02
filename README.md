# Port Killer

[한국어](README_ko.md)

Port Killer is a simple Go CLI tool that finds processes listening on TCP ports and allows you to easily terminate them.

It is useful for identifying which process is occupying a port and immediately killing it when you encounter an "Address already in use" error during development.

## Features

- Lists TCP ports in LISTEN state using the `lsof` command
- Displays process information in a user-friendly table format (Command, PID, User, Address)
- Interactive process termination via number selection
- Forcefully terminates using `kill -9` (SIGKILL) signal

## Requirements

- **OS**: Linux or macOS (Windows is currently not supported due to lack of `lsof` support)
- **Dependencies**: `lsof` must be installed.
- **Go**: Go is required to build or run the source code.

## Installation and Execution

### 1. Run Source Code

```bash
go run main.go
```

### 2. Build and Run

```bash
go build -o port-killer main.go
./port-killer
```

To use it globally, move the built binary to a directory included in your PATH (e.g., `/usr/local/bin`).

## Usage

When you run the program, a list of currently open ports is displayed.

```text
 List of currently running ports
No   Command    PID    User     Address
--   -------    ---    ----     -------
1    node       12345  user     *:3000
2    main       67890  user     *:8080

 Select a number to terminate (Cancel: 0):
```

Enter the number of the process you want to terminate and press Enter.

> **Note**: You may need `sudo` privileges to terminate system processes or processes owned by other users.

```bash
sudo ./port-killer
# or
sudo go run main.go
```

## Contribution

If you find a bug or have an improvement, please register an issue or send a PR.