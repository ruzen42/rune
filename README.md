# rune

`rune` is a minimal privilege escalation utility for Linux and NixOS that works similarly to `sudo`. It executes commands as `root` (UID 0) if the calling user's UID matches the value stored in `/etc/runeid`.

## Features

- Access control based on `/etc/runeid`
- Executes commands with arguments (`rune <command> [args...]`)
- Full support for NixOS (`nix-shell`) and standard Linux distributions
- Environment sanitization to prevent privilege escalation vulnerabilities

## Setup and Installation

### 1. Configure `/etc/runeid`

Create `/etc/runeid` containing the UID of the allowed user:

```sh
echo $(id -u) | sudo tee /etc/runeid
sudo chmod 644 /etc/runeid
```

### 2. Build and Install

```sh
# Build binary and set Linux capabilities
make build

# Install binary to /usr/local/bin with SUID permission
sudo make install
```

## Makefile Commands

| Command | Description |
| --- | --- |
| `make` / `make all` | Format code and build |
| `make build` | Format code, build binary into `./bin/rune`, set capabilities |
| `make install` | Copy binary to `/usr/local/bin` and set SUID bit |
| `make fmt` | Format source code (`go fmt`) |
| `make run` | Build and run locally |
| `make clean` | Remove `./bin` directory |

## Usage

```sh
rune <command> [args...]
```

Examples:

```sh
# Check effective user
rune id

# List root directory
rune ls -la /root

# Read protected file
rune cat /etc/shadow
```

## License

This is free and unencumbered software released into the public domain (Unlicense).
