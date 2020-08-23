package parser

import (
	"encoding/csv"
	"io"
	"os"
)

// Attribute is an individual attribute of the data
type Attribute struct {
	Name   string
	Values []string
}

// Row represents and individual row of data
type Row struct {
	Values map[string]string
	Result bool
}

// InputData is the input to this program
type InputData struct {
	Attributes     []Attribute
	AttributesMap  map[string]Attribute
	AttributeNames []string
	Rows           []Row
	ResultName     string
	Length         uint
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
				inputData.AttributeNames[i] = record[i]
				inputData.Attributes[i].Name = record[i]
			}
			inputData.ResultName = record[size]
		} else {
			if len(inputData.Rows) == idx-1 {
				inputData.doubleRows()
			}
			for i := uint(0); i < size; i++ {
				inputData.Rows[idx-1].Values[inputData.AttributeNames[i]] = record[i]
				if !containsString(inputData.Attributes[i].Values, record[i]) {
					inputData.Attributes[i].Values = append(inputData.Attributes[i].Values, record[i])
				}
			}
			inputData.Rows[idx-1].Result = true
			if record[size] == "no" {
				inputData.Rows[idx-1].Result = false
			}
			inputData.Length++
		}
		idx++
	}
	for i := range inputData.AttributeNames {
		inputData.AttributesMap[inputData.AttributeNames[i]] = inputData.Attributes[i]
	}
	return inputData, nil
}

func newInputData(size uint) *InputData {
	inputData := new(InputData)
	inputData.AttributeNames = make([]string, size)
	inputData.Attributes = make([]Attribute, size)
	inputData.AttributesMap = make(map[string]Attribute)
	inputData.Rows = make([]Row, 100)
	for i := range inputData.Rows {
		inputData.Rows[i].Values = make(map[string]string)
	}
	inputData.Length = 0
	return inputData
}

func (data *InputData) doubleRows() {
	newRows := make([]Row, len(data.Rows))
	for i := range newRows {
		newRows[i].Values = make(map[string]string)
	}
	data.Rows = append(data.Rows, newRows...)
}

func containsString(arr []string, elem string) bool {
	for _, val := range arr {
		if elem == val {
			return true
		}
	}
	return false
}
