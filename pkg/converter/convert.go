package converter

import (
	"bytes"
	"encoding/csv"
	"strconv"
	"strings"
)

// ReadCSV to read the content of CSV File
func ConvertCSVToJSON(csvBytes []byte, separator rune) []byte {

	csvReader := csv.NewReader(bytes.NewReader(csvBytes))
	csvReader.Comma = separator
	csvContent, _ := csvReader.ReadAll()

	var buffer bytes.Buffer

	// Begin Writing JSON. Start with opening array bracket "["
	buffer.WriteString("[")
	for i, d := range csvContent[1:] {

		// Beginning of csv row. Begin new JSON object
		buffer.WriteString("{")
		for j, y := range d {
			buffer.WriteString(`"` + csvContent[0][j] + `":`)

			// Determine if cell type is Float
			_, fErr := strconv.ParseFloat(y, 32)
			// Determine if cell type is Bool
			_, bErr := strconv.ParseBool(y)
			if fErr == nil {
				buffer.WriteString(y)
			} else if bErr == nil {
				buffer.WriteString(strings.ToLower(y))
			} else {
				buffer.WriteString((`"` + y + `"`))
			}

			// End of properties in single row. If
			if j < len(d)-1 {
				buffer.WriteString(",")
			}

		}

		// End of object of the array
		buffer.WriteString("}")
		if i < len(csvContent)-2 {
			buffer.WriteString(",")
		}
	}

	buffer.WriteString(`]`)
	return buffer.Bytes()
}
