package main

import (
	"fmt"
	"os"

	"github.com/nablaone/vertabelo2gorm"
)

func main() {

	if len(os.Args) != 4 {
		fmt.Println("Usage: vertabelo2gorm input.xml package_name output.xml")
	}

	in := os.Args[1]
	packageName := os.Args[2]
	out := os.Args[3]

	processor := &vertabelo2gorm.Processor{}
	processor.Package = packageName

	must := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	fin, err := os.Open(in)
	must(err)

	fout, err := os.Create(out)
	must(err)

	err = processor.Process(fin, fout)
	must(err)
}
