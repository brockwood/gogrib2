/*
Package gogrib2 allows a user to decode a GRIB2 formatted file.  These
types of files are used to transmit meteorological data, both historyic
and current.

The library will simply takes a path to a GRIB2 file and return all
sections of a GRIB2 file.

*/
package gogrib2

import (
	"io"
)

// DecodeGribData takes a reader and will attempt to find then decode a GRIB2 record.
// This returns the number of bytes read and the record.  If a record cannot be found,
// then it returns an error.  Note that the bytes read can be used in case the file being
// read is part of a collection.
func DecodeGribData(data *io.Reader) (uint, error) {
	return 0, nil
}

func findGribRecordStart() (uint, error) {
	return 0, nil
}
