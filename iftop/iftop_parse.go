package main

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("/home/guid/Work/go-projects/pfexporter/iftop/iftop.txt")

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
		//	println(num[line], ip[line], dir[line], two[line], ten[line], fourty[line], totes[line])
		line++
	}
	b40s := make(map[int]int)
	b10s := make(map[int]int)
	mre := regexp.MustCompile("M")
	kre := regexp.MustCompile("K")
	//bre := regexp.MustCompile("B")
	for line, value := range fourty {
		if mre.Match(([]byte(value))) {
			newval := regexp.MustCompile("MB").Split(value, -1)
			newval2 := strings.Join(newval[:], "")
			println(newval2, reflect.TypeOf(newval2))
			newval3 := newval2 + "000000"
			println(newval3, reflect.TypeOf(newval3))
			intval, _ := strconv.Atoi(newval3)
			println(intval, reflect.TypeOf(intval))
			//add to new list
			b40s[line] = intval
		} else if kre.Match(([]byte(value))) {
			newval := regexp.MustCompile("KB").Split(value, -1)
			newval2 := strings.Join(newval[:], "")
			newval3 := newval2 + "000"
			intval, _ := strconv.Atoi(newval3)
			//add to new list
			b40s[line] = intval
		} else {
			newval := regexp.MustCompile("B").Split(value, -1)
			newval2 := strings.Join(newval[:], "")
			intval, _ := strconv.Atoi(newval2)
			finval := intval * 1
			//add to new list
			b40s[line] = finval
		}
	}
	for line, value := range ten {
		if mre.Match(([]byte(value))) {
			newval := regexp.MustCompile("MB").Split(value, -1)
			newval2 := strings.Join(newval[:], "")
			newval3 := newval2 + "000000"
			intval, _ := strconv.Atoi(newval3)
			//add to new list
			b10s[line] = intval
		} else if kre.Match(([]byte(value))) {
			newval := regexp.MustCompile("KB").Split(value, -1)
			newval2 := strings.Join(newval[:], "")
			newval3 := newval2 + "000"
			intval, _ := strconv.Atoi(newval3)
			println(intval, reflect.TypeOf(intval))
			//add to new list
			b10s[line] = intval
		} else {
			newval := regexp.MustCompile("B").Split(value, -1)
			newval2 := strings.Join(newval[:], "")
			intval, _ := strconv.Atoi(newval2)
			//add to new list
			b10s[line] = intval
		}
	}
	// now we need to get it into /metrics

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
