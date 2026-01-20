package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		entry, err := ParseLine(scanner.Text())

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("timestamp: ", (entry.Timestamp))
		fmt.Println("level: ", (entry.Level))
		fmt.Println("fields: ", (entry.Fields))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
