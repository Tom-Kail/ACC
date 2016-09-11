package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"runtime/pprof"
	//	"fmt"
)

//go test -v -bench=".*"
//o test -bench=".*" -cpuprofile=cpu.prof -c
//import "fmt"
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func deferFunc() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			//			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
	}()
	panic(777)
}

func normalFunc() error {
	err := errors.New("777")
	if err != nil {
		//		fmt.Println("dd")
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	deferFunc()
}
