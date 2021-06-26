package main

import (
	"os"

	"github.com/split-cube-studios/ardent/aautil"
)

func main() {

	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			aautil.CreateAssets(os.Args[i])
		}
	} else {
		aautil.CreateAssets("./")
	}
}
