package utils

import (
	"RedDocMD/decision_tree/parser"
	"errors"
)

// AttributePartition partitions rows of input data on basis of attribute
func AttributePartition(data *parser.InputData, attribute string) (map[string][]parser.Row, error) {
	stat, _ := containsString(data.AttributeNames, attribute)
	if !stat {
		return nil, errors.New("Invalid attribute")
	}
	partitionedRows := make(map[string][]parser.Row)
	for _, row := range data.Rows {
		partitionedRows[row.Values[attribute]] = append(partitionedRows[row.Values[attribute]], row)
	}
	return partitionedRows, nil
}

func containsString(arr []string, elem string) (bool, int) {
	for i, val := range arr {
		if elem == val {
			return true, i
		}
	}
	return false, -1
}
