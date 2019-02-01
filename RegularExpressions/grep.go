package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

var result = [][]string{}

func grep(re, filename string) {
	reg, err := regexp.Compile(re)
	if err != nil {
		return
	}
	fh, err := os.Open(filename)
	defer fh.Close()
	if err != nil {
		return
	}
	f := bufio.NewReader(fh)
	buf := make([]byte, 1024)
	for {
		buf, _, err = f.ReadLine()
		if err != nil {
			return
		}
		s := string(buf)
		data := reg.FindAllString(s, -1)
		if len(data) == 0 {
			continue
		}
		result = append(result, data)
	}
}


func main() {
	flag.Parse()
	if flag.NArg() == 2 {
		grep(flag.Arg(0), flag.Arg(1))
		fmt.Println(result)
	} else {
		fmt.Println("Wrong number of arguments.")
		os.Exit(1)
	}
}
