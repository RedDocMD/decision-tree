package ida3

import (
	"RedDocMD/decision_tree/parser"
	"testing"
)

func TestMaxEntropyAttribute(t *testing.T) {
	inputData := getInputData(t)
	bestAtttribute := attributeForMaxEntropyGain(inputData.Rows, inputData, nil)
	if bestAtttribute == "" {
		t.Error("Expected an attribute")
	}
	t.Log(bestAtttribute)
}

func TestIDA3(t *testing.T) {
	inputData := getInputData(t)
	decisionTree := IDA3(inputData)
	t.Log("\n", decisionTree.String())
}

func getInputData(t *testing.T) *parser.InputData {
	filename := "/home/deep/work/go/decision-tree/data/data1_19.csv"
	inputData, err := parser.ParseFile(filename)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	return inputData
}
