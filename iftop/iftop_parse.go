package main

import (
	"bufio"
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
	fileTextLines := make(map[int]string)
	i := 1
	for fileScanner.Scan() {
		fileTextLines[i] = fileScanner.Text()
		i++
	}
	// this has the indended effect of splitting words out
	readFile.Close()

	start, end, line := 3, (len(fileTextLines) - 9), 1

	num := make(map[int]string)
	ip := make(map[int]string)
	dir := make(map[int]string)
	two := make(map[int]string)
	ten := make(map[int]string)
	fourty := make(map[int]string)
	totes := make(map[int]string)

	for line <= end {
		if line > start {
			words := strings.Fields(fileTextLines[line])
			if line%2 == 0 {
				num[line], ip[line], dir[line], two[line], ten[line], fourty[line], totes[line] = words[0], words[1], words[2], words[3], words[4], words[5], words[6]
			} else {
				ip[line], dir[line], two[line], ten[line], fourty[line], totes[line] = words[0], words[1], words[2], words[3], words[4], words[5]
			}
		}
		println(num[line], ip[line], dir[line], two[line], ten[line], fourty[line], totes[line])
		line++
	}
	/*
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
			   https://haydz.github.io/2020/04/12/ParsingStrings.html
		*
		i := 1
		//reader := bufio.NewReader(fileTextLines)
		for _, eachline := range fileTextLines {
			//scanner := reader.ReadString(eachline)
			print(eachline)
			i++
		}*/
}
