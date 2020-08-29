package lab

import (
	"lab/internal/remote"
	"lab/internal/tracker"
)

/*
 *	todo list:
 *	exclude path
 *
 *	save tracking path to sync new labs
 */

type Lab struct {
	tracker *tracker.Tracker
	remotes []*remote.Remote
}

func New() (*Lab, error) {
	tracker, err := tracker.New()
	if err != nil {
		return nil, err
	}
	l := Lab{
		tracker,
		make([]*remote.Remote, 0),
	}
	return &l, err
}

func (l *Lab) AddRemote(r *remote.Remote) {
	l.remotes = append(l.remotes, r)
}

func (l *Lab) RemoveRemote(remoteName string) {
	for i, r := range l.remotes {
		if r.Name == remoteName {
			l.remotes = append(l.remotes[:i], l.remotes[i+1:]...)
			return
		}
	}
}

func (l *Lab) Track(path string) {
	l.tracker.Track(path)
}

func (l *Lab) Start() {
	events := l.tracker.Start()
	go func() {
		for {
			select {
			case event := <-events:
				for _, r := range l.remotes {
					r.Sync(event.Name)
				}
			}
		}
	}()
}
