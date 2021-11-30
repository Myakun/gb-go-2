package set

import "sync"

var (
	data  = make(map[int]float64)
	mutex sync.Mutex
)

func ReadData(i int) float64 {
	mutex.Lock()
	val := data[i]
	mutex.Unlock()

	return val
}

func WriteData(i int, val float64) {
	mutex.Lock()
	data[i] = val
	mutex.Unlock()
}
