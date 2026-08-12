# Go Minimal Template

A lightweight, opinionated Go project starter template configured with a Makefile for rapid local development, build automation, and easy installation.

---

## Features

* **Automated Workflow:** Pre-configured Makefile for building, running, formatting, and installing binaries.
* **Version Injection:** Automatic compile-time versioning passed directly into `main.Version` via `-ldflags`.
* **Clean Structure:** Standardized output path for compiled binaries (`output/`).

---

## Prerequisites

* [Go](https://go.dev/doc/install) 1.26.4 or higher
* `make` utility (pre-installed on Linux/macOS, available via WSL or Chocolatey on Windows)

---

## Quick Start

### Clone and Run

```bash
# Clone the template repository
git clone https://github.com/your-username/your-repo-name.git
cd your-repo-name

# Build and run the project
make run

```

---

## Available Commands

All primary development tasks are managed via `make`:

| Target | Command | Description |
| --- | --- | --- |
| `all` | `make` or `make all` | Default target. Builds the project binary. |
| `build` | `make build` | Runs `go fmt`, creates `output/`, and compiles `output/main`. |
| `run` | `make run` | Compiles the binary and immediately executes it. |
| `fmt` | `make fmt` | Runs `go fmt` on the codebase to enforce standard formatting. |
| `install` | `make install` | Builds the binary and copies it to `/usr/local/bin/main`. |
| `clean` | `make clean` | Removes the `output/` build directory and compiled binary. |

---

## Build Customization

You can customize the build parameters directly in the `Makefile` or override them at runtime:

```makefile
BUILD_DIR=output
VERSION=0.1.0
NAME=main

```

### Overriding Variables at Runtime

```bash
# Build with a custom version tag
make build VERSION=1.2.3

# Build with a custom output binary name
make build NAME=myapp

```

---

## Project Structure

```text
.
├── output/           # Generated binary directory (git-ignored)
├── go.mod            # Go module definition
├── main.go           # Application entrypoint
├── Makefile          # Build automation script
└── README.md         # Project documentation

```

---

## License

This template is licensed under the [MIT License](https://www.google.com/search?q=LICENSE).
