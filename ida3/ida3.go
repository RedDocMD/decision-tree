package ida3

import (
	"RedDocMD/decision_tree/parser"
	"RedDocMD/decision_tree/utils"
	"errors"
)

func entropyGain(rows []parser.Row, attribute string, data *parser.InputData) (float64, error) {
	originalEntropy := utils.Entropy(rows)
	partitions, err := utils.AttributePartition(data, rows, attribute)
	if err != nil {
		return 0.0, errors.New("Invalid atttribute")
	}
	newTotalEntropy := 0.0
	for _, partitionedRows := range partitions {
		newEntropy := utils.Entropy(partitionedRows)
		newTotalEntropy += float64(len(partitionedRows)) / float64(len(rows)) * newEntropy
	}
	return originalEntropy - newTotalEntropy, nil
}

func containsString(arr []string, elem string) bool {
	for _, val := range arr {
		if elem == val {
			return true
		}
	}
	return false
}

func attributeForMaxEntropyGain(rows []parser.Row, data *parser.InputData) string {
	maxGain := 0.0
	bestAttribute := ""
	for _, attribute := range data.AttributeNames {
		gain, _ := entropyGain(rows, attribute, data)
		if gain >= maxGain {
			maxGain = gain
			bestAttribute = attribute
		}
	}
	return bestAttribute
}
