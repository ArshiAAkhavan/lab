package tracker

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	"github.com/fsnotify/fsnotify"
)

type Tracker struct {
	watcher *fsnotify.Watcher
}

func New() (*Tracker, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	t := &Tracker{
		watcher,
	}
	return t, nil
}

func (t *Tracker) Close() {
	t.watcher.Close()
}

func (t *Tracker) Track(file string) error {
	err := t.watcher.Add(file)
	if err != nil {
		return err
	}
	return nil
}

func track() {
	// defer watcher.Close()

	// done := make(chan bool)
	// go func() {
	// 	for {
	// 		select {
	// 		case event, ok := <-watcher.Events:
	// 			if !ok {
	// 				return
	// 			}
	// 			log.Println("event:", event)
	// 			if event.Op&fsnotify.Write == fsnotify.Write {
	// 				log.Println("modified file:", event.Name)
	// 			}
	// 		case err, ok := <-watcher.Errors:
	// 			if !ok {
	// 				return
	// 			}
	// 			log.Println("error:", err)
	// 		}
	// 	}
	// }()

	// <-done
}

func sync(remote string, path string) {
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
