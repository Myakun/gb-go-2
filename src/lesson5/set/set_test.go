package set_test

import (
	"lesson5/set"
	"testing"
)

func BenchmarkSyncMutexW10R90(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 1; j <= 100; j++ {
			go func(index int, val float64) {
				if index%10 == 0 {
					set.WriteDataSyncMutex(index, val)
					return
				}

				_ = set.ReadDataSyncMutex(index)
			}(i, float64(j))
		}
	}
}

func BenchmarkSyncMutexW50R50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 1; j <= 100; j++ {
			go func(index int, val float64) {
				if index%2 == 0 {
					set.WriteDataSyncMutex(index, val)
					return
				}

				_ = set.ReadDataSyncMutex(index)
			}(i, float64(j))
		}
	}
}

func BenchmarkSyncMutexW90R10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 1; j <= 100; j++ {
			go func(index int, val float64) {
				if index%10 == 0 {
					_ = set.ReadDataSyncMutex(index)
					return
				}

				set.WriteDataSyncMutex(index, val)
			}(i, float64(j))
		}
	}
}

func BenchmarkSyncRWMutexW10R90(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 1; j <= 100; j++ {
			go func(index int, val float64) {
				if index%10 == 0 {
					set.WriteDataSyncRWMutex(index, val)
					return
				}

				_ = set.ReadDataSyncRWMutex(index)
			}(i, float64(j))
		}
	}
}

func BenchmarkSyncRWMutexW50R50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 1; j <= 100; j++ {
			go func(index int, val float64) {
				if index%2 == 0 {
					set.WriteDataSyncRWMutex(index, val)
					return
				}

				_ = set.ReadDataSyncRWMutex(index)
			}(i, float64(j))
		}
	}
}

func BenchmarkSyncRWMutexW90R10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 1; j <= 100; j++ {
			go func(index int, val float64) {
				if index%10 == 0 {
					_ = set.ReadDataSyncRWMutex(index)
					return
				}

				set.WriteDataSyncRWMutex(index, val)
			}(i, float64(j))
		}
	}
}
