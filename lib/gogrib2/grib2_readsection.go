package gogrib2

import (
	"io"
)

// fetchSection expects the reader to be at the start of a GRIB2 section.
// It does not matter which one.
// This function returns the section number, a slice containing the data or an error if the section could not be read properly.
func fetchSection(data io.Reader) (int, []byte, error) {
	return 0, nil, nil
}
