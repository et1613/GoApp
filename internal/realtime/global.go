package realtime

import "sync"

var (
	globalHub     *Hub
	globalHubOnce sync.Once
)

// GetGlobalHub returns the singleton hub instance
func GetGlobalHub() *Hub {
	globalHubOnce.Do(func() {
		globalHub = NewHub()
		go globalHub.Run()
	})
	return globalHub
}
