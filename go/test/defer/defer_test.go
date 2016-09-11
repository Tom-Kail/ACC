package main

import (
	"testing"
)

//func Test_defer(t *testing.T) {
//	username, _ := getAdmin(1)
//	if username != "admin" {
//		t.Error("getAdmin get data error")
//	}
//}

func Benchmark_defer(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		deferFunc()
	}
}

func Benchmark_normal(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		normalFunc()
	}
}
