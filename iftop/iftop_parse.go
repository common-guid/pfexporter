package iftopParse

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	num    = make(map[int]string)
	ip     = make(map[int]string)
	dir    = make(map[int]string)
	two    = make(map[int]string)
	ten    = make(map[int]string)
	fourty = make(map[int]string)
	totes  = make(map[int]string)
	b40s   = make(map[int]float64)
	b10s   = make(map[int]float64)
	bTotal = make(map[int]float64)
)

// need to delete file after parsing [ ]
// update from concatenating to doing math to get values because of decimals [x]
// add in cumulative list [x]
// file monitoring to fetch updated  [x]
func Parse() {
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
	for line, item := range dir {
		if item == "<=" {
			dir[line] = "dest"
		} else {
			dir[line] = "src"
		}
	}

	mre := regexp.MustCompile("M")
	kre := regexp.MustCompile("K")
	//compile the 40s list of values
	for line, value := range fourty {
		if mre.Match(([]byte(value))) {
			newval := regexp.MustCompile("MB").Split(value, -1)
			newval2 := strings.Join(newval[:], "")
			intval, _ := strconv.ParseFloat(newval2, 64)
			//	println(newval2, reflect.TypeOf(newval2))
			finval := intval * 1000000
			//	println(finval, reflect.TypeOf(finval))
			//add to new list
			b40s[line] = finval
		} else if kre.Match(([]byte(value))) {
			newval := regexp.MustCompile("KB").Split(value, -1)
			newval2 := strings.Join(newval[:], "")
			intval, _ := strconv.ParseFloat(newval2, 64)
			finval := intval * 1000
			//add to new list
			b40s[line] = finval
		} else {
			newval := regexp.MustCompile("B").Split(value, -1)
			newval2 := strings.Join(newval[:], "")
			finval, _ := strconv.ParseFloat(newval2, 64)
			//add to new list
			b40s[line] = finval
		}
	}
	// compile the 10s list of values
	for line, value := range ten {
		if mre.Match(([]byte(value))) {
			newval := regexp.MustCompile("MB").Split(value, -1)
			newval2 := strings.Join(newval[:], "")
			intval, _ := strconv.ParseFloat(newval2, 64)
			//	println(newval2, reflect.TypeOf(newval2))
			finval := intval * 1000000
			//	println(finval, reflect.TypeOf(finval))
			//add to new list
			b10s[line] = finval
		} else if kre.Match(([]byte(value))) {
			newval := regexp.MustCompile("KB").Split(value, -1)
			newval2 := strings.Join(newval[:], "")
			intval, _ := strconv.ParseFloat(newval2, 64)
			finval := intval * 1000
			//add to new list
			b10s[line] = finval
		} else {
			newval := regexp.MustCompile("B").Split(value, -1)
			newval2 := strings.Join(newval[:], "")
			finval, _ := strconv.ParseFloat(newval2, 64)
			//add to new list
			b10s[line] = finval
		}
	}
	// compile the cumulative list of values
	for line, value := range totes {
		if mre.Match(([]byte(value))) {
			newval := regexp.MustCompile("MB").Split(value, -1)
			newval2 := strings.Join(newval[:], "")
			intval, _ := strconv.ParseFloat(newval2, 64)
			//	println(line, newval2, reflect.TypeOf(newval2))
			finval := intval * 1000000
			//	println(line, finval, reflect.TypeOf(finval))
			//add to new list
			bTotal[line] = finval
		} else if kre.Match(([]byte(value))) {
			newval := regexp.MustCompile("KB").Split(value, -1)
			newval2 := strings.Join(newval[:], "")
			println(line, newval2, reflect.TypeOf(newval2))
			intval, _ := strconv.ParseFloat(newval2, 64)
			finval := intval * 1000
			println(line, finval, reflect.TypeOf(finval))
			//add to new list
			bTotal[line] = finval
		} else {
			newval := regexp.MustCompile("B").Split(value, -1)
			newval2 := strings.Join(newval[:], "")
			finval, _ := strconv.ParseFloat(newval2, 64)
			//add to new list
			bTotal[line] = finval
		}
	}
	/*
		   https://gist.github.com/jamesrr39/c45a1aff4d3fd9dc2949
		   https://bountify.co/golang-parse-stdout
		   https://haydz.github.io/2020/04/12/ParsingStrings.html

		   original:
		   if mre.Match(([]byte(value))) {
				newval := regexp.MustCompile("MB").Split(value, -1)
				newval2 := strings.Join(newval[:], "")
				finval := newval2 + "000000"
				intval, _ := strconv.Atoi(finval)
				//add to new list
				b10s[line] = intval
	*/
}
func Prom() {

	// need to check that values are associating correctly
	// figure out a better naming convention, maybe just ip
	for line, value := range dir {
		sline := strconv.Itoa(line)
		if value == "src" {
			var (
				ip_sender = prometheus.NewGauge(prometheus.GaugeOpts{
					Name:        "ip_sender" + sline,
					Help:        "ip addr of host sending communications",
					ConstLabels: prometheus.Labels{"ip": ip[line]},
				})
			)
			prometheus.MustRegister(ip_sender)
			ip_sender.Add(float64(b10s[line]))
		} else {
			var (
				ip_receiver = prometheus.NewGauge(prometheus.GaugeOpts{
					Name:        "ip_receiver" + sline,
					Help:        "ip addr of host receiving communications",
					ConstLabels: prometheus.Labels{"ip": ip[line]},
				})
			)
			prometheus.MustRegister(ip_receiver)
			ip_receiver.Add(float64(b10s[line]))
		}
	}
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
