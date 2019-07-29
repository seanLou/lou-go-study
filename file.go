package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file, e := os.Open("/Users/louguanyang/Downloads/ttco_result_20190417.txt")
	if e != nil {
		panic(e)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	xiaomiMap := make(map[string]string)
	for {
		lineByte, _, e := reader.ReadLine()
		if e == io.EOF {
			break
		}
		line := string(lineByte)

		split := strings.Split(line, "|+|")
		wfId := split[0]
		xiaomiMap[wfId] = line
	}
	end := time.Now()
	fmt.Println("read file use:", end.Sub(start))
	start = end
	resultFilePath := "/Users/louguanyang/Downloads/ttco_result_20190417_2.txt"
	resultFile, e := os.OpenFile(resultFilePath, os.O_CREATE|os.O_WRONLY, 0666)
	if e != nil {
		panic(e)
	}
	defer resultFile.Close()
	for _, v := range xiaomiMap {
		v = v + "\n"
		_, e := resultFile.WriteString(v)
		if e != nil {
			fmt.Println(e)
		}
	}
	end = time.Now()
	fmt.Println("write file use:", end.Sub(start))
}
