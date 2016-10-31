package main

import (
	"io/ioutil"
	"fmt"
	"flag"
	"strconv"
	"strings"
)

var (
	inputFilename = flag.String("input_filename", "input.txt", "Name of input file containing list of newline-terminated integers.")
)

func main() {

	data, err := ioutil.ReadFile(*inputFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := string(data)
	for _, line := range strings.Split(s, "\r\n") {
		if line == "" {
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Couldn't convert %q to integer\n", line)
			return
		}
		fmt.Println(n)
	}
}