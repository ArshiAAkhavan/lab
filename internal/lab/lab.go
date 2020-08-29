package lab

import (
	"lab/internal/remote"
	"lab/internal/tracker"
)

type Lab struct {
	tracker *tracker.Tracker
	remotes []remote.Remote
}

func New() (*Lab, error) {
	tracker, err := tracker.New()
	if err != nil {
		return nil, err
	}
	l := Lab{
		tracker,
		make([]remote.Remote, 0),
	}
	return &l, err
}

func (l *Lab) AddRemote(r remote.Remote) {
	l.remotes = append(l.remotes, r)
}

func (l *lab) RemoveRemote(remoteName string) {
	for _,r := l.remotes
}
