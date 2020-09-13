package main

import (
	"lab/cmd"
	"lab/internal/lab"
)

/*
 * todo:
 * add ssh open shell to remotes
 */

func main() {


	l, _ := lab.New()
	cmd.CMD_init(l)
	// l.AddRemote(remote.New("global-vm", "root", "172.16.8.223", "/root/ArshiA"))
	// l.AddRemote(remote.New("global-vm2", "root", "172.16.8.223", "/root/ArshiA2"))
	// l.Track("./")
	// l.Start()
	// l.AllowSync("global-vm")
	// time.Sleep(time.Second * 10)
	// l.AllowSync("global-vm2")

	// wait := make(chan bool)
	// <-wait
}
