package solarlunar

import (
	"fmt"
	"testing"
)

func TestSolarToChineseLuanr(t *testing.T) {
	solarDate := "1990-05-06"
	fmt.Println(SolarToChineseLuanr(solarDate))
}

func TestSolarToSimpleLunar(t *testing.T) {
	solarDate := "1990-05-06"
	fmt.Println(SolarToSimpleLuanr(solarDate))
	solarDate = "2019-02-05"
	fmt.Println(SolarToSimpleLuanr(solarDate))
	y, m, d, bLeap := Solar2Lunar(2019, 2, 5)
	fmt.Println("Lunar: ", y, m, d, bLeap)
}

func TestLunarToSolar(t *testing.T) {
	lunarDate := "1990-04-12"
	fmt.Println(LunarToSolar(lunarDate, false))
	lunarDate = "2019-01-01"
	fmt.Println(LunarToSolar(lunarDate, false))
	y, m, d := Lunar2Solar(2019, 1, 1, false)
	fmt.Println("Solar: ", y, m, d)
}

func BenchmarkLunarToSolar(b *testing.B) {
	lunarDate := "2018-01-01"
	for i := 0; i < b.N; i++ {
		_ = LunarToSolar(lunarDate, false)
	}
}

func BenchmarkSolarToLunar(b *testing.B) {
	solarDate := "2018-02-16"
	for i := 0; i < b.N; i++ {
		_ = SolarToSimpleLuanr(solarDate)
	}
}

func BenchmarkLunar2Solar(b *testing.B) {
	y, m, d := 2018, 1, 1
	for i := 0; i < b.N; i++ {
		_, _, _ = Lunar2Solar(y, m, d, false)
	}
}

func BenchmarkSolar2Lunar(b *testing.B) {
	y, m, d := 2018, 2, 16
	for i := 0; i < b.N; i++ {
		_, _, _, _ = Solar2Lunar(y, m, d)
	}
}
