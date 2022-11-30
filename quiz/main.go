package main

import (
	"flag"
	"fmt"
	"os"
)

const bold = "\033[1m"
const normal = "\033[0m"

//func checkerr(string msg) {
//	if err != nil {
//		fmt.Println(bold + "Error: " + normal + msg)
//	}
//}

func main() {

	// What is flag library?
	// the flag library allows you to pass various parameters
	// to the executable.

	// flag.String("stringType", "defaultValue",
	// "helpDescription")

	csvfile := flag.String("csv", "problems.csv", "a csv file with 'questions, answer'")
	flag.Parse()

	file, err := os.Open(*csvfile)
	if err != nil {
		fmt.Println(bold + "Error: " + normal + "No such file named " + *csvfile)
		os.Exit(1)
	}
	_ = file
}

// 4.08 seconds ...
