package controllers

import (
	"net/http"
	"testing"

	"github.com/astaxie/beego"
)

//func Test_defer(t *testing.T) {
//	username, _ := getAdmin(1)
//	if username != "admin" {
//		t.Error("getAdmin get data error")
//	}
//}

//func Benchmark_Get(b *testing.B) {
//	for i := 0; i < b.N; i++ { //use b.N for looping
//		_, err := http.Get("http://localhost:8080/")
//		if err != nil {
//			panic(err)
//		}
//		//		beego.Info(rsp.Body)
//	}
//}

func Benchmark_PB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := http.Get("http://localhost:8080/")
		if err != nil {
			beego.Info(err)
		}
	}

	//	b.RunParallel(ParallelBenchmark_PB)
}

func ParallelBenchmark_PB(pb *testing.PB) {
	for pb.Next() {
		_, err := http.Get("http://localhost:8080/")
		if err != nil {

			beego.Info(err)
		}
	}
}

//func Test_Get(t *testing.T) {
//	t.Log("I'am a log message!")
//	t.Log("start test Get")
//	_, err := http.Get("http://localhost:8080/")
//	if err != nil {
//		t.Error(err)
//	}
//	t.Log("finish test Get")

//}

//func Benchmark_normal(b *testing.B) {
//	for i := 0; i < b.N; i++ { //use b.N for looping
//		normalFunc()
//	}
//}
