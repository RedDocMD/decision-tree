package parser

import (
	"encoding/csv"
	"io"
	"os"
)

// Attribute is an individual attribute of the data
type Attribute struct {
	name   string
	values []string
}

// Row represents and individual row of data
type Row struct {
	values map[string]string
	result bool
}

// InputData is the input to this program
type InputData struct {
	attributes     []Attribute
	attributeNames []string
	rows           []Row
	resultName     string
}

// ParseFile parses a CSV file to give an InputData pointer
func ParseFile(filename string) (*InputData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	idx := 0
	var inputData *InputData
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if idx == 0 {
			size := uint(len(record)) - 1
			inputData = newInputData(size)
			for i := uint(0); i < size; i++ {
				inputData.attributeNames[i] = record[i]
				inputData.attributes[i].name = record[i]
			}
			inputData.resultName = record[size]
		}
	}
}

func newInputData(size uint) *InputData {
	inputData := new(InputData)
	inputData.attributeNames = make([]string, size)
	inputData.attributes = make([]Attribute, size)
	inputData.rows = make([]Row, 100)
	return inputData
}
