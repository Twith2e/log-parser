package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	argument := os.Args
	if len(argument) != 2 {
		log.Fatal("usage: logparser <logfile>")
	}

	file, err := os.Open(argument[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		timestamp, level, rest, ok := splitHeader(scanner.Text())
		if !ok {
			fmt.Println("error reading file line")
			continue
		}
		err := validTimestamp(timestamp)
		if err != nil {
			fmt.Println("invalid timestamp")
			continue
		}
		isLevel := validLevel(level)
		if !isLevel {
			fmt.Println("invalid level")
			continue
		}
		parsedRest := parseKeyValues(rest)
		fmt.Println()
		fmt.Println("ts: ", timestamp)
		fmt.Println()
		fmt.Println("lvl: ", level)
		fmt.Println()
		fmt.Println("rest: ", parsedRest)
		fmt.Println()

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func splitHeader(line string) (timestamp, level, rest string, ok bool) {
	parts := strings.SplitN(line, " ", 3)

	if len(parts) < 3 {
		return "", "", "", false
	}

	return parts[0], parts[1], parts[2], true
}

func validTimestamp(ts string) error {
	_, err := time.Parse(time.RFC3339, ts)
	return err
}

func validLevel(lvl string) bool {
	switch lvl {
	case "INFO", "ERROR", "WARN", "DEBUG":
		return true
	default:
		return false
	}

}

func parseKeyValues(input string) map[string]string {
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

	return fields
}
