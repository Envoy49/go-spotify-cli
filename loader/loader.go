package loader

import (
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/sirupsen/logrus"
)

var (
	instance *spinner.Spinner
	once     sync.Once
	mu       sync.Mutex
)

func InitializeSpinner() {
	once.Do(func() {
		instance = spinner.New(spinner.CharSets[17], 10*time.Millisecond)
		if err := instance.Color("bold", "yellow"); err != nil {
			logrus.Error(err)
		}
	})
}

// Start begins the spinner animation.
func Start() {
	mu.Lock()
	defer mu.Unlock()
	instance.Start()
}

// Stop ends the spinner animation.
func Stop() {
	mu.Lock()
	defer mu.Unlock()
	instance.Stop()
}
