package set

import "sync"

var (
	data    = make(map[int]float64)
	mutex   sync.Mutex
	rwMutex sync.RWMutex
)

func ReadDataSyncMutex(i int) float64 {
	mutex.Lock()
	val := data[i]
	mutex.Unlock()

	return val
}

func ReadDataSyncRWMutex(i int) float64 {
	rwMutex.RLock()
	val := data[i]
	rwMutex.RUnlock()

	return val
}

func WriteDataSyncMutex(i int, val float64) {
	mutex.Lock()
	data[i] = val
	mutex.Unlock()
}

func WriteDataSyncRWMutex(i int, val float64) {
	rwMutex.Lock()
	data[i] = val
	rwMutex.Unlock()
}
