package realize

// this code is imported from moby, unfortunately i can't import it directly as dependencies from its repo,
// cause there was a problem between moby vendor and fsnotify
// i have just added only the walk methods and some little changes to polling interval, originally set as static.

import (
	"errors"
	"github.com/fsnotify/fsnotify"
)

var (
	// errPollerClosed is returned when the poller is closed
	errPollerClosed = errors.New("poller is closed")
	// errNoSuchWatch is returned when trying to remove a watch that doesn't exist
	errNoSuchWatch = errors.New("watch does not exist")
)

type (
	// FileWatcher is an interface for implementing file notification watchers
	FileWatcher interface {
		Close() error
		Add(string) error
		Walk(string, bool) string
		Remove(string) error
		Errors() <-chan error
		Events() <-chan fsnotify.Event
	}
	// fsNotifyWatcher wraps the fsnotify package to satisfy the FileNotifier interface
	fsNotifyWatcher struct {
		*fsnotify.Watcher
	}
)

// EventWatcher returns an fs-event based file watcher
func EventWatcher() (FileWatcher, error) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	return &fsNotifyWatcher{Watcher: w}, nil
}

// Errors returns the fsnotify error channel receiver
func (w *fsNotifyWatcher) Errors() <-chan error {
	return w.Watcher.Errors
}

// Events returns the fsnotify event channel receiver
func (w *fsNotifyWatcher) Events() <-chan fsnotify.Event {
	return w.Watcher.Events
}

// Walk fsnotify
func (w *fsNotifyWatcher) Walk(path string, init bool) string {
	if err := w.Add(path); err != nil {
		return ""
	}
	return path
}
