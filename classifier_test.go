package main

import (
	"RedDocMD/decision_tree/parser"
	"RedDocMD/decision_tree/utils"
	"testing"
)

func TestAttributePartitionCorrect(t *testing.T) {
	filename := "/home/deep/work/go/decision-tree/data/data1_19.csv"
	inputData, err := parser.ParseFile(filename)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	attribute := "pclass"
	partitions, err := utils.AttributePartition(inputData, attribute)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	t.Log("For attribute ", attribute)
	for _, variant := range inputData.AttributesMap[attribute].Values {
		t.Log("\t", variant, " : ", len(partitions[variant]), " rows")
	}
}
