package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//read word header and footer
	headerBytes, err := ioutil.ReadFile(HeaderFile)
	if err != nil {
		panic(err)
	}
	footerBytes, err := ioutil.ReadFile(FooterFile)
	if err != nil {
		panic(err)
	}
	output, err := os.Create(OutputFile)
	defer output.Close()
	if err != nil {
		panic(err)
	}
	// write header of word file
	_, err = output.Write(headerBytes)
	if err != nil {
		panic(err)
	}
	// write table header
	_, err = output.WriteString(TableHeader)
	if err != nil {
		panic(err)
	}
	// write table row
	row1 := fmt.Sprintf(TableRaw, "今天", "天气", "真好")
	row2 := fmt.Sprintf(TableRaw, "明天", "天气", "更好")
	row3 := fmt.Sprintf(TableRaw, "后天", "天气", "好极了")
	rows := []string{row1, row2, row3}
	for i := 0; i < 20; i++ {
		for _, row := range rows {
			_, err = output.WriteString(row)
			if err != nil {
				panic(err)
			}
		}
	}
	// write table footer
	_, err = output.WriteString(TableFooter)
	if err != nil {
		panic(err)
	}
	// write footer of word file
	_, err = output.Write(footerBytes)
	if err != nil {
		panic(err)
	}

}
