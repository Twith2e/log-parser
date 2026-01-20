# LogParser

A simple Go program to parse structured log files.  
It reads a log file line by line, validates timestamps and log levels, and extracts key-value pairs from each line.

---

## Features

- Parses logs with a timestamp, level, and key-value fields.
- Validates timestamp format (RFC3339).
- Checks log levels: INFO, WARN, ERROR, DEBUG.
- Detects unclosed quotes in key-value fields.
- Outputs parsed log entries in a readable format.

---

## Log Format

Each log line should follow this structure:
`<timestamp> <level> key=value ...`

Example:
2026-01-12T08:16:01Z INFO user=105 action=view_page page=/dashboard msg="Dashboard viewed"

- `timestamp` → RFC3339 format (e.g., `2026-01-20T12:34:56Z`)
- `level` → INFO, WARN, ERROR, DEBUG
- `key=value` → key-value pairs; values may be quoted if they contain spaces

## Installation

1. Make sure you have [Go](https://golang.org/dl/) installed
2. Clone this repository:

   ```bash
       git clone http://github.com/Twith2e/log-parser
       cd log-parser
       go mod tidy
   ```

## Usage

```bash
    go run . app.log
```

Sample output:

```bash
    timestamp: 2026-01-12 08:24:55 +0000 UTC
    level: ERROR
    fields: map[action:auth msg:Invalid token provided user:999]
```

## Error handling

- running the command with an incomplete arg

```bash
    go run .
    usage: logparser <logfile>
```

- invalid timestamp -> prints `invalid timestamp`
- invalid log level -> prints `invalid level`
- unclosed quotes -> prints `unclosed quotes in line ...`

## File structure

logparser/
├── main.go # CLI and file reading
├── parser.go # Log parsing logic
├── app.log # Sample log file
├── go.mod # Go module
└── README.md
