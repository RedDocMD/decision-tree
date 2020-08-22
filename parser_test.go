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
	t.Logf("%+v", inputData)
}
