package tracker

import (
	"log"

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

func (t *Tracker) start(events chan<- fsnotify.Event) {
	for {
		select {
		case event, ok := <-t.watcher.Events:
			if !ok {
				continue
			}
			log.Println("event:", event)
			if event.Op != 0 /*&fsnotify.Write == fsnotify.Write*/ {
				log.Println("modified file:", event.Name)
				events <- event
				// t.sync("root@172.16.8.223:/root/ArshiA", event.Name)
			}
		case err, ok := <-t.watcher.Errors:
			if !ok {
				continue
			}
			log.Println("error:", err)
		}
	}
}

func (t *Tracker) Start() <-chan fsnotify.Event {
	events := make(chan fsnotify.Event)
	go func() {
		t.start(events)
	}()
	return events
}
