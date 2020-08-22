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
		size := uint(len(record)) - 1
		if idx == 0 {
			inputData = newInputData(size)
			for i := uint(0); i < size; i++ {
				inputData.attributeNames[i] = record[i]
				inputData.attributes[i].name = record[i]
			}
			inputData.resultName = record[size]
		} else {
			if len(inputData.rows) == idx-1 {
				inputData.doubleRows()
			}
			for i := uint(0); i < size; i++ {
				inputData.rows[idx-1].values[inputData.attributeNames[i]] = record[i]
				if !containsString(inputData.attributes[i].values, record[i]) {
					inputData.attributes[i].values = append(inputData.attributes[i].values, record[i])
				}
			}
			inputData.rows[idx-1].result = true
			if record[size] == "no" {
				inputData.rows[idx-1].result = false
			}
		}
		idx++
	}
	return inputData, nil
}

func newInputData(size uint) *InputData {
	inputData := new(InputData)
	inputData.attributeNames = make([]string, size)
	inputData.attributes = make([]Attribute, size)
	for _, attribute := range inputData.attributes {
		attribute.values = make([]string, 0)
	}
	inputData.rows = make([]Row, 100)
	for i := range inputData.rows {
		inputData.rows[i].values = make(map[string]string)
	}
	return inputData
}

func (data *InputData) doubleRows() {
	newRows := make([]Row, len(data.rows))
	for i := range newRows {
		newRows[i].values = make(map[string]string)
	}
	data.rows = append(data.rows, newRows...)
}

func containsString(arr []string, elem string) bool {
	for _, val := range arr {
		if elem == val {
			return true
		}
	}
	return false
}
