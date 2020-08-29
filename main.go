package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

func syncFile(remote string, path string) {

	args := strings.Split(fmt.Sprintf("rsync -a %s %s", path, remote), " ")
	log.Println(args)
	command := exec.Command(args[0], args[1:]...)
	outStream, _ := command.StdoutPipe()
	errStream, _ := command.StderrPipe()

	command.Start()

	output, _ := ioutil.ReadAll(outStream)
	errput, _ := ioutil.ReadAll(errStream)

	command.Wait()
	log.Println(string(errput))
	fmt.Println(string(output))
	return
}

func track(path string, wg *sync.WaitGroup) {
	defer wg.Done()

	filename := path

	lastModifiedTime := time.Now()
	for {
		// get last modified time
		file, err := os.Stat(filename)

		if err != nil {
			fmt.Println(err)
		}
		modifiedTime := file.ModTime()
		if lastModifiedTime != modifiedTime {
			syncFile("root@172.16.8.223:/root/ArshiA", path)
			fmt.Println("Last modified time : ", modifiedTime)
		}
		lastModifiedTime = modifiedTime
	}
}

// func main() {

// 	var wg sync.WaitGroup

// 	wg.Add(1)
// 	go track("a.txt", &wg)
// 	wg.Wait()
// }

func main() {
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
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("./")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
