package main

import (
	"flag"

	"github.com/boxidau/crtdtomsl/megasquirtcan"

	"github.com/golang/glog"
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "WARNING")
	flag.Set("v", "2")
	flag.Parse()
}

func main() {
	glog.Info("test")
	d := megasquirtcan.Database
}
