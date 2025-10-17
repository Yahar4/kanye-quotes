# Kanye Quotes CLI

A beautiful terminal app that displays random Kanye West quotes.

## Install

```bash
go install github.com/Yahar4/kanye-quotes@latest
```

## Usage

```bash
kanye-quotes
```

**Controls:**
- `r`, `space`, `enter` - New quote
- `q`, `ctrl+c`, `esc` - Quit

## Build from Source

```bash
git clone https://github.com/Yahar4/kanye-quotes

cd kanye-quotes

go build -o kanye-quotes ./cmd/kanye-quotes

./kanye-quotes
```

## Requirements

- Go 1.21+
- Terminal with true color support
