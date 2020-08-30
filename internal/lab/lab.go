package lab

import (
	"lab/internal/remote"
	"lab/internal/tracker"
	"sync"
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

func (l *Lab) getRemoteByName(name string) *remote.Remote {
	for _, r := range l.remotes {
		if r.Name == name {
			return r
		}
	}
	return nil
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

func (l *Lab) AllowSync(remoteName string) {
	r := l.getRemoteByName(remoteName)
	r.SetSyncMode(true)
	l.syncAll(r)
}

func (l *Lab) DisableSync(remoteName string) {
	l.getRemoteByName(remoteName).SetSyncMode(false)
}

func (l *Lab) Track(path string) {
	l.tracker.Track(path)
}

func (l *Lab) syncFile(r *remote.Remote, file string, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		r.Sync(file)
		wg.Done()
	}()
}

func (l *Lab) syncAll(r *remote.Remote) {
	var wg sync.WaitGroup
	for _, p := range l.tracker.Paths() {
		l.syncFile(r, p, &wg)
	}
}

func (l *Lab) Start() {
	events := l.tracker.Start()
	var wg sync.WaitGroup
	go func() {
		for {
			select {
			case event := <-events:
				for _, r := range l.remotes {
					l.syncFile(r, event.Name, &wg)
				}
			}
		}
	}()
}
