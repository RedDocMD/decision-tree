package utils

import (
	"RedDocMD/decision_tree/parser"
	"testing"
)

func getInputData(t *testing.T) *parser.InputData {
	filename := "/home/deep/work/go/decision-tree/data/data1_19.csv"
	inputData, err := parser.ParseFile(filename)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	return inputData
}

func TestAttributePartitionCorrect(t *testing.T) {
	inputData := getInputData(t)
	attribute := "pclass"
	partitions, err := AttributePartition(inputData, inputData.Rows, attribute)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	t.Log("For attribute ", attribute)
	for _, variant := range inputData.AttributesMap[attribute].Values {
		t.Log("\t", variant, " : ", len(partitions[variant]), " rows")
	}
}

func TestAttributePartitionWrong(t *testing.T) {
	inputData := getInputData(t)
	attribute := "garbage"
	_, err := AttributePartition(inputData, inputData.Rows, attribute)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestEntropyWhole(t *testing.T) {
	inputData := getInputData(t)
	entropy := Entropy(inputData.Rows)
	if entropy < 0.0 || entropy > 1.0 {
		t.Errorf("%f entropy value out of range", entropy)
	}
	t.Log(entropy, " entropy value")
}
