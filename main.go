/*
https://www.programmersought.com/article/12297316144/
https://medium.com/@skdomino/watch-this-file-watching-in-go-5b5a247cf71f
*/

package main

import (
	"log"

	iftopParse "github.com/common-guid/pfexporter/iftop"
	"github.com/fsnotify/fsnotify"
)

func exporter() {
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
					iftopParse.Parse()

				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("/home/guid/Work/go-projects/pfexporter/iftop/file_watch/if2.txt")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
