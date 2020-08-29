package lab

import (
	"fmt"
)

type Lab struct {
	user string
	host string
	path string
}

func New(user, host, path string) *Lab {
	return &Lab{
		user, host, path,
	}
}

func (l *Lab) Adress() string {
	return fmt.Sprintf("%s@%s:%s", l.user, l.host, l.path)
}
