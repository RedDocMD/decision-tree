package main

import (
	"RedDocMD/decision_tree/parser"
	"testing"
)

func TestParser(t *testing.T) {
	filename := "/home/deep/work/go/decision-tree/data/data1_19.csv"
	inputData, err := parser.ParseFile(filename)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	t.Logf("Attribute names: %+v", inputData.AttributeNames)
	t.Logf("Result name: %s", inputData.ResultName)
	for _, attr := range inputData.Attributes {
		t.Logf("Attribute: %s Values: %+v", attr.Name, attr.Values)
	}
	expectedLen := uint(2201)
	if inputData.Length != expectedLen {
		t.Errorf("Expected %d rows, got %d values", expectedLen, inputData.Length)
	}
	t.Logf("%+v", inputData.Rows[0])
}

func TestParserWithWrongFilename(t *testing.T) {
	filename := "hello"
	_, err := parser.ParseFile(filename)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
