package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBase := flag.Int("from", 10, "input base")
	outputBase := flag.Int("to", 10, "output base")
	inputSeparator := flag.String("is", " ", "input separator")
	outputSeparator := flag.String("os", " ", "output separator")
	outputASCII := flag.Bool("to-ascii", false, "output ascii")
	flag.Parse()

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		os.Exit(1)
	}

	inputBytes, err := ioutil.ReadAll(os.Stdin)
	defer os.Stdin.Close()
	if err != nil {
		os.Exit(2)
	}

	chunks := strings.Split(string(inputBytes), *inputSeparator)
	if err != nil {
		os.Exit(4)
	}
	inputs := []int64{}
	for _, chunk := range chunks {
		cleaned := strings.TrimSpace(chunk)
		i, err := strconv.ParseInt(cleaned, *inputBase, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not parse base %d value: %s\n", *inputBase, cleaned)
			os.Exit(3)
		}
		inputs = append(inputs, i)
	}
	if *outputASCII {
		b := []byte{}
		for _, input := range inputs {
			if input&0xff == input {
				b = append(b, byte(input))
			} else {
				os.Exit(5)
			}
		}
		fmt.Println(string(b))
		os.Exit(0)
	}
	outputs := []string{}
	for _, input := range inputs {
		outputs = append(outputs, strconv.FormatInt(input, *outputBase))
	}
	fmt.Println(strings.Join(outputs, *outputSeparator))
}
