package main

import (
	"lab/internal/lab"
	"lab/internal/remote"
)

/*
 * todo:
 * add ssh open shell to remotes
 */

func main() {

	l, _ := lab.New()
	l.AddRemote(remote.New("global-vm", "root", "172.16.8.223", "/root/ArshiA"))
	l.AddRemote(remote.New("global-vm2", "root", "172.16.8.223", "/root/ArshiA2"))
	l.Track("./")
	l.Start()

	wait := make(chan bool)
	<-wait
}
