package logcollector

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/golang/glog"
	"go.einride.tech/can/pkg/descriptor"
	"go.einride.tech/can/pkg/generated"
)

type ColumnInfo struct {
	name  string
	unit  string
	index int
}

type Record []string

type DataLog struct {
	database         *descriptor.Database
	columns          map[string]ColumnInfo
	columnOrder      []string
	records          map[int]*Record
	bucketOrder      []int
	bucketSize       int
	mostRecentRecord *Record
}

func NewDataLog(d *descriptor.Database, bucketSize int) *DataLog {
	dl := &DataLog{
		database:   d,
		columns:    map[string]ColumnInfo{"time": ColumnInfo{"time", "s", 0}},
		records:    map[int]*Record{},
		bucketSize: bucketSize,
	}
	dl.columnOrder = []string{"time"}

	idx := 1
	for _, m := range d.Messages {
		for _, s := range m.Signals {
			dl.columns[s.Name] = ColumnInfo{s.Name, s.Unit, idx}
			dl.columnOrder = append(dl.columnOrder, s.Name)
			idx++
		}
	}
	return dl
}

func frameToMap(m *generated.Message) map[string]string {
	dataMap := map[string]string{}
	f := (*m).Frame()
	for _, s := range (*m).Descriptor().Signals {
		key := s.Name
		switch {
		case s.Length == 1: // bool
			val := s.UnmarshalBool(f.Data)
			dataMap[key] = "0"
			if val {
				dataMap[key] = "1"
			}
		default:
			dataMap[key] = strconv.FormatFloat(s.UnmarshalPhysical(f.Data), 'g', 8, 64)
		}
	}
	return dataMap
}

func (dl *DataLog) calcTimeBucket(microseconds int) int {
	if dl.bucketSize < 1 {
		return microseconds
	}
	return microseconds / (dl.bucketSize * 1000)
}

func (dl *DataLog) getRecord(microseconds int) (*Record, error) {
	bucket := dl.calcTimeBucket(microseconds)

	// does the bucket ID already exist?
	r, existingBucket := dl.records[bucket]
	if existingBucket {
		// return it for value update
		return r, nil
	}

	record := make(Record, len(dl.columnOrder))
	if dl.mostRecentRecord == nil {
		// first record in log, needs to be filled with "0"
		for i := 0; i < len(dl.columnOrder); i++ {
			record[i] = "0"
		}
	} else {
		// new bucket, a previous record exists, copy data
		copy(record, *dl.mostRecentRecord)
	}

	record[0] = strconv.FormatFloat(float64(microseconds)/1_000_000, 'g', -1, 64)

	dl.mostRecentRecord = &record
	dl.records[bucket] = &record
	dl.bucketOrder = append(dl.bucketOrder, bucket)

	return &record, nil
}

func (dl *DataLog) AddDataRecord(m *generated.Message, microseconds int) error {
	record, err := dl.getRecord(microseconds)
	if err != nil {
		return err
	}

	dataMap := frameToMap(m)
	for colName, val := range dataMap {
		column, ok := dl.columns[colName]
		if !ok {
			return fmt.Errorf("Invalid column name: %s", colName)
		}
		(*record)[column.index] = val
	}

	return nil
}

func (dl *DataLog) Write(ofp *os.File) (int, int, error) {

	outputWriter := bufio.NewWriter(ofp)
	defer outputWriter.Flush()
	defer ofp.Close()

	bytesWritten := 0
	linesWritten := 0

	b, err := outputWriter.WriteString(strings.Join(dl.columnOrder, "\t") + "\n")
	if err != nil {
		return linesWritten, bytesWritten, err
	}
	bytesWritten += b
	linesWritten++

	units := make([]string, 0, len(dl.columnOrder))
	for _, colName := range dl.columnOrder {
		col, ok := dl.columns[colName]
		if !ok {
			return linesWritten,
				bytesWritten,
				fmt.Errorf("Unable to write headers, unknown colName: ", colName)
		}
		units = append(units, col.unit)
	}
	b, err = outputWriter.WriteString(strings.Join(units, "\t") + "\n")
	if err != nil {
		return linesWritten, bytesWritten, err
	}
	bytesWritten += b
	linesWritten++

	for _, bucket := range dl.bucketOrder {
		record, ok := dl.records[bucket]
		if !ok {
			glog.Fatalf("Unable to find record for bucket %d", bucket)
		}
		b, err = outputWriter.WriteString(strings.Join(*record, "\t") + "\n")
		if err != nil {
			return linesWritten, bytesWritten, err
		}
		bytesWritten += b
		linesWritten++
	}
	return linesWritten, bytesWritten, nil
}
