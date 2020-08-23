package ida3

import (
	"RedDocMD/decision_tree/parser"
	"testing"
)

func TestMaxEntropyAttribute(t *testing.T) {
	inputData := getInputData(t)
	bestAtttribute := attributeForMaxEntropyGain(inputData.Rows, inputData)
	if bestAtttribute == "" {
		t.Error("Expected an attribute")
	}
	t.Log(bestAtttribute)
}

func getInputData(t *testing.T) *parser.InputData {
	filename := "/home/deep/work/go/decision-tree/data/data1_19.csv"
	inputData, err := parser.ParseFile(filename)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	return inputData
}
