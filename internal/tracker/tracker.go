package tracker

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

type Tracker struct {
	watcher *fsnotify.Watcher
	paths   []string
}

func New() (*Tracker, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	t := &Tracker{
		watcher, make([]string, 0),
	}
	return t, nil
}

func (t *Tracker) Close() {
	t.watcher.Close()
}

func (t *Tracker) Track(path string) error {
	err := t.watcher.Add(path)
	if err != nil {
		return err
	}
	t.paths = append(t.paths, path)
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

func (t *Tracker) Paths() []string {
	return t.paths
}
