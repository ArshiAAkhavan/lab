package remote

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

type Remote struct {
	user string
	host string
	path string

	Name string

	syncMode bool
}

func New(name, user, host, path string) *Remote {
	return &Remote{
		user,
		host,
		path,
		name,
		false,
	}
}

func (r *Remote) address() string {
	return fmt.Sprintf("%s@%s:%s", r.user, r.host, r.path)
}

func (r *Remote) SetSyncMode(syncMode bool) {
	r.syncMode = syncMode
}

func (r *Remote) Sync(file string) {
	//todo delete still not working
	args := strings.Split(fmt.Sprintf("rsync -a --delete %s %s", file, r.address()), " ")
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
