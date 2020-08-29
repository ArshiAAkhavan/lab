package main

import (
	"lab/internal/lab"
	"lab/internal/remote"
	"sync"
)

func main() {

	l, _ := lab.New()

	l.AddRemote(remote.New("global-vm", "root", "172.16.8.223", "/root/ArshiA"))
	l.Track("./")
	l.Start()
	var wg sync.WaitGroup
	// tracker, err := tracker.New()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // tracker.Track("./a.txt")
	// tracker.Track("./")

	wg.Add(1)
	wg.Wait()
}
