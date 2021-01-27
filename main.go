package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/boxidau/crtdtomsl/crtd"
	"github.com/boxidau/crtdtomsl/logcollector"
	"github.com/boxidau/crtdtomsl/megasquirtcan"

	"github.com/golang/glog"
)

var absoluteTimeFlag = flag.Bool(
	"absolutetime",
	false,
	"Set if you don't want to rebase timestamps to be relative to the first record",
)

var bucketSizeFlag = flag.Int(
	"bucket",
	50,
	"Produce 1 output row for every X ms of data collected",
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "WARNING")
	flag.Set("v", "2")
	flag.Parse()
}

func processFile(inputFile *os.File, datalog *logcollector.DataLog) {
	scanner := bufio.NewScanner(inputFile)
	timebase := 0
	first := true

	framesProcessed := 0
	totalLines := 0

	for scanner.Scan() {
		totalLines++
		record, err := crtd.ParseCRTDDataRecord(scanner.Text())
		if err != nil {
			continue
		}
		if *absoluteTimeFlag == false {
			if first {
				timebase = record.Microseconds
				record.Microseconds = 0
				first = false
			} else {
				record.Microseconds -= timebase
			}
		}

		msg, err := megasquirtcan.Messages().UnmarshalFrame(record.Frame)
		if err != nil {
			glog.Error(err)
			continue
		}
		err = datalog.AddDataRecord(&msg, record.Microseconds)
		if err != nil {
			glog.Errorf("Cannot add data record: %v", err)
		}
		framesProcessed++
	}

	glog.Infof("Successfully processed %d of %d records", framesProcessed, totalLines)

}

func main() {
	if flag.NArg() != 1 {
		glog.Fatal("No input file specified")
	}

	inputFileName := flag.Args()[0]
	ifp, err := os.Open(inputFileName)
	if err != nil {
		glog.Fatal(err)
	}
	defer ifp.Close()

	outputFileName := inputFileName + ".msl"
	ofp, err := os.Create(outputFileName)
	if err != nil {
		glog.Fatal(err)
	}

	datalog := logcollector.NewDataLog(
		megasquirtcan.Messages().Database(),
		*bucketSizeFlag,
	)

	processFile(ifp, datalog)
	l, b, err := datalog.Write(ofp)
	if err != nil {
		glog.Error(err)
	}
	glog.Infof("Wrote log file to %s", outputFileName)
	glog.Infof("Wrote %d lines (%d KiB)", l, b/1024)

}
