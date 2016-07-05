package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
)

func getMem() (int, error) {
	cmd := exec.Command("/bin/sh", "-c", `free -m | sed -n '2p' | awk '{print ""$3/$2*100"%"}'`)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		return 0, err
	}
	fmt.Println(cmd.Args)
	err = cmd.Wait()
	if err != nil {
		return 0, err
	}
	tmpRst := out.String()
	length := len(tmpRst)
	rst := tmpRst[0 : length-2]

	f, err := strconv.ParseFloat(rst, 64)
	return int(f), err
}

func getCpu() (int, error) {
	cmd := exec.Command("/bin/sh", "-c", `vmstat | sed -n '3p' | awk '{print ""$13+$14""}'`)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		return 0, err
	}
	fmt.Println(cmd.Args)
	err = cmd.Wait()
	if err != nil {
		return 0, err
	}

	tmpRst := out.String()
	length := len(tmpRst)
	rst := tmpRst[0 : length-1]

	f, err := strconv.ParseInt(rst, 10, 64)
	return int(f), err
}

func OSStatus() (cpu int, mem int, err error) {
	cpu, err = getCpu()
	if err != nil {
		return
	}
	mem, err = getMem()
	return
}
