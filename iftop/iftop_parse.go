package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("iftop.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	// split file by line
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}
	readFile.Close()

	// for each line from file process based on string position
	var ip []string
	//var eline []string
	for _, eachline := range fileTextLines {
		scanner := bufio.NewScanner(strings.NewReader(eachline))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			ip = append(ip, scanner.Text())
		}
	}

	fmt.Println(ip)

	/*
	   https://gist.github.com/jamesrr39/c45a1aff4d3fd9dc2949
	   https://bountify.co/golang-parse-stdout
	*/
}
