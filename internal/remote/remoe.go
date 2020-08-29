package remote

import (
	"fmt"
)

type Remote struct {
	user string
	host string
	path string
}

func New(user, host, path string) *Remote {
	return &Remote{
		user, host, path,
	}
}

func (l *Remote) Adress() string {
	return fmt.Sprintf("%s@%s:%s", l.user, l.host, l.path)
}
