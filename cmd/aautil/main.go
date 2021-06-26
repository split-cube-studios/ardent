package main

import (
	"flag"

	"github.com/split-cube-studios/ardent/aautil"
)

var dir string

func init() {
	flag.StringVar(&dir, "d", "assets/", "Asset directory")
	flag.Parse()
}

func main() {
	aautil.CreateAssets(dir)
}
