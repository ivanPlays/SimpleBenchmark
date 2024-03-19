package main

import (
    "fmt"
    "time"
    "math"
)

func IsPrime(x uint64) bool {
	if x < 2 {
		return false
	}
	if x % 2 == 0 {
		return false
	}
	if x < 9 {
		return false
	}
	for d := uint64(3); d < uint64(math.Sqrt(float64(x))); d += 2 {
		if x % d == 0 {
			return false
		}
	}
	return true
}

func bench(f func(uint64)bool) uint32 {
	var x uint64 = 0
	var k uint32
	t0 := time.Now()
	for time.Now().Sub(t0).Seconds() < 10.0 {
		if f(x) {
			k++
		}
		x++
	}
	return k
}

func main() {
	fmt.Println("------ ИННОВАЦИОННЫЙ БЕНЧМАРК ИМЕНИ В.И.ПОПОК ------")
	fmt.Printf("%s - Бенчмарк запущен\n", time.Now())
	k := bench(IsPrime)
	fmt.Printf("%s - Бенчмарк завершён\n", time.Now())
	fmt.Printf("За 10 секунд было найдено %v простых чисел\n", k)
	fmt.Println("------ ИСПОЛНЕНИЕ ЗАВЕРШЕНО ------")
	time.Sleep(15 * time.Second)
}
