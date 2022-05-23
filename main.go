/*
https://www.programmersought.com/article/12297316144/
https://medium.com/@skdomino/watch-this-file-watching-in-go-5b5a247cf71f
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	iftopParse "github.com/common-guid/pfexporter/iftop"
	"github.com/fsnotify/fsnotify"
)

var fileLoc string = "/home/guid/Work/go-projects/pfexporter/if2.txt"

func main() {
	http.HandleFunc("iftop/iftop.txt", func(rw http.ResponseWriter, req *http.Request) { // testing this <---------
		http.ServeFile(rw, req, req.URL.Path[1:]) // testing this <--------------

	})
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					//  log.Println("modified file:", event.Name)

					// I think prom is blocking the loop from continuing
					if fileLen() > 1 {
						iftopParse.Parse()
						//delete file - put somewhere
						//os.Remove(fileLoc)
						iftopParse.Prom()
						fmt.Println("past prom")
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(fileLoc)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

/* make sure this is working as intended
first time with return and declaring int after function
count lines in iftop file to be used in if stmt in main()
*/
func fileLen() int {
	// wc implementation from https://amehta.github.io/posts/2019/03/wc-implementation-in-go-lang/
	file, err := os.Open(fileLoc)
	if err != nil {
		fmt.Println("Err ", err)
	}
	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}
	return lines
}
