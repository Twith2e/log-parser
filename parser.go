package main

import (
	"fmt"
	"strings"
	"time"
)

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Fields    map[string]string
}

func ParseLine(line string) (*LogEntry, error) {
	ts, lvl, rest, ok := splitHeader(line)

	if !ok {
		return nil, fmt.Errorf("invalid header")
	}

	parsedTime, err := validTimestamp(ts)
	if err != nil {
		return nil, err
	}

	if !validLevel(lvl) {
		return nil, fmt.Errorf("invalid level")
	}

	fields, err := parseKeyValues(rest)

	if err != nil {
		return nil, err
	}

	return &LogEntry{
		Timestamp: parsedTime,
		Level:     lvl,
		Fields:    fields,
	}, nil
}

func validLevel(level string) bool {
	switch level {
	case "INFO", "WARN", "ERROR", "DEBUG":
		return true
	default:
		return false
	}
}

func validTimestamp(ts string) (time.Time, error) {
	parsedTime, err := time.Parse(time.RFC3339, ts)
	return parsedTime, err
}

func splitHeader(line string) (timestamp string, level, rest string, ok bool) {
	parts := strings.SplitN(line, " ", 3)

	if len(parts) < 3 {
		return "", "", "", false
	}

	return parts[0], parts[1], parts[2], true
}

func parseKeyValues(input string) (map[string]string, error) {
	fields := make(map[string]string)

	var key, value strings.Builder
	inQuotes := false
	readingKey := true

	for i := 0; i < len(input); i++ {
		ch := input[i]

		switch ch {
		case '=':
			if readingKey {
				readingKey = false
			} else {
				value.WriteByte(ch)
			}
		case '"':
			inQuotes = !inQuotes
		case ' ':
			if inQuotes {
				value.WriteByte(ch)
			} else {
				if key.Len() > 0 {
					fields[key.String()] = value.String()
					key.Reset()
					value.Reset()
					readingKey = true
				}
			}
		default:
			if readingKey {
				key.WriteByte(ch)
			} else {
				value.WriteByte(ch)
			}
		}

	}

	if key.Len() > 0 {
		fields[key.String()] = value.String()
	}

	if inQuotes {
		return nil, fmt.Errorf("unclosed quotes in line %s\n", input)
	}

	return fields, nil
}
