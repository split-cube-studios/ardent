package main

import (
	"flag"
	"github.com/split-cube-studios/ardent/assetutil"
)
var dir string

func init() {
	flag.StringVar(&dir, "d", "assets/", "Asset directory")
	flag.Parse()
}

func main() {
	assetutil.CreateAssets(dir)
}

