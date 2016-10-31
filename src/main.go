package main

import (
	"io/ioutil"
	"fmt"
	"flag"
	"strconv"
	"strings"

	"pearls"
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

	var max int
	b := pearls.NewBitmap()

	for _, line := range strings.Split(s, "\r\n") {
		if line == "" {
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Couldn't convert %q to integer\n", line)
			return
		}
		if n > max {
			max = n
		}
		b.SetBit(n)
	}

	for i := 0; i < max; i ++ {
		if b.GetBit(i) {
			fmt.Println(i)
		}
	}
}