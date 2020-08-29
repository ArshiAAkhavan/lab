package remote

import (
	"fmt"
)

type Remote struct {
	user string
	host string
	path string

	Name string
}

func New(name, user, host, path string) *Remote {
	return &Remote{
		user, host, path, name,
	}
}

func (l *Remote) Address() string {
	return fmt.Sprintf("%s@%s:%s", l.user, l.host, l.path)
}
