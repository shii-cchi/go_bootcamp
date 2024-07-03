package ex01

import (
	"day07/ex00"
	"testing"
)

func BenchmarkMinCoins(b *testing.B) {
	coins := []int{1, 2, 5, 10, 20, 100}
	val := 12345
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ex00.MinCoins(val, coins)
	}
}

func BenchmarkMinCoins2(b *testing.B) {
	coins := []int{1, 2, 5, 10, 20, 100}
	val := 12345
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ex00.MinCoins2(val, coins)
	}
}
