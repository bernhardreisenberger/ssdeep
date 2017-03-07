package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bernhardreisenberger/ssdeep"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Println("Please provide a file path: ./ssdeep /tmp/file")
		return
	}
	if len(flag.Args()) == 1 {
		sdeep := ssdeep.NewSSDEEP()
		file, err := os.Open(flag.Args()[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		hash, err := sdeep.Fuzzy(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%v", &hash)
	}
	if len(flag.Args()) == 2 {
		sdeep := ssdeep.NewSSDEEP()
		f1, err := os.Open(flag.Args()[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		h1, err := sdeep.Fuzzy(f1)
		if err != nil {
			fmt.Println(err)
			return
		}
		sdeep = ssdeep.NewSSDEEP()
		f2, err := os.Open(flag.Args()[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		h2, err := sdeep.Fuzzy(f2)
		if err != nil {
			fmt.Println(err)
			return
		}
		score := ssdeep.HashDistance(*h1, *h2)
		fmt.Println(score)
	}
}
