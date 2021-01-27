package crtd

import (
	"fmt"
	"strings"

	"go.einride.tech/can"
)

type CRTDCommentRecord struct {
	timestamp   float64
	messageType string
	message     string
}

// CRTDFrameRecord Represents a canbus data message from a CRTD log
type CRTDDataRecord struct {
	Microseconds int
	BusID        uint8
	Frame        can.Frame
}

func ParseCRTDDataRecord(input string) (CRTDDataRecord, error) {
	record := CRTDDataRecord{}

	tokens := strings.Split(input, " ")

	rxtx := ""
	headerLen := 0
	_, err := fmt.Sscanf(tokens[1], "%d%1s%d", &record.BusID, &rxtx, &headerLen)
	if err != nil {
		return CRTDDataRecord{}, err
	}
	seconds := 0
	microseconds := 0

	_, err = fmt.Sscanf(tokens[0], "%d.%d", &seconds, &microseconds)
	if err != nil {
		return CRTDDataRecord{}, err
	}
	if microseconds < 1000 { // only 3 digit precision
		microseconds *= 1000
	}
	seconds *= 1_000_000
	record.Microseconds = seconds + microseconds

	record.Frame = can.Frame{}
	if rxtx == "R" {
		record.Frame.IsRemote = true
	}
	if headerLen == 29 {
		record.Frame.IsExtended = true
	}

	_, err = fmt.Sscan(tokens[2], &record.Frame.ID)
	if err != nil {
		return CRTDDataRecord{}, err
	}

	for offset := 0; offset < len(tokens)-3; offset++ {
		if err != nil {
			return CRTDDataRecord{}, err
		}
		fmt.Sscanf(tokens[3+offset], "%x", &record.Frame.Data[offset])
		record.Frame.Length = uint8(offset) + 1
	}
	err = record.Frame.Validate()

	return record, err
}
