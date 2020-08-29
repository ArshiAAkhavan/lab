package main

import (
	"lab/internal/tracker"
	"log"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	tracker, err := tracker.New()
	if err != nil {
		log.Fatal(err)
	}
	tracker.Track("./a.txt")
	tracker.Track("./ali.txt")
	wg.Add(1)
	wg.Wait()
}
