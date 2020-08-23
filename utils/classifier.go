package utils

import (
	"RedDocMD/decision_tree/parser"
	"errors"
	"math"
)

// AttributePartition partitions rows of input data on basis of attribute
func AttributePartition(data *parser.InputData, rows []parser.Row, attribute string) (map[string][]parser.Row, error) {
	stat, _ := containsString(data.AttributeNames, attribute)
	if !stat {
		return nil, errors.New("Invalid attribute")
	}
	partitionedRows := make(map[string][]parser.Row)
	for _, row := range rows {
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

// Entropy finds the entropy of a given list of Rows
func Entropy(rows []parser.Row) float64 {
	tot := len(rows)
	pos := 0
	for _, row := range rows {
		if row.Result {
			pos++
		}
	}
	neg := tot - pos
	posFrac := float64(pos) / float64(tot)
	negFrac := float64(neg) / float64(tot)
	return -posFrac*math.Log(posFrac) - negFrac*math.Log(negFrac)
}
