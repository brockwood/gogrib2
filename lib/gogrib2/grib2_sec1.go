package gogrib2

import (
	"fmt"
)

// Section1 contains data found in the indicator section of a GRIB2 record.
// See http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_sect0.shtml for the format.
type Section1 struct {
	Discipline  uint8
	Edition     uint8
	MessageSize uint64
}

func decodeSec1(data []byte) error {
	return nil
}

// PrintSec1 will dump the section out as text in an easily readible format.
func (s *Section1) PrintSec1() {
	fmt.Println("Hello again!")
}
