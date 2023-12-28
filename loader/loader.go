package loader

import (
	"github.com/briandowns/spinner"
	"sync"
	"time"
)

var (
	instance *spinner.Spinner
	once     sync.Once
	mu       sync.Mutex
)

func InitializeSpinner() {
	once.Do(func() {
		instance = spinner.New(spinner.CharSets[17], 10*time.Millisecond)
		instance.Color("bold", "yellow")
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
